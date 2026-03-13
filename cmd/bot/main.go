package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"yes_bot/config"
	"yes_bot/internal/adapter/telegram"
	"yes_bot/internal/infra/logger"
	"yes_bot/internal/usecase"
)

func main() {
	cfg := config.Load()

	if err := logger.Init(cfg.LogLevel); err != nil {
		log.Fatalf("failed to init logger: %v", err)
	}
	defer logger.Sync()

	logg := logger.L()

	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		logg.Fatalf("failed to create bot: %v", err)
	}

	logg.Infow("bot authorized", "username", bot.Self.UserName)

	sender := telegram.NewSender(bot)
	handleMessageUC := usecase.NewHandleMessageUseCase(sender, cfg.StickerFileID)
	updateHandler := telegram.NewUpdateHandler(bot, handleMessageUC)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	updateHandler.Run(ctx)
}
