package src

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

func infoTemplateMessage(mc *MessageConfig) linebot.SendingMessage {
	return linebot.NewFlexMessage(infoAltText, infoBubbleContainer(mc))
}

func infoBubbleContainer(mc *MessageConfig) *linebot.BubbleContainer {
	headerTextCon := &linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Text:   dennyTsai,
		Size:   linebot.FlexTextSizeTypeXl,
		Align:  linebot.FlexComponentAlignTypeCenter,
		Weight: linebot.FlexTextWeightTypeBold,
		Color:  black,
	}

	header := &linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeVertical,
		Contents: []linebot.FlexComponent{headerTextCon},
	}

	hero := &linebot.ImageComponent{
		Type:            linebot.FlexComponentTypeImage,
		Size:            linebot.FlexImageSizeTypeFull,
		AspectRatio:     linebot.FlexImageAspectRatioType3to4,
		AspectMode:      linebot.FlexImageAspectModeTypeCover,
		BackgroundColor: white,
		URL:             mc.InfoTemplate.IMG,
	}

	charac := &linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Text:   characteristic,
		Size:   linebot.FlexTextSizeTypeLg,
		Align:  linebot.FlexComponentAlignTypeStart,
		Weight: linebot.FlexTextWeightTypeBold,
	}

	characContent := &linebot.TextComponent{
		Type:  linebot.FlexComponentTypeText,
		Color: waterBlue,
		Text:  mc.InfoTemplate.Characteristic,
		Wrap:  true,
	}

	habit := &linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Text:   habit,
		Size:   linebot.FlexTextSizeTypeLg,
		Align:  linebot.FlexComponentAlignTypeStart,
		Weight: linebot.FlexTextWeightTypeBold,
	}
	habitContent := &linebot.TextComponent{
		Type:  linebot.FlexComponentTypeText,
		Color: lightOrange,
		Text:  mc.InfoTemplate.Habit,
		Wrap:  true,
	}
	motto := &linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Text:   motto,
		Size:   linebot.FlexTextSizeTypeLg,
		Align:  linebot.FlexComponentAlignTypeStart,
		Weight: linebot.FlexTextWeightTypeBold,
	}
	mottoContent := &linebot.TextComponent{
		Type:  linebot.FlexComponentTypeText,
		Color: lightRed,
		Text:  mc.InfoTemplate.Motto,
		Wrap:  true,
	}
	sep := &linebot.SeparatorComponent{Type: linebot.FlexComponentTypeSeparator}

	body := &linebot.BoxComponent{
		Type:   linebot.FlexComponentTypeBox,
		Layout: linebot.FlexBoxLayoutTypeVertical,
		Contents: []linebot.FlexComponent{
			charac, sep, characContent,
			habit, sep, habitContent,
			motto, sep, mottoContent},
	}

	footer := &linebot.BoxComponent{
		Type:   linebot.FlexComponentTypeBox,
		Layout: linebot.FlexBoxLayoutTypeHorizontal,
		Contents: []linebot.FlexComponent{
			&linebot.FillerComponent{Type: linebot.FlexComponentTypeFiller},
		},
	}

	style := &linebot.BubbleStyle{
		Header: &linebot.BlockStyle{BackgroundColor: lightGreenBlue},
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
