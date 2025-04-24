package sending

import (
	"com.github/dsvdev/telego/internal/client/models/requests"
	"com.github/dsvdev/telego/internal/client/models/requests/keyboard"
)

func SendMessageToRequest(message *SendMessage) *requests.SendMessageRequest {
	req := &requests.SendMessageRequest{
		ChatID: message.ChatID,
		Text:   message.Text,
	}

	if len(message.Keyboard.Buttons) > 0 {
		req.ReplyMarkup = keyboardButtonsToRequest(message.Keyboard.Buttons)
	}

	return req
}

func keyboardButtonsToRequest(buttons KeyboardButtons) *keyboard.InlineKeyboardMarkup {
	res := &keyboard.InlineKeyboardMarkup{
		InlineKeyboard: make([][]keyboard.InlineKeyboardButton, 0, len(buttons)),
	}

	for _, row := range buttons {
		buttonRow := make([]keyboard.InlineKeyboardButton, 0, len(row))
		for _, label := range row {
			button := keyboard.InlineKeyboardButton{
				Text:         label,
				CallbackData: label,
			}
			buttonRow = append(buttonRow, button)
		}
		res.InlineKeyboard = append(res.InlineKeyboard, buttonRow)
	}

	return res
}
