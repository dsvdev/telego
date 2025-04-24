package requests

import "github.com/dsvdev/telego/internal/client/models/requests/keyboard"

type SendMessageRequest struct {
	ChatID      int64             `json:"chat_id"`
	Text        string            `json:"text"`
	ReplyMarkup keyboard.Keyboard `json:"reply_markup,omitempty"`
}

type SendPhotoByIdRequest struct {
	ChatID      int64             `json:"chat_id"`
	Text        string            `json:"caption,omitempty"`
	ReplyMarkup keyboard.Keyboard `json:"reply_markup,omitempty"`
	Photo       string            `json:"photo"`
}
