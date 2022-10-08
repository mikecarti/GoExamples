package telegram

type UpdatesResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	ID            int              `json:"update_id"`
	Message       *IncomingMessage `json:"message"`
	CallbackQuery *CallbackQuery   `json:"callback_query"`
}

type IncomingMessage struct {
	MessageID      int                  `json:"message_id"`
	Text           string               `json:"text"`
	From           From                 `json:"from"`
	Chat           Chat                 `json:"chat"`
	InlineKeyboard InlineKeyboardMarkup `json:"reply_markup"`
}

type From struct {
	Username string `json:"username"`
}

type Chat struct {
	ID int `json:"id"`
}

type InlineKeyboardMarkup struct {
	KeyboardArray [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}

type CallbackQuery struct {
	ID      string           `json:"ID"`
	From    From             `json:"from"`
	Data    string           `json:"data"`
	Message *IncomingMessage `json:"message"`
}
