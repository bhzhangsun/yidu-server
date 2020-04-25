package core

import "regexp"

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
