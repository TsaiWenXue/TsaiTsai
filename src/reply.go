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
		if _, err := bot.ReplyMessage(event.ReplyToken, messageWithQuickReply(welcome, mConfig)).Do(); err != nil {
			return err
		}
	}

	return nil
}

func handleTextMessage(mConfig *MessageConfig, text string) linebot.SendingMessage {
	switch specialWord(text) {
	case handsomePhoto:
		return randHandsomePhoto(mConfig)
	case help:
		return messageWithQuickReply(string(helpReply), mConfig)
	case project:
		return projectCarousel(mConfig)
	default:
		return messageWithQuickReply(defaultReply, mConfig)
	}
}

func randHandsomePhoto(msg *MessageConfig) linebot.SendingMessage {
	randNum := rand.Intn(len(msg.HandsomePhoto))
	url := msg.HandsomePhoto[randNum]

	return linebot.NewImageMessage(url, url)
}

func projectCarousel(msg *MessageConfig) linebot.SendingMessage {
	columns := []*linebot.CarouselColumn{}

	for _, c := range msg.ProjectCarousel {
		act := linebot.NewURIAction(c.Actions[0].Label, c.Actions[0].URI)
		column := linebot.NewCarouselColumn(c.ThumbnailImageURL, c.Title, c.Text, act)
		columns = append(columns, column)
	}

	carTemplate := linebot.NewCarouselTemplate(columns...)

	return linebot.NewTemplateMessage(projectAltText, carTemplate)
}

func messageWithQuickReply(msg string, mConfig *MessageConfig) linebot.SendingMessage {
	items := []*linebot.QuickReplyButton{}

	for _, q := range mConfig.Welcome {
		ma := linebot.NewMessageAction(q.Label, q.Text)
		qrb := linebot.NewQuickReplyButton(q.ImageURL, ma)
		items = append(items, qrb)
	}

	quickReply := &linebot.QuickReplyItems{Items: items}

	return linebot.NewTextMessage(msg).WithQuickReplies(quickReply)
}
