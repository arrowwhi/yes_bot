package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Sender struct {
	bot *tgbotapi.BotAPI
}

func NewSender(bot *tgbotapi.BotAPI) *Sender {
	return &Sender{bot: bot}
}

func (s *Sender) SendSticker(chatID int64, stickerFileID string, replyToMessageID int) error {
	msg := tgbotapi.NewSticker(chatID, tgbotapi.FileID(stickerFileID))
	if replyToMessageID > 0 {
		msg.ReplyToMessageID = replyToMessageID
	}
	_, err := s.bot.Send(msg)
	return err
}
