package common

import "com.github/dsvdev/telego/pkg/common/sending"

type Message struct {
	ChatID int64
	Text   string
}

// MessageProcessor определяет функцию-обработчик входящих сообщений.
//
// Принимает:
//   - inputUpdate: указатель на входящее сообщение (*Message)
//   - outbox: канал для отправки исходящих сообщений (chan sending.TelegramSendable)
//
// Используется в longpolling-боте для обработки каждого входящего сообщения.
type MessageProcessor func(inputUpdate *Message, outbox chan sending.TelegramSendable)
