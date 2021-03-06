package src

import (
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
	newsMap    map[string]*cnnNews
	newsChan   chan *cnnNews
	timeTicker *time.Ticker
}

type cnnNews struct {
	area      string
	areaLink  string
	imagePath string
	title     []string
	link      []string
}

// InitNewsScheduler init NewsScheduler
func InitNewsScheduler(mc *MessageConfig) *NewsScheduler {
	return &NewsScheduler{
		mu:         &sync.Mutex{},
		newsMap:    map[string]*cnnNews{},
		newsChan:   make(chan *cnnNews, mc.News.ChanBuffer),
		timeTicker: time.NewTicker(time.Duration(mc.News.RefreshPeriod) * time.Minute),
	}
}

// AddToQueue adds a news to news channel.
func (ns *NewsScheduler) AddToQueue(n *cnnNews) {
	ns.newsChan <- n
}

// PopNewsChan update news map from news in the news channel.
func (ns *NewsScheduler) PopNewsChan() {
	for n := range ns.newsChan {
		ns.updateNewsMap(n)
	}
}

func (ns *NewsScheduler) updateNewsMap(n *cnnNews) {
	ns.mu.Lock()
	defer ns.mu.Unlock()

	if _, ok := ns.newsMap[n.area]; !ok {
		ns.newsMap[n.area] = n
	}

	ns.newsMap[n.area] = n
}

// PopTicker timely crawl news.
func (ns *NewsScheduler) PopTicker() {
	for range ns.timeTicker.C {
		ns.RefreshNews()
	}
}

// RefreshNews crawl news from cnn.
func (ns *NewsScheduler) RefreshNews() {
	c := colly.NewCollector()

	c.OnHTML("#world-zone-2 .zn__containers", func(e *colly.HTMLElement) {
		e.ForEach("div ul", ns.newCNNNews)
	})
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	c.Visit(cnnDomain + "/world")
}

func (ns *NewsScheduler) newCNNNews(i int, h *colly.HTMLElement) {
	var (
		image, area, areaLink string
		class, titles, links  []string
	)
	re := regexp.MustCompile(`(?P<one>src=")(?P<two>.*)(?P<three>")`)
	imgText := h.ChildText("a noscript")
	matchStrings := re.FindStringSubmatch(imgText)
	if len(matchStrings) >= 3 {
		image = https + matchStrings[2]
	}

	area = h.ChildText("a h2")
	areaLink = h.ChildAttr("a", "href")
	if areaLink != "" && len(areaLink) < len(cnnDomain) {
		areaLink = cnnDomain + areaLink
	}

	class = h.ChildAttrs("li article", "class")
	for _, c := range class {
		c = strings.Replace(c, " ", ".", -1)
		title := h.ChildText("." + c + " h3")
		link := h.ChildAttr("."+c+" h3 a", "href")

		if link != "" && title != "" {
			titles = append(titles, title)
			links = append(links, cnnDomain+link)
		}
	}

	if image == "" {
		image = cnnDefaultImg
	}
	if len(links) == len(titles) {
		ns.AddToQueue(&cnnNews{
			area:      area,
			areaLink:  areaLink,
			title:     titles,
			link:      links,
			imagePath: image,
		})
	}
}
