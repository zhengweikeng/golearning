package main

import (
	"golearning/crawler/engine"
	"golearning/crawler/schedule"
	"golearning/crawler/zhenai/parser"

	"golearning/crawler_distributed/persist/client"
)

func main() {
	itemChan, err := client.ItemSaver(":1234")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Schedule:    &schedule.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
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
