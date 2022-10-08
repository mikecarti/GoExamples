package telegram

import (
	"errors"
	"psy_bot/lib/e"
	"psy_bot/lib/i"
)

func (p *Processor) buttonPressed(data string, chatID int, messageID int, username string) error {
	switch data {
	case i.ChooseIssueButton:
		if err := p.deleteKeyboard(chatID, messageID); err != nil {
			return e.Wrap("can not delete a keyboard", err)
		}
		return p.sendPatientIssueChoose(chatID)

	case i.ChoosePatientButton:
		if err := p.deleteKeyboard(chatID, messageID); err != nil {
			return e.Wrap("can not delete a keyboard", err)
		}
		// if patient(chatID) is fully in DataBase:
		//		continue
		// if patient(chatID) is not fully in DataBase:
		//		ask to fill missing data
		return p.sendPatientCityChoose(chatID)

	case i.ChoosePsychologistButton:
		return p.sendPsychologistMenu(chatID)

	case i.City1, i.City2, i.City3, i.City4, i.City5, i.City6, i.City7, i.City8:
		// AddCityToUser(username, data)
		return nil
	default:
		err := errors.New("button data number is unknown")
		return e.Wrap("Could not process button call:", err)
	}
}
