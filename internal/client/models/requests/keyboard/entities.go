package keyboard

type Keyboard interface {
	IsKeyboard()
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

func (*InlineKeyboardMarkup) IsKeyboard() {}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}

type ReplyKeyboardMarkup struct {
	Keyboard       [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard bool               `json:"resize_keyboard"`
}

func (*ReplyKeyboardMarkup) IsKeyboard() {}

type KeyboardButton struct {
	Text string `json:"text"`
}
