package main

import (
	"com.github/dsvdev/telego/pkg/bot"
	"com.github/dsvdev/telego/pkg/common"
	"com.github/dsvdev/telego/pkg/common/sending"
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
