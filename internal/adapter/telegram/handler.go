package telegram

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"yes_bot/internal/domain/entity"
	"yes_bot/internal/infra/logger"
	"yes_bot/internal/usecase"
)

type UpdateHandler struct {
	bot     *tgbotapi.BotAPI
	useCase *usecase.HandleMessageUseCase
}

func NewUpdateHandler(bot *tgbotapi.BotAPI, useCase *usecase.HandleMessageUseCase) *UpdateHandler {
	return &UpdateHandler{
		bot:     bot,
		useCase: useCase,
	}
}

func (h *UpdateHandler) Run(ctx context.Context) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := h.bot.GetUpdatesChan(u)

	for {
		select {
		case <-ctx.Done():
			logger.L().Info("shutting down update handler")
			return
		case update, ok := <-updates:
			if !ok {
				logger.L().Warn("updates channel closed")
				return
			}

			if update.Message == nil {
				continue
			}

			if update.Message.Text == "" {
				continue
			}

			logger.L().Debugw(
				"received message",
				"chat_id", update.Message.Chat.ID,
				"message_id", update.Message.MessageID,
				"text", update.Message.Text,
			)

			msg := entity.Message{
				ChatID:    update.Message.Chat.ID,
				MessageID: update.Message.MessageID,
				Text:      update.Message.Text,
			}

			if err := h.useCase.Handle(msg); err != nil {
				logger.L().Errorw("handle message error", "err", err)
			}
		}
	}
}
