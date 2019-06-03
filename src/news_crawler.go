package src

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly"
)

// NewsScheduler used to timely refresh news
type NewsScheduler struct {
	mu         *sync.Mutex
	newsMap    map[string]map[string]*cnnNews
	newsChan   chan *cnnNews
	timeTicker *time.Ticker
}

type cnnNews struct {
	id         string
	area       string
	imagePath  string
	title      string
	link       string
	effectTime time.Time
}

// InitNewsScheduler init NewsScheduler
func InitNewsScheduler(mc *MessageConfig) *NewsScheduler {
	return &NewsScheduler{
		mu:         &sync.Mutex{},
		newsMap:    map[string]map[string]*cnnNews{},
		newsChan:   make(chan *cnnNews, mc.News.ChanBuffer),
		timeTicker: time.NewTicker(time.Duration(mc.News.RefreshPeriod) * time.Minute),
	}
}

// AddToQueue adds a news to newsChan
func (ns *NewsScheduler) AddToQueue(n *cnnNews) {
	ns.newsChan <- n
}

// PopNewsChan call updateNews when news is in the newsChan.
func (ns *NewsScheduler) PopNewsChan() {
	for n := range ns.newsChan {
		if err := ns.updateNewsMap(n); err != nil {
			log.Println(err)
		}
	}
}

func (ns *NewsScheduler) updateNewsMap(n *cnnNews) error {
	ns.mu.Lock()
	defer ns.mu.Unlock()
	if n.title == "" || n.area == "" || n.link == "" {
		return fmt.Errorf("cnnNews field empty")
	}

	if _, ok := ns.newsMap[n.area]; !ok {
		ns.newsMap[n.area] = map[string]*cnnNews{n.id: n}
	}
	ns.newsMap[n.area][n.id] = n

	log.Println(ns.newsMap)
	return nil
}

// PopTicker call RefreshNews a period of time.
func (ns *NewsScheduler) PopTicker() {
	for range ns.timeTicker.C {
		ns.RefreshNews()
	}
}

// RefreshNews crawl news from cnn
func (ns *NewsScheduler) RefreshNews() {
	c := colly.NewCollector()

	c.OnHTML("#world-zone-2 .zn__containers", func(e *colly.HTMLElement) {
		e.ForEach("div ul li article", ns.newCNNNews)
	})
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	c.Visit(cnnDomain + "/world")
}

func (ns *NewsScheduler) newCNNNews(i int, h *colly.HTMLElement) {
	var (
		image, title, link, area string
	)
	re := regexp.MustCompile(`(?P<one>src=")(?P<two>.*)(?P<three>")`)
	imgText := h.ChildText("a noscript")
	matchStrings := re.FindStringSubmatch(imgText)
	if len(matchStrings) >= 3 {
		image = matchStrings[2]
	}
	// link=/2019/06/01/asia/bts-kpop-us-intl/index.html
	link = h.ChildAttr("a", "href")
	title = h.ChildText("h3")

	if len(link) > 0 {
		tmpArea := link[12:]
		area = tmpArea[:strings.Index(tmpArea, "/")]
	}

	ns.AddToQueue(&cnnNews{
		id:         fmt.Sprintf("%s-%d", area, i),
		area:       strings.ToUpper(string(area[0])) + area[1:],
		title:      title,
		link:       cnnDomain + link,
		imagePath:  https + image,
		effectTime: time.Now().UTC(),
	})
}
