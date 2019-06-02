package main

import (
	"log"
	"net/http"
	"os"

	"github.com/TsaiWenXue/TsaiTsai/src"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	mConfig, err := src.InitMessage()
	if err != nil {
		panic(err)
	}

	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(http.StatusBadGateway)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		for _, event := range events {
			if err := src.HandleEvent(bot, event, mConfig); err != nil {
				log.Println(err)
			}
		}
	})

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
