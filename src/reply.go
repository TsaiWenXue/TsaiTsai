package src

import (
	"math/rand"

	"github.com/line/line-bot-sdk-go/linebot"
)

// HandleEvent deal with event sent from users.
func HandleEvent(bot *linebot.Client, event *linebot.Event, mConfig *MessageConfig) error {
	switch event.Type {
	case linebot.EventTypeMessage:
		switch message := event.Message.(type) {
		case *linebot.TextMessage:
			if _, err := bot.ReplyMessage(event.ReplyToken, handleTextMessage(mConfig, message.Text)).Do(); err != nil {
				return err
			}
		}
	case linebot.EventTypeJoin:
		if _, err := bot.ReplyMessage(event.ReplyToken, welcomeMessage(mConfig)).Do(); err != nil {
			return err
		}
	}

	return nil
}

func handleTextMessage(mConfig *MessageConfig, text string) linebot.SendingMessage {
	switch text {
	case string(handsomePhoto):
		return randHandsomePhoto(mConfig)
	case string(help):
		return linebot.NewTextMessage(string(helpReply))
	default:
		return linebot.NewTextMessage(string(defaultReply))
	}
}

func randHandsomePhoto(msg *MessageConfig) linebot.SendingMessage {
	randNum := rand.Intn(len(msg.HandsomePhoto))
	url := msg.HandsomePhoto[randNum]

	return linebot.NewImageMessage(url, url)
}

func welcomeMessage(msg *MessageConfig) linebot.SendingMessage {
	items := []*linebot.QuickReplyButton{}

	for _, q := range msg.Welcome {
		ma := linebot.NewMessageAction(q.Label, q.Text)
		qrb := linebot.NewQuickReplyButton(q.ImageURL, ma)
		items = append(items, qrb)
	}

	quickReply := &linebot.QuickReplyItems{Items: items}

	return linebot.NewTextMessage(string(welcome)).WithQuickReplies(quickReply)
}
