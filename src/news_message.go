package src

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

func newsTemplateMessage(mc *MessageConfig) linebot.SendingMessage {
	Scheduler.mu.Lock()
	defer Scheduler.mu.Unlock()

	carCont := &linebot.CarouselContainer{Type: linebot.FlexContainerTypeCarousel}
	for _, n := range Scheduler.newsMap {
		bubCont := newBubbleContainer(n)
		for i := 0; i < len(n.title); i++ {
			bubCont.Body.Contents = append(bubCont.Body.Contents, &linebot.TextComponent{
				Type:   linebot.FlexComponentTypeText,
				Text:   n.title[i],
				Size:   linebot.FlexTextSizeTypeSm,
				Weight: linebot.FlexTextWeightTypeBold,
				Action: &linebot.URIAction{
					Label: n.area,
					URI:   n.link[i],
				},
			})

			bubCont.Body.Contents = append(bubCont.Body.Contents, &linebot.SeparatorComponent{
				Type:   linebot.FlexComponentTypeSeparator,
				Margin: linebot.FlexComponentMarginTypeSm,
				Color:  gray,
			})

		}
		carCont.Contents = append(carCont.Contents, bubCont)
	}

	return linebot.NewFlexMessage(newsAltText, carCont)
}

func newBubbleContainer(n *cnnNews) *linebot.BubbleContainer {
	htextCon := &linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Text:   n.area,
		Size:   linebot.FlexTextSizeTypeXl,
		Align:  linebot.FlexComponentAlignTypeCenter,
		Weight: linebot.FlexTextWeightTypeBold,
		Color:  white,
	}
	if n.areaLink != "" {
		htextCon.Action = &linebot.URIAction{
			Label: n.area,
			URI:   n.areaLink,
		}
	}

	header := &linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeVertical,
		Contents: []linebot.FlexComponent{htextCon},
	}

	hero := &linebot.ImageComponent{
		Type:        linebot.FlexComponentTypeImage,
		Size:        linebot.FlexImageSizeTypeFull,
		AspectRatio: linebot.FlexImageAspectRatioType20to13,
		AspectMode:  linebot.FlexImageAspectModeTypeFit,
		URL:         n.imagePath,
	}
	body := &linebot.BoxComponent{
		Type:   linebot.FlexComponentTypeBox,
		Layout: linebot.FlexBoxLayoutTypeVertical,
	}
	footer := &linebot.BoxComponent{
		Type:   linebot.FlexComponentTypeBox,
		Layout: linebot.FlexBoxLayoutTypeVertical,
		Contents: []linebot.FlexComponent{
			&linebot.FillerComponent{Type: linebot.FlexComponentTypeFiller},
		},
	}
	style := &linebot.BubbleStyle{
		Header: &linebot.BlockStyle{BackgroundColor: black},
		Footer: &linebot.BlockStyle{BackgroundColor: white},
	}

	return &linebot.BubbleContainer{
		Type:      linebot.FlexContainerTypeBubble,
		Direction: linebot.FlexBubbleDirectionTypeLTR,
		Header:    header,
		Hero:      hero,
		Body:      body,
		Footer:    footer,
		Styles:    style,
	}
}
