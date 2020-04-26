package core

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

type Collector interface {
	Reducer(url string)
}

type Classifier struct {
	model map[*regexp.Regexp]Collector
}

var classifier Classifier = Classifier{model: make(map[*regexp.Regexp]Collector, 10)}

func GetClassifier() *Classifier {
	return &classifier
}

func (this *Classifier) RegisterClassfier(regx *regexp.Regexp, c Collector) {
	this.model[regx] = c
}

func (this *Classifier) GetCollector(url string) Collector {
	for k, v := range this.model {
		if k.MatchString(url) {
			return v
		}
	}
	return nil
}

func GetProxy(*http.Request) (u *url.URL, err error) {
	data := struct {
		Proxy      string `json:"proxy"`
		FailCount  int    `json:"fail_count"`
		Region     string `json:"region"`
		Type       string `json:"type"`
		Source     string `json:"source"`
		ChankCount int    `json:"check_count"`
		LastStatus int    `json:"last_status"`
		LastTime   string `json:"last_time"`
	}{}
	u = &url.URL{}
	err = nil
	if resp, err := http.Get("http://127.0.0.1:5010/get"); err != nil {
		err = errors.New("get proxy failure!")
	} else {
		defer resp.Body.Close()
		str, err := ioutil.ReadAll(resp.Body)
		if err = json.Unmarshal(str, &data); err != nil {
			err = errors.New("proxy data decode error!")
		}
	}

	u.Host = data.Proxy
	if data.Type == "" {
		u.Scheme = "http"
	} else {
		u.Scheme = data.Type
	}
	return
}
