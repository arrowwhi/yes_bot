package usecase

import (
	"strings"

	"yes_bot/internal/domain/entity"
	"yes_bot/internal/domain/port"
)

type HandleMessageUseCase struct {
	sender        port.MessageSender
	stickerFileID string
}

func NewHandleMessageUseCase(sender port.MessageSender, stickerFileID string) *HandleMessageUseCase {
	return &HandleMessageUseCase{
		sender:        sender,
		stickerFileID: stickerFileID,
	}
}

func (uc *HandleMessageUseCase) Handle(msg entity.Message) error {
	text := strings.TrimSpace(msg.Text)
	lower := strings.ToLower(text)
	if strings.HasSuffix(lower, "да") {
		return uc.sender.SendSticker(msg.ChatID, uc.stickerFileID, msg.MessageID)
	}

	return nil
}
