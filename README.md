# Telegram Bot Helper

![Go Version](https://img.shields.io/badge/Go-%3E%3D1.22.5-blue) ![License](https://img.shields.io/badge/license-MIT-green)

## 📌 Описание
**Telego** — это библиотека на Go, которая упрощает работу с Telegram Bot API, предоставляя удобные методы для обработки сообщений, команд и кнопок.

## 🚀 Установка

Для установки используйте `go get`:

```sh
 go get github.com/dsvdev/telego
```

## 📖 Быстрый старт

### 1. Инициализация бота

```go
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
```

### 2. Обработка команд

```go
// not implemented yet
```

### 3. Кнопки и inline-клавиатура

```go
// not implemented yet
```

## 📌 Возможности
✅ exemple

## 📞 Контакты

Если у вас есть вопросы или предложения, пишите:
- Telegram: [@dsvtlg](https://t.me/dsvtlg)
- GitHub Issues

