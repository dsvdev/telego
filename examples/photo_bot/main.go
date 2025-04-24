package main

import (
	config "github.com/dsvdev/telego/examples"
	"github.com/dsvdev/telego/pkg/bot"
	"github.com/dsvdev/telego/pkg/common"
	"github.com/dsvdev/telego/pkg/common/sending"
	"sync"
)

func main() {
	echoBot := bot.NewLongpollingTelegramBot(config.BOT_TOKEN)

	echo := func(inputMessage *common.Message, outbox chan sending.TelegramSendable) {
		if inputMessage.PhotoID == "" {
			outbox <- &sending.SendMessage{
				ChatID: inputMessage.ChatID,
				Text:   "Пришлите фото для определения ID",
			}
			return
		}

		outbox <- &sending.SendPhotoById{
			ChatID:  inputMessage.ChatID,
			PhotoID: inputMessage.PhotoID,
			Text:    inputMessage.PhotoID,
		}
	}

	echoBot.StartProcessUpdates(echo)

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
