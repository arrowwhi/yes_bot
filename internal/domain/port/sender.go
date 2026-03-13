package port

type MessageSender interface {
	SendSticker(chatID int64, stickerFileID string, replyToMessageID int) error
}
