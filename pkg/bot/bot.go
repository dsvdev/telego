package bot

import (
	"github.com/dsvdev/telego/internal/client"
	"github.com/dsvdev/telego/pkg/common"
	"github.com/dsvdev/telego/pkg/common/sending"
	"log"
	"sync"
)

type TelegramBot interface {
	StartProcessUpdates(processor common.MessageProcessor)
	SendMessage(message *sending.SendMessage)
}

func NewLongpollingTelegramBot(token string) TelegramBot {
	newBot := &longpollingTelegramBot{
		token:          token,
		offset:         0,
		inputUpdates:   make(chan *common.Message),
		outputMessages: make(chan sending.TelegramSendable),
	}

	go newBot.gettingUpdates()
	go newBot.processOutputMessages()
	return newBot
}

type longpollingTelegramBot struct {
	token          string
	offset         int64
	inputUpdates   chan *common.Message
	outputMessages chan sending.TelegramSendable
	once           sync.Once
}

func (b *longpollingTelegramBot) StartProcessUpdates(processor common.MessageProcessor) {
	b.once.Do(func() {
		log.Println("Starting Longpolling Telegram Bot")

		go func() {
			for {
				msg := <-b.inputUpdates
				log.Printf("Processing message %v", msg)
				processor(msg, b.outputMessages)
			}
		}()
	})
}

func (b *longpollingTelegramBot) SendMessage(message *sending.SendMessage) {
	b.outputMessages <- message
}

func (b *longpollingTelegramBot) gettingUpdates() {
	for {
		updates, err := client.GetUpdates(b.token, b.offset)
		if err != nil {
			log.Printf("Error getting updates: %v", err)
		}
		if updates == nil {
			continue
		}
		for _, update := range *updates {
			if update.Message != nil {
				log.Printf("Got update: %v", update.Message)
				msg := &common.Message{
					ChatID: update.Message.Chat.ID,
					Text:   update.Message.Text,
				}
				if len(update.Message.Photos) > 0 {
					msg.PhotoID = update.Message.Photos[0].FileID
				}
				b.inputUpdates <- msg
			}
			b.offset = update.ID + 1
		}
	}
}

func (b *longpollingTelegramBot) processOutputMessages() {
	for {
		b.processSendable(<-b.outputMessages)
	}
}

func (b *longpollingTelegramBot) processSendable(msg sending.TelegramSendable) {
	log.Printf("Processing output message %v", msg)
	err := msg.SendToTelegram(b.token)
	if err != nil {
		log.Printf("Error sending telegram message: %v", err)
	}
}
