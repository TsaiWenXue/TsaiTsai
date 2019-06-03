package src

import (
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

func newsTemplateMessage(mc *MessageConfig) linebot.SendingMessage {
	Scheduler.mu.Lock()
	defer Scheduler.mu.Unlock()
	now := time.Now().UTC()
	carCont := &linebot.CarouselContainer{Type: linebot.FlexContainerTypeCarousel}
	for area, m := range Scheduler.newsMap {
		bubCont := newBubbleContainer(area)
		for k, n := range m {
			// if a news exists too long, delete it
			if now.Sub(n.effectTime) > time.Duration(mc.News.EffectTime)*time.Hour {
				delete(m, k)
			}
			if n.imagePath != "" {
				bubCont.Hero = &linebot.ImageComponent{
					Type:        linebot.FlexComponentTypeImage,
					URL:         n.imagePath,
					Size:        linebot.FlexImageSizeTypeFull,
					AspectRatio: linebot.FlexImageAspectRatioType20to13,
					AspectMode:  linebot.FlexImageAspectModeTypeFit,
					Action: &linebot.URIAction{
						Label: n.id,
						URI:   n.link,
					},
				}
			} else {
				bubCont.Body.Contents = append(bubCont.Body.Contents, &linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   n.title,
					Size:   linebot.FlexTextSizeTypeLg,
					Weight: linebot.FlexTextWeightTypeBold,
					Action: &linebot.URIAction{
						Label: n.id,
						URI:   n.link,
					},
				})
				bubCont.Body.Contents = append(bubCont.Body.Contents, &linebot.SeparatorComponent{
					Type:   linebot.FlexComponentTypeSeparator,
					Margin: linebot.FlexComponentMarginTypeXs,
					Color:  gray,
				})
			}

		}
		carCont.Contents = append(carCont.Contents, bubCont)
	}

	return linebot.NewFlexMessage(newsAltText, carCont)
}

func newBubbleContainer(area string) *linebot.BubbleContainer {
	header := &linebot.BoxComponent{
		Type:   linebot.FlexComponentTypeBox,
		Layout: linebot.FlexBoxLayoutTypeVertical,
		Contents: []linebot.FlexComponent{
			&linebot.TextComponent{
				Type:   linebot.FlexComponentTypeText,
				Text:   area,
				Size:   linebot.FlexTextSizeTypeXl,
				Align:  linebot.FlexComponentAlignTypeCenter,
				Weight: linebot.FlexTextWeightTypeBold,
				Color:  white,
			},
		},
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
		Header: &linebot.BlockStyle{BackgroundColor: white},
		Footer: &linebot.BlockStyle{BackgroundColor: black},
	}
	return &linebot.BubbleContainer{
		Type:      linebot.FlexContainerTypeBubble,
		Direction: linebot.FlexBubbleDirectionTypeLTR,
		Header:    header,
		Body:      body,
		Footer:    footer,
		Styles:    style,
	}
}
