package main

import (
	"golearning/crawler_distributed/rpcsupport"
	"golearning/crawler_distributed/worker"
	"log"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc(":9000", worker.CrawlService{}))
}
