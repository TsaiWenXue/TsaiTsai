package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

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

	go scheduleCrawl()

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

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}

func scheduleCrawl() {
	ticker := time.NewTicker(time.Duration(25) * time.Minute)
	body := []byte(`project=default&spider=nba`)
	for range ticker.C {
		log.Println("start crawl")
		if resp, err := http.Post("https://scrapy-nba.herokuapp.com/schedule.json", "application/x-www-form-urlencoded", bytes.NewBuffer(body)); err != nil {
			log.Println(err)
		} else {
			defer resp.Body.Close()
			r, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
			} else {
				log.Println((string(r)))
			}
		}

	}
}
