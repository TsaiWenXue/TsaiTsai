package main

import (
	"log"
	"net/http"
	"os"

	"github.com/TsaiWenXue/TsaiTsai/src"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	// Init config
	mc, err := src.InitMessageConfig()
	if err != nil {
		panic(err)
	}

	// Init news scheduler
	src.Scheduler = src.InitNewsScheduler(mc)
	go src.Scheduler.PopNewsChan()
	go src.Scheduler.PopTicker()
	src.Scheduler.RefreshNews()

	// Init line bot
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	tth := &tsaitsaiHandler{
		bot: bot,
		mc:  mc,
	}
	
	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", tth.handleBotRequest)
	http.HandleFunc("/web-img", tth.webImgRequest)

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}

