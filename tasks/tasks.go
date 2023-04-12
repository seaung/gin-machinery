package tasks

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type Message struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Tatget  string `json:"target"`
}

func DecodingMessage(msg string, task interface{}) (err error) {
	bs64, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		return
	}

	content := []byte(bs64)
	err = json.Unmarshal(content, task)
	if err != nil {
		return
	}
	return
}

func SendMessage(content string) (bool, error) {
	message := Message{}

	DecodingMessage(content, &message)

	fmt.Printf("the message id : %d", message.ID)
	fmt.Printf("the content : %s", message.Content)
	fmt.Printf("the target : %s", message.Tatget)

	return true, nil
}
