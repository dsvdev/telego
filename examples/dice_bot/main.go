package main

import (
	config "github.com/dsvdev/telego/examples"
	"github.com/dsvdev/telego/pkg/bot"
	"github.com/dsvdev/telego/pkg/common"
	"github.com/dsvdev/telego/pkg/common/sending"
	"strconv"
	"sync"
	"time"
)

func main() {
	echoBot := bot.NewLongpollingTelegramBot(config.BOT_TOKEN)

	echo := func(inputMessage *common.Message, outbox chan sending.TelegramSendable) {
		getValue := sending.NewSendDice(inputMessage.ChatID, outbox)
		value := getValue()
		time.Sleep(3 * time.Second)
		outbox <- &sending.SendMessage{
			ChatID: inputMessage.ChatID,
			Text:   "Вам выпало: " + strconv.Itoa(value),
		}
	}

	echoBot.StartProcessUpdates(echo)

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
