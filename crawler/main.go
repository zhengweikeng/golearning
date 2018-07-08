package main

import (
	"golearning/crawler/engine"
	"golearning/crawler/schedule"
	"golearning/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Schedule:    &schedule.QueuedScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})
}
