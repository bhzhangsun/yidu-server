package main

import (
	"fmt"

	"yidu.4se.cool/services/core/mfxsyd"
	"yidu.4se.cool/web"
)

func main() {
	novels := mfxsyd.NewNovels()
	novels.RegisterToClassifier()
	for i := 0; i < 2000; i++ {
		novels.Reducer(fmt.Sprintf("http://www.mianfeixiaoshuoyueduwang.com/book/%d", i))
	}
	web.Run()
}
