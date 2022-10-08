package telegram

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"psy_bot/lib/e"
	"strconv"
)

const (
	getUpdatesMethod   = "getUpdates"
	sendMessageMethod  = "sendMessage"
	editKeyboardMethod = "editMessageReplyMarkup"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

//tg-bot.com/bot<token>

func New(host string, token string) *Client {
	return &Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

func (c *Client) Updates(offset int, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest(getUpdatesMethod, q)
	if err != nil {
		return nil, err
	}

	var res UpdatesResponse

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res.Result, nil
}

func (c *Client) SendMessage(chatID int, text string, keyboard ...InlineKeyboardMarkup) error {

	q := url.Values{}

	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)

	// ↓ q.Add("keyboard_markup", keyboard)
	if len(keyboard) == 1 {
		jsonKeyboard, err := json.Marshal(keyboard[0])
		if err != nil {
			return e.Wrap("can't convert a keyboard to json", err)
		}
		stringKeyboard := string(jsonKeyboard[:])

		q.Add("reply_markup", stringKeyboard)
	}

	_, err := c.doRequest(sendMessageMethod, q)
	if err != nil {
		return e.Wrap("can't send a message", err)
	}

	return nil
}

func (c *Client) EditKeyboard(chatID int, messageID int, keyboard InlineKeyboardMarkup) error {
	q := url.Values{}

	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("message_id", strconv.Itoa(messageID))

	// ↓ q.Add("keyboard_markup", keyboard)
	jsonKeyboard, err := json.Marshal(keyboard)
	if err != nil {
		return e.Wrap("can't convert a keyboard to json", err)
	}
	stringKeyboard := string(jsonKeyboard[:])

	q.Add("reply_markup", stringKeyboard)

	_, err2 := c.doRequest(editKeyboardMethod, q)
	if err2 != nil {
		return e.Wrap("can't send a message", err)
	}

	return nil

}

func (c *Client) doRequest(method string, query url.Values) (data []byte, err error) {
	defer func() { err = e.WrapIfErr("can't do request", err) }()

	const errMsg = "can't do request"

	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
