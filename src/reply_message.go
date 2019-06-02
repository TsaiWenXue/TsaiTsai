package src

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

// WelcomeMessage is message that TsaiTsai will reply when making a new friend.
func WelcomeMessage(msg *MessageConfig) linebot.SendingMessage {
	items := []*linebot.QuickReplyButton{}

	for _, q := range msg.Welcome {
		ma := linebot.NewMessageAction(q.Label, q.Text)
		qrb := linebot.NewQuickReplyButton(q.ImageURL, ma)
		items = append(items, qrb)
	}

	quickReply := &linebot.QuickReplyItems{Items: items}

	return linebot.NewTextMessage(string(welcom)).WithQuickReplies(quickReply)
}
