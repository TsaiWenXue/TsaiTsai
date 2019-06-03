package src

import (
	"math/rand"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

var (
	// Scheduler is the news scheduler
	Scheduler *NewsScheduler
)

// HandleEvent deal with event sent from users.
func HandleEvent(bot *linebot.Client, event *linebot.Event, mc *MessageConfig) error {
	switch event.Type {
	case linebot.EventTypeMessage:
		switch message := event.Message.(type) {
		case *linebot.TextMessage:
			if _, err := bot.ReplyMessage(event.ReplyToken, handleTextMessage(mc, message.Text)).Do(); err != nil {
				return err
			}
		}
	case linebot.EventTypeJoin:
		if _, err := bot.ReplyMessage(event.ReplyToken, messageWithQuickReply(welcome, mc)).Do(); err != nil {
			return err
		}
	}

	return nil
}

func handleTextMessage(mc *MessageConfig, text string) linebot.SendingMessage {
	switch specialWord(strings.ToLower(text)) {
	case handsomePhoto:
		return randHandsomePhoto(mc)
	case help:
		return messageWithQuickReply(string(helpReply), mc)
	case project:
		return projectCarousel(mc)
	case hello, hi:
		pkgID := stickersPackageMap[rand.Intn(len(stickersPackageMap))]
		stickerID := stickersMap[pkgID][rand.Intn(len(stickersMap[pkgID]))]
		return linebot.NewStickerMessage(pkgID, stickerID)
	case "news":
		return newsTemplateMessage(mc)
	default:
		return messageWithQuickReply(defaultReply, mc)
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

func messageWithQuickReply(msg string, mc *MessageConfig) linebot.SendingMessage {
	items := []*linebot.QuickReplyButton{}

	for _, q := range mc.Welcome {
		ma := linebot.NewMessageAction(q.Label, q.Text)
		qrb := linebot.NewQuickReplyButton(q.ImageURL, ma)
		items = append(items, qrb)
	}

	quickReply := &linebot.QuickReplyItems{Items: items}

	return linebot.NewTextMessage(msg).WithQuickReplies(quickReply)
}
