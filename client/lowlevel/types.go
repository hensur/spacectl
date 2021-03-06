package lowlevel

import (
	"fmt"
	"net/url"
	"strings"
)

type Message struct {
	Message string      `json:"msg"`
	Error   interface{} `json:"error"`
}

func (m Message) String() string {
	if m.Message != "" && m.Error != "" && m.Error != nil {
		return fmt.Sprintf("%s (%s)", m.Message, m.Error)
	}

	if m.Message != "" {
		return m.Message
	}

	if m.Error != nil {
		switch e := m.Error.(type) {
		case string:
			if e != "" {
				return e
			}
		default:
			return fmt.Sprintf("%v", e)
		}
	}

	return "Unknown"
}

type Link struct {
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Method string `json:"method"`
}

func (l Link) WithParam(param, value string) Link {
	return Link{
		Href:   strings.Replace(l.Href, "{"+param+"}", value, 1),
		Rel:    l.Rel,
		Method: l.Method,
	}
}

func (l Link) Execute(client *SpacesLowlevelClient, body interface{}, result interface{}) error {
	switch strings.ToUpper(l.Method) {
	case "POST":
		return l.Post(client, body, result)
	case "PUT":
		return l.Put(client, body, result)
	case "DELETE":
		return l.Delete(client, result)
	case "GET":
		return l.Get(client, result)
	default:
		return l.Get(client, result)
	}
}

func (l Link) Get(client *SpacesLowlevelClient, result interface{}) error {
	return client.Get(l.Href, result)
}

func (l Link) GetWithQuery(q url.Values, client *SpacesLowlevelClient, result interface{}) error {
	if len(q) == 0 {
		return l.Get(client, result)
	}

	u := fmt.Sprintf("%s?%s", l.Href, q.Encode())
	return client.Get(u, result)
}

func (l Link) Post(client *SpacesLowlevelClient, body interface{}, result interface{}) error {
	return client.Post(l.Href, body, result)
}

func (l Link) Put(client *SpacesLowlevelClient, body interface{}, result interface{}) error {
	return client.Put(l.Href, body, result)
}

func (l Link) Delete(client *SpacesLowlevelClient, result interface{}) error {
	return client.Delete(l.Href, result)
}

type LinkList []Link

func (l LinkList) HasLink(rel string) bool {
	for i := range l {
		if l[i].Rel == rel {
			return true
		}
	}

	return false
}

func (l LinkList) GetLinkByRel(rel string) (*Link, error) {
	for i := range l {
		if l[i].Rel == rel {
			return &l[i], nil
		}
	}

	return nil, ErrLinkNotFound{rel}
}

type Linkeable struct {
	Links   LinkList `json:"_links"`
	Actions LinkList `json:"_actions"`
}
