package src

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// MessageConfig is a json file that put all custom messages.
type MessageConfig struct {
	DefaultReply    []string                   `json:"default_reply"`
	Welcome         []*quickReply              `json:"welcome"`
	HandsomePhoto   []string                   `json:"handsome_photo"`
	ProjectCarousel []*projectCarouselTemplate `json:"project_carousel"`
	News            *newsConfig                `json:"news"`
	InfoTemplate    *infoTemplateConfig        `json:"info_template"`
}

type quickReply struct {
	ImageURL string `json:"image_url"`
	Label    string `json:"label"`
	Text     string `json:"text"`
}

type projectCarouselTemplate struct {
	ThumbnailImageURL string       `json:"thumbnailImageUrl,omitempty"`
	Title             string       `json:"title,omitempty"`
	Text              string       `json:"text"`
	Actions           []*uriAction `json:"actions"`
}

type uriAction struct {
	URI   string `json:"uri"`
	Label string `json:"label"`
}

type newsConfig struct {
	RefreshPeriod int `json:"refresh_period"`
	ChanBuffer    int `json:"chan_buffer"`
}

type infoTemplateConfig struct {
	IMG            string `json:"img"`
	Characteristic string `json:"characteristic"`
	Habit          string `json:"habit"`
	Motto          string `json:"motto"`
}

// InitMessageConfig init the all custom message config by message.json.
func InitMessageConfig() (*MessageConfig, error) {
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
