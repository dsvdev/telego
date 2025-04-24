package sending

import (
	"github.com/dsvdev/telego/internal/client"
)

type TelegramSendable interface {
	SendToTelegram(token string) error
}

type KeyboardType string

const (
	InlineKeyboard KeyboardType = "InlineKeyboard"
	ReplyKeyboard  KeyboardType = "ReplyKeyboard"
)

type SendMessage struct {
	ChatID   int64
	Text     string
	Keyboard SendKeyboard
}

func (s *SendMessage) SendToTelegram(token string) error {
	smr := SendMessageToRequest(s)
	_, err := client.SendMessage(token, smr)
	return err
}

type SendKeyboard struct {
	Type    KeyboardType
	Buttons KeyboardButtons
}

type KeyboardButtons [][]string

type SendPhotoById struct {
	ChatID   int64
	Text     string
	PhotoID  string
	Keyboard SendKeyboard
}

func (s *SendPhotoById) SendToTelegram(token string) error {
	smr := SendPhotoByIdToRequest(s)
	_, err := client.SendPhoto(token, smr)
	return err
}
