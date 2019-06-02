package src

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// MessageConfig is a json file that put all custom messages
type MessageConfig struct {
	Welcome []*quickReply `json:"welcome"`
}

type quickReply struct {
	ImageURL string `json:"image_url"`
	Label    string `json:"label"`
	Text     string `json:"text"`
}

// InitMessage init the all custom message by message.json
func InitMessage() (*MessageConfig, error) {
	messageFile, err := os.Open(string(messagePath))
	if err != nil {
		return nil, err
	}
	defer messageFile.Close()

	byteMsg, err := ioutil.ReadAll(messageFile)
	if err != nil {
		return nil, err
	}

	m := &MessageConfig{}
	if err := json.Unmarshal(byteMsg, m); err != nil {
		return nil, err
	}

	return m, nil

}
