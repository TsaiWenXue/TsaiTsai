package main

import (
	"log"
	"net/http"

	"github.com/TsaiWenXue/TsaiTsai/src"
	"github.com/line/line-bot-sdk-go/linebot"
)

type tsaitsaiHandler struct {
	bot *linebot.Client
	mc  *src.MessageConfig
}

func (tth *tsaitsaiHandler) handleBotRequest(w http.ResponseWriter, req *http.Request) {
	events, err := tth.bot.ParseRequest(req)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(http.StatusBadGateway)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	for _, event := range events {
		if err := src.HandleEvent(tth.bot, event, tth.mc); err != nil {
			log.Println(err)
		}
	}
}
