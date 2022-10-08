package telegram

import (
	"fmt"
	"log"
	"psy_bot/lib/e"
	"strings"
)

const (
	RndCmd   = "/rnd"
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s'", text, username)

	//if isAddCmd(text) {
	//	return p.savePage(chatID, text, username)
	//}

	switch text {
	//case RndCmd:
	//	return p.sendRandom(chatID, username)
	case HelpCmd:
		return p.sendHelp(chatID)
	case StartCmd:
		return p.sendHello(chatID)
	default:
		return p.tg.SendMessage(chatID, msgUnknownCommand)
	}
}

//func (p *Processor) savePage(chatID int, pageURL string, username string) (err error) {
//	defer func() { err = e.WrapIfErr("can't do command: save page", err) }()
//
//	page := &storage.Page{
//		URL:      pageURL,
//		UserName: username,
//	}
//
//	IsExists, err := p.storage.IsExists(page)
//	if err != nil {
//		return err
//	}
//	if IsExists {
//		return p.tg.SendMessage(chatID, noCommand)
//	}
//
//	if err := p.storage.Save(page); err != nil {
//		return err
//	}
//
//	if err := p.tg.SendMessage(chatID, noCommand); err != nil {
//		return err
//	}
//
//	return nil
//}

/*
func (p *Processor) sendRandom(chatID int, username string) (err error)  {
	defer func() {err = e.WrapIfErr("can't do command: can't send random", err)}()

	page, err := p.storage.PickRandom(username)
	if err != nil && !errors.Is(err, storage.ErrNoSavedPages){
		return err
	}

	if errors.Is(err, storage.ErrNoSavedPages) {
		return p.tg.SendMessage(chatID, msgNoSavedPages)
	}

	if err:=p.tg.SendMessage(chatID, page.URL); err!=nil{
		return err
	}

	return p.storage.Remove(page)

}

func isAddCmd(text string) bool {
	return isURL(text)
}

func isURL(text string) bool {
	u, err := url.Parse(text)

	return err == nil && u.Host != ""
}
*/

func (p *Processor) sendHelp(chatID int) error {
	return p.tg.SendMessage(chatID, msgHelp)
}

func (p *Processor) sendHello(chatID int) error {
	inlineKeyboard := createKeyboardStart()
	return p.tg.SendMessage(chatID, msgHello, inlineKeyboard)
}

func (p *Processor) sendHelloPatient(chatID int) error {
	return p.tg.SendMessage(chatID, msgHelloPatient)
}

func (p *Processor) sendPatientCityChoose(chatID int) error {
	fmt.Println("Patient")
	inlineKeyboard := createKeyboardCityChoose()

	if err := p.sendHelloPatient(chatID); err != nil {
		return e.Wrap("can't use function <sendHelloPatient>", err)
	}

	return p.tg.SendMessage(chatID, msgChooseCity, inlineKeyboard)
}

func (p *Processor) sendPsychologistMenu(chatID int) error {
	fmt.Println("Psychologist")
	// Menu for psychologist login, registration
	return nil
}

func (p *Processor) sendPatientIssueChoose(chatID int) error {
	// Меню для того чтобы пациент указал с какой проблемой он обращается
	return nil
}
