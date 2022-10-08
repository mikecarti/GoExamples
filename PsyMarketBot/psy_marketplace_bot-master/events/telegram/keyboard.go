package telegram

import (
	"errors"
	"psy_bot/clients/telegram"
	"psy_bot/lib/e"
	"psy_bot/lib/i"
)

type Keyboard telegram.InlineKeyboardMarkup
type Button telegram.InlineKeyboardButton

func (keyboard Keyboard) Markup() telegram.InlineKeyboardMarkup {
	return keyboard.Markup()
}

func (button Button) Button() telegram.InlineKeyboardButton {
	return button.Button()
}

func createKeyboardStart() (inlineKeyboard telegram.InlineKeyboardMarkup) {
	patientButton := telegram.InlineKeyboardButton{
		Text:         "Я Клиент",
		CallbackData: i.ChoosePatientButton,
	}

	psyButton := telegram.InlineKeyboardButton{
		Text:         "Я Психолог",
		CallbackData: i.ChoosePsychologistButton,
	}

	arrButtons := [][]telegram.InlineKeyboardButton{
		{patientButton, psyButton},
	}

	inlineKeyboard = telegram.InlineKeyboardMarkup{
		KeyboardArray: arrButtons,
	}

	return inlineKeyboard
}

func createKeyboardCityChoose() (inlineKeyboard telegram.InlineKeyboardMarkup) {
	city1Button := telegram.InlineKeyboardButton{
		Text:         "city1",
		CallbackData: i.City1,
	}

	city2Button := telegram.InlineKeyboardButton{
		Text:         "city2",
		CallbackData: i.City2,
	}
	city3Button := telegram.InlineKeyboardButton{
		Text:         "city3",
		CallbackData: i.City3,
	}

	city4Button := telegram.InlineKeyboardButton{
		Text:         "city4",
		CallbackData: i.City4,
	}

	city5Button := telegram.InlineKeyboardButton{
		Text:         "city5",
		CallbackData: i.City5,
	}

	city6Button := telegram.InlineKeyboardButton{
		Text:         "city6",
		CallbackData: i.City6,
	}
	city7Button := telegram.InlineKeyboardButton{
		Text:         "city7",
		CallbackData: i.City7,
	}

	city8Button := telegram.InlineKeyboardButton{
		Text:         "city8",
		CallbackData: i.City8,
	}

	arrButtons := [][]telegram.InlineKeyboardButton{
		{city1Button, city2Button, city3Button},
		{city4Button, city5Button, city6Button},
		{city7Button, city8Button},
	}

	inlineKeyboard = telegram.InlineKeyboardMarkup{
		KeyboardArray: arrButtons,
	}

	return inlineKeyboard
}

func createEmptyKeyboard() (emptyKeyboard telegram.InlineKeyboardMarkup) {
	var arrButtons [][]telegram.InlineKeyboardButton
	{
	}
	emptyKeyboard = telegram.InlineKeyboardMarkup{
		KeyboardArray: arrButtons,
	}
	return emptyKeyboard
}

// interface{} because can be either string or int
func (p *Processor) deleteKeyboard(chatID int, MessageID int) error {
	emptyKeyboard := createEmptyKeyboard()
	return p.tg.EditKeyboard(chatID, MessageID, emptyKeyboard)
}

func CreateCustomKeyboard(length int, height int, titles []string, callbackIDs []string) (telegram.InlineKeyboardMarkup, error) {
	var arrButtons [][]telegram.InlineKeyboardButton

	currIndex := 0

	// Error Handling
	if len(titles) != length*height {
		err := errors.New("titles length is not equal to quantity of buttons")
		return telegram.InlineKeyboardMarkup{}, e.Wrap("", err)
	}

	if len(callbackIDs) != length*height {
		err := errors.New("callbackIDs length is not equal to quantity of buttons")
		return telegram.InlineKeyboardMarkup{}, e.Wrap("", err)
	}

	// making an InlineKeyboard
	for n := 0; n < height; n++ {
		var currSlice []telegram.InlineKeyboardButton
		for j := 0; j < length; j++ {
			button := Button{
				Text:         titles[currIndex],
				CallbackData: callbackIDs[currIndex],
			}
			// подозрительная строка из-за button.Button()
			currSlice = append(currSlice, button.Button())
			currIndex++
		}
		arrButtons = append(arrButtons, currSlice)
	}

	keyboard := Keyboard{
		KeyboardArray: arrButtons,
	}

	return keyboard.Markup(), nil

}
