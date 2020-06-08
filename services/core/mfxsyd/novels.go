package mfxsyd

import (
	"crypto/md5"
	"fmt"
	"log"
	"net"
	"net/http"
	"regexp"
	"strings"
	"time"

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
	crawler := colly.NewCollector()
	crawler.WithTransport(&http.Transport{
		Proxy: core.GetProxy,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	})
	return Novels{
		//在此设置收集器配置
		crawler: crawler,
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
	this.crawler.OnRequest(func(req *colly.Request) {
		log.Println("onrequest!")
	})

	this.crawler.OnError(func(resp *colly.Response, err error) {
		log.Println(err)
	})
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
				Sources: map[string]string{data.Sources[0]: elem.Attr("href")},
			}
			menu = append(menu, row)
		},
	)

	this.crawler.OnScraped(func(c *colly.Response) {
		var hash [16]byte = md5.Sum([]byte(data.Name))
		data.ID = fmt.Sprintf("%x", hash)[:16]

		for i, _ := range menu {
			key := md5.Sum([]byte(strings.Join([]string{data.Name, menu[i].Name, string(i)}, "")))
			menu[i].Sequence = i
			menu[i].ID = fmt.Sprintf("%x", key)[:16]
			menu[i].Novel = data.ID
		}

		// 更新小说数据
		nov := data
		if res, err := services.DB.Get(&nov); !res {
			if err == nil {
				services.DB.InsertOne(&data)
			}
		} else {
			i := 0
			for ; i < len(nov.Sources); i++ {
				if nov.Sources[i] == data.Sources[0] { // 当前网站在不在之前的源站列表里
					break // 在不做处理
				}
			}
			if i >= len(nov.Sources) { // 不在，只更新source字段
				nov.Sources = append(data.Sources, nov.Sources...)
				services.DB.Update(&nov, &models.Novel{
					ID: data.ID,
				})
			}
		}

		// 更新章节数据
		catelog := []models.Chapter{}                                                                                        // 数据库数据
		if res, err := services.DB.Where("novel_id = ?", data.ID).OrderBy("sequence asc").FindAndCount(&catelog); res <= 0 { //出错或数据库无数据
			if err == nil {
				services.DB.Insert(&menu)
			}
		} else {
			for i := range catelog { // catelog已有数据更新
				if i < len(menu) {
					catelog[i].Sources[data.Sources[0]] = menu[i].Sources[data.Sources[0]]
				}
			}
			services.DB.Update(&catelog)
			services.DB.Insert(menu[len(catelog):])
		}
	})

	this.crawler.Visit(url)
}
