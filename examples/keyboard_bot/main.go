package main

import (
	config "github.com/dsvdev/telego/examples"
	"github.com/dsvdev/telego/pkg/bot"
	"github.com/dsvdev/telego/pkg/common"
	"github.com/dsvdev/telego/pkg/common/sending"
	"sync"
)

func main() {
	keyboardBot := bot.NewLongpollingTelegramBot(config.BOT_TOKEN)

	keyboard := func(inputMessage *common.Message, outbox chan sending.TelegramSendable) {
		if inputMessage.Text == "Message with inline" {
			outbox <- &sending.SendMessage{
				ChatID: inputMessage.ChatID,
				Text:   "Example with inline keyboard",
				Keyboard: sending.SendKeyboard{
					Type: sending.InlineKeyboard,
					Buttons: [][]string{
						{"Button 1", "Button 2"},
						{"Cancel"},
					},
				},
			}
		} else {
			outbox <- &sending.SendMessage{
				ChatID: inputMessage.ChatID,
				Text:   "Example with reply keyboard",
				Keyboard: sending.SendKeyboard{
					Type: sending.ReplyKeyboard,
					Buttons: [][]string{
						{"Button 1", "Button 2"},
						{inputMessage.Text},
						{"Message with inline"},
					},
				},
			}
		}
	}

	keyboardBot.StartProcessUpdates(keyboard)

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
