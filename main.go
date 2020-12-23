package main

import (
	"fmt"

	"yidu.4se.cool/services/core/mfxsyd"
	"yidu.4se.cool/web"
)

func main() {
	novels := mfxsyd.NewNovels()
	novels.RegisterToClassifier()
	// for i := 0; i < 00; i++ {
	novels.Reducer(fmt.Sprintf("http://www.mianfeixiaoshuoyueduwang.com/book/%d", 1))
	// }
	web.Run()
}
