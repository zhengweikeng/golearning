package engine

import (
	"golearning/crawler/model"
	"log"
)

type ConcurrentEngine struct {
	Schedule    Schedule
	WorkerCount int
}

type Schedule interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Schedule.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Schedule.WorkerChan(), out, e.Schedule)
	}
	for _, r := range seeds {
		if isDuplicate(r.Url) {
			continue
		}
		e.Schedule.Submit(r)
	}

	profileCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			if _, ok := item.(model.Profile); ok {
				log.Printf("Got profile #%d: %v", profileCount, item)
				profileCount++
			}

		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Schedule.Submit(request)
		}
	}
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}

	visitedUrls[url] = true
	return false
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
