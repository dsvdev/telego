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
