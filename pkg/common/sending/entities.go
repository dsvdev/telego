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

// NewSendDice создаёт и инициализирует объект отправки "dice" сообщения в Telegram,
// и возвращает функцию, которая при вызове инициирует отправку и ожидание результата броска кубика.
//
// Параметры:
//   - chatID: идентификатор чата в Telegram, куда будет отправлено сообщение с анимацией dice.
//   - outbox: канал типа TelegramSendable, через который объекты отправляются для отправки сообщений.
//
// Поведение:
// Функция создаёт внутренний объект типа sendDice, настроенный для отправки "dice" сообщения.
// Она возвращает функцию, которая при вызове:
//   - отправляет объект sendDice в указанный канал outbox для асинхронной отправки сообщения в Telegram;
//   - блокирует выполнение до получения результата от Telegram API;
//   - после получения ответа возвращает значение выпавшего кубика.
//
// Возвращаемое значение:
//   - Функция без аргументов типа func() int.
//     При вызове этой функции происходит отправка запроса и ожидание ответа.
//     Возвращается значение выпавшего кубика (целое число от 1 до 6).
//
// Примечания:
//   - Вызов возвращённой функции блокирует выполнение до получения ответа.
//   - Если Telegram API не вернёт ответ (например, из-за сетевой ошибки),
//     выполнение может заблокироваться навсегда, если не предусмотрен таймаут или отмена.
//
// Пример использования:
//
//	getValue := NewSendDice(chatID, outbox)
//	value := getValue()
//	fmt.Printf("Выпало значение: %d\n", value)
func NewSendDice(chatID int64, outbox chan TelegramSendable) func() int {
	sd := &sendDice{
		ChatID: chatID,
		done:   make(chan struct{}),
		outbox: outbox,
	}

	return sd.GetValue
}

type sendDice struct {
	ChatID int64
	value  int
	done   chan struct{}
	outbox chan TelegramSendable
}

func (s *sendDice) SendToTelegram(token string) error {
	sdr := SendDiceToRequest(s)
	response, err := client.SendDice(token, sdr)
	if err != nil {
		return err
	}
	s.value = response.Dice.Value
	close(s.done)
	return nil
}

func (s *sendDice) GetValue() int {
	s.outbox <- s
	<-s.done
	return s.value
}
