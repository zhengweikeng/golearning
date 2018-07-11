package main

import (
	"golearning/crawler/engine"
	"golearning/crawler/schedule"
	"golearning/crawler/zhenai/parser"

	itemSaver "golearning/crawler_distributed/persist/client"
	worker "golearning/crawler_distributed/worker/client"
)

func main() {
	itemChan, err := itemSaver.ItemSaver(":1234")
	if err != nil {
		panic(err)
	}

	processor, err := worker.CreateProcessor()

	e := engine.ConcurrentEngine{
		Schedule:         &schedule.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})
}
