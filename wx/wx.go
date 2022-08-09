package wx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var ErrEmptyWebHook = fmt.Errorf("empty webhook")

// {
// 	"msgtype": "text",
// 	"text": {
// 		"content": "hello world",
// 		"mentioned_list":["@all"],
// 	}
// }
type Msg struct {
	MsgType string `json:"msgtype"`
	Text    Text   `json:"text"`
}

type Text struct {
	Content       string   `json:"content"`
	MentionedList []string `json:"mentioned_list"`
}

func SendMsg(content string) error {
	var (
		webhook string
		msg     Msg
		msgStr  []byte
		err     error
	)
	webhook = os.Getenv("WX_WEB_HOOK")
	if webhook == "" {
		return ErrEmptyWebHook
	}
	msg = Msg{
		MsgType: "text",
		Text: Text{
			Content:       content,
			MentionedList: []string{"@all"},
		},
	}
	if msgStr, err = json.Marshal(msg); err != nil {
		return err
	}
	_, err = http.Post(webhook, "application/json", bytes.NewReader(msgStr))
	return err
}
