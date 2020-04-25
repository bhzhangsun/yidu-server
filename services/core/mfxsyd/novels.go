package mfxsyd

import (
	"crypto/md5"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
	"yidu.4se.tech/models"
	"yidu.4se.tech/services"
	"yidu.4se.tech/services/core"
)

type Novels struct {
	crawler *colly.Collector
}

const PATTERN string = `^http://www.mianfeixiaoshuoyueduwang.com/book/.+`

func NewNovels() Novels {
	return Novels{
		//在此设置收集器配置
		crawler: colly.NewCollector(),
	}
}

func (this *Novels) RegisterToClassifier() {
	c := core.GetClassifier()
	if regx, err := regexp.Compile(PATTERN); err == nil {
		c.RegisterClassfier(regx, this)
	}
}

func (this *Novels) Reducer(url string) {
	var data models.Novel
	var menu []models.Chapter = make([]models.Chapter, 0)
	this.crawler.OnResponse(func(resp *colly.Response) {
		data.Sources = []string{resp.Request.URL.String()}
		log.Println(resp.Request.URL.String())
	})

	this.crawler.OnHTML(
		"body > div.wrapper > div:nth-child(2) > div > img",
		func(elem *colly.HTMLElement) {
			data.Image = elem.Attr("src")
		},
	)

	this.crawler.OnHTML(
		"body > div.wrapper > div:nth-child(2) > div > h3 > span",
		func(elem *colly.HTMLElement) {
			data.Name = elem.Text
		},
	)

	this.crawler.OnHTML(
		"body > div.wrapper > div:nth-child(2) > div > p:nth-child(3) > span",
		func(elem *colly.HTMLElement) {
			data.Author = elem.Text
		},
	)
	this.crawler.OnHTML(
		"body > div.wrapper > div:nth-child(2) > div > p:nth-child(4) > span",
		func(elem *colly.HTMLElement) {
			data.Tags = []string{elem.Text}
		},
	)
	this.crawler.OnHTML(
		"body > div.wrapper > div:nth-child(2) > div > p:nth-child(5) > span",
		func(elem *colly.HTMLElement) {
			var state2int = map[string]int8{
				"完结": 0,
				"连载": 1,
			}
			data.State = state2int[elem.Text]
		},
	)
	this.crawler.OnHTML(
		"body > div.wrapper > div:nth-child(2) > p",
		func(elem *colly.HTMLElement) {
			data.Description = elem.Text
		},
	)

	this.crawler.OnHTML(
		"body > div.wrapper > div:nth-child(4) > ul > li > a",
		func(elem *colly.HTMLElement) {
			row := models.Chapter{
				Name:    elem.ChildText("span"),
				Sources: []string{elem.Attr("href")},
			}
			menu = append(menu, row)
		},
	)

	this.crawler.OnScraped(func(c *colly.Response) {
		var hash [16]byte = md5.Sum([]byte(data.Name))
		data.ID = fmt.Sprintf("%x", hash)[:16]

		for i, _ := range menu {
			key := md5.Sum([]byte(strings.Join([]string{data.Name, menu[i].Name, string(i)}, "")))
			menu[i].ID = fmt.Sprintf("%x", key)[:16]
			menu[i].Novel = data.ID
		}

		if _, err := services.DB.InsertOne(&data); err != nil {
			log.Println("insert novel error")
		}
		if _, err := services.DB.Insert(&menu); err != nil {
			log.Println("insert menu error")
		}
		result := models.Novel{ID: data.ID}
		if _, err := services.DB.Get(&result); err != nil {
			log.Println("get error")
		}

	})

	this.crawler.Visit(url)
}
