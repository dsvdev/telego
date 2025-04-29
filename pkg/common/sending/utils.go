package sending

import (
	"github.com/dsvdev/telego/internal/client/models/requests"
	"github.com/dsvdev/telego/internal/client/models/requests/keyboard"
	"log"
)

func SendMessageToRequest(message *SendMessage) *requests.SendMessageRequest {
	req := &requests.SendMessageRequest{
		ChatID:      message.ChatID,
		Text:        message.Text,
		ReplyMarkup: SendKeyboardToRequest(message.Keyboard),
	}

	return req
}

func SendKeyboardToRequest(sendKeyboard SendKeyboard) keyboard.Keyboard {
	if len(sendKeyboard.Buttons) == 0 {
		return nil
	}

	switch sendKeyboard.Type {
	case InlineKeyboard:
		return keyboardButtonsToInlineKeyboardRequest(sendKeyboard.Buttons)
	case ReplyKeyboard:
		return keyboardButtonsToReplyKeyboardRequest(sendKeyboard.Buttons)
	default:
		log.Printf("Unknown keyboard type: %v", sendKeyboard.Type)
		return nil
	}
}

func SendPhotoByIdToRequest(message *SendPhotoById) *requests.SendPhotoByIdRequest {
	req := &requests.SendPhotoByIdRequest{
		ChatID:      message.ChatID,
		Text:        message.Text,
		Photo:       message.PhotoID,
		ReplyMarkup: SendKeyboardToRequest(message.Keyboard),
	}

	return req
}

func SendDiceToRequest(message *sendDice) *requests.SendDiceRequest {
	return &requests.SendDiceRequest{
		ChatID: message.ChatID,
	}
}

func keyboardButtonsToInlineKeyboardRequest(buttons KeyboardButtons) *keyboard.InlineKeyboardMarkup {
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

func keyboardButtonsToReplyKeyboardRequest(buttons KeyboardButtons) *keyboard.ReplyKeyboardMarkup {
	res := &keyboard.ReplyKeyboardMarkup{
		Keyboard: make([][]keyboard.KeyboardButton, 0, len(buttons)),
	}

	for _, row := range buttons {
		buttonRow := make([]keyboard.KeyboardButton, 0, len(row))
		for _, label := range row {
			button := keyboard.KeyboardButton{
				Text: label,
			}
			buttonRow = append(buttonRow, button)
		}
		res.Keyboard = append(res.Keyboard, buttonRow)
	}
	res.ResizeKeyboard = true

	return res
}
