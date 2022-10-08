package telegram

import (
	"errors"
	"psy_bot/clients/telegram"
	"psy_bot/events"
	"psy_bot/lib/e"
	"psy_bot/lib/storage"
)

type Processor struct {
	tg      *telegram.Client
	offset  int
	storage storage.Storage
}

type Meta struct {
	ChatID    int
	Username  string
	messageID int
	//// fields for callbackQuery (not sure)
	Data    string
	Message *telegram.IncomingMessage
}

var (
	ErrUnknownEventType = errors.New("unknown event type")
	ErrUnknownMetaType  = errors.New("unknown meta type")
)

func New(client *telegram.Client, storage storage.Storage) *Processor {
	return &Processor{
		tg:      client,
		storage: storage,
	}
}

func (p *Processor) Fetch(limit int) ([]events.Event, error) {
	updates, err := p.tg.Updates(p.offset, limit)
	if err != nil {
		return nil, e.Wrap("can't get events", err)
	}

	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]events.Event, 0, len(updates))

	for _, u := range updates {
		res = append(res, event(u))
	}

	p.offset = updates[len(updates)-1].ID + 1

	return res, nil
}

func (p *Processor) Process(event events.Event) error {
	switch event.Type {
	case events.Message:
		return p.processMessage(event)
	case events.CallbackQuery:
		return p.processCallback(event)
	default:
		return e.Wrap("can't process message", ErrUnknownEventType)
	}
}

func (p *Processor) processMessage(event events.Event) error {
	meta, err := meta(event)
	if err != nil {
		return e.Wrap("can't process message", err)
	}

	if err := p.doCmd(event.Text, meta.ChatID, meta.Username); err != nil {
		return e.Wrap("can't process message", err)
	}

	return nil

}

func (p *Processor) processCallback(event events.Event) error {

	meta, err := meta(event)
	if err != nil {
		return e.Wrap("can't process callback", err)
	}

	//  Я не хочу использовать здесь doCmd, чтобы не позволять пользователям
	//  иметь доступ ко всем кнопочным коммандам

	if err := p.buttonPressed(meta.Data, meta.Message.Chat.ID, meta.Message.MessageID, meta.Username); err != nil {
		return e.Wrap("can't process callback", err)
	}

	return nil
}

func meta(event events.Event) (Meta, error) {
	res, ok := event.Meta.(Meta)
	if !ok {
		return Meta{}, e.Wrap("can't get meta", ErrUnknownMetaType)
	}

	return res, nil

}

func event(update telegram.Update) events.Event {
	updType := fetchType(update)

	res := events.Event{
		Type: updType,
		Text: fetchText(update),
	}

	if updType == events.Message {
		res.Meta = Meta{
			ChatID:    update.Message.Chat.ID,
			Username:  update.Message.From.Username,
			messageID: update.Message.MessageID,
		}
	}

	if updType == events.CallbackQuery {
		res.Meta = Meta{
			Username: update.CallbackQuery.From.Username,
			Data:     update.CallbackQuery.Data,
			Message:  update.CallbackQuery.Message,
		}
	}

	return res
}

func fetchText(upd telegram.Update) string {
	if upd.Message == nil {
		return ""
	}
	return upd.Message.Text
}

func fetchType(upd telegram.Update) events.Type {
	if upd.Message != nil {
		return events.Message
	}

	if upd.CallbackQuery != nil {
		return events.CallbackQuery
	}

	return events.Unknown
}
