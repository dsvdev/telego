package requests

import "github.com/dsvdev/telego/internal/client/models/requests/keyboard"

type SendMessageRequest struct {
	ChatID      int64                          `json:"chat_id"`
	Text        string                         `json:"text"`
	ReplyMarkup *keyboard.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}
