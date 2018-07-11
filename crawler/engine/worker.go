package engine

import (
	"golearning/crawler/fetcher"
	"log"
)

func Worker(r Request) (ParseResult, error) {
	//log.Printf("Fetching %s", r.Url)

	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.Parser.Parse(body, r.Url), nil
}
