package core

import "regexp"

type Collector interface {
	Structure(uri string)
	Store()
}

type Classifier struct {
	model map[*regexp.Regexp] Collector
}

func RegisterClassfier(regx *regexp.Regexp, c Collector) {

}