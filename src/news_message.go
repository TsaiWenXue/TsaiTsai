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
			if bubCont.Hero.URL == "" {
				bubCont.Hero.URL = n.imagePath
			}

			bubCont.Body.Contents = append(bubCont.Body.Contents, &linebot.TextComponent{
				Type:   linebot.FlexComponentTypeText,
				Text:   n.title,
				Size:   linebot.FlexTextSizeTypeSm,
				Weight: linebot.FlexTextWeightTypeBold,
				Action: &linebot.URIAction{
					Label: n.id,
					URI:   n.link,
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
	hero := &linebot.ImageComponent{
		Type:        linebot.FlexComponentTypeImage,
		Size:        linebot.FlexImageSizeTypeFull,
		AspectRatio: linebot.FlexImageAspectRatioType20to13,
		AspectMode:  linebot.FlexImageAspectModeTypeFit,
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
