package main

import (
	"golearning/crawler_distributed/persist"
	"golearning/crawler_distributed/rpcsupport"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func main() {
	log.Fatal(serveRpc(":1234", "dating_profile"))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
