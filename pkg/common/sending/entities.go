package sending

import (
	"com.github/dsvdev/telego/internal/client"
)

type TelegramSendable interface {
	SendToTelegram(token string) error
}

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
	Type    interface{}
	Buttons KeyboardButtons
}

type KeyboardButtons [][]string
