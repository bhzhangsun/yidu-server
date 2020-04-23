package mfxsyd

import (
	"log"
	"regexp"

	"github.com/gocolly/colly"
	"yidu.4se.tech/models"
	"yidu.4se.tech/services"
	"yidu.4se.tech/services/core"
)

type Novels struct {
	crawler *colly.Collector
	data    models.Novel
}

const PATTERN string = `^http://www.mianfeixiaoshuoyueduwang.com/book/.+`

func NewNovels() Novels {
	return Novels{
		crawler: colly.NewCollector(),
		data:    models.Novel{},
	}
}

func (this *Novels) RegisterToClassifier() {
	c := core.GetClassifier()
	if regx, err := regexp.Compile(PATTERN); err == nil {
		c.RegisterClassfier(regx, this)
	}
}

func (this *Novels) Resolve(url string) {
	// c.configuration
	this.selector()
	this.crawler.Visit(url)
	this.crawler.Wait()
}

func (this *Novels) Store() {
	// DB 存储数据
	services.DB.InsertOne(&this.data)
}

// private
func (this *Novels) selector() {
	this.crawler.OnResponse(func(resp *colly.Response) {
		log.Println(string(resp.Body))
	})

	this.crawler.OnHTML(
		"body > div.wrapper > div:nth-child(2) > div > h3 > span",
		func(elem *colly.HTMLElement) {
			log.Println(elem.Attr("itemprop"), ":", elem.Text)
		},
	)

	this.crawler.OnHTML(
		"body > div.wrapper > div:nth-child(2) > div > p",
		func(elem *colly.HTMLElement) {
		},
	)
}
