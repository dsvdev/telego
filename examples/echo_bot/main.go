package main

import (
	"github.com/dsvdev/telego/pkg/bot"
	"github.com/dsvdev/telego/pkg/common"
	"github.com/dsvdev/telego/pkg/common/sending"
	"sync"
)

const token = "YOUR_TOKEN_HERE"

func main() {
	echoBot := bot.NewLongpollingTelegramBot(token)

	echo := func(inputMessage *common.Message, outbox chan sending.TelegramSendable) {
		outbox <- &sending.SendMessage{
			ChatID: inputMessage.ChatID,
			Text:   "Your inputMessage: " + inputMessage.Text,
		}
	}

	echoBot.StartProcessUpdates(echo)

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
