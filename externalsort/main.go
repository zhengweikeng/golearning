package main

import (
	"bufio"
	"fmt"
	"golearning/pipeline"
	"os"
	"strconv"
)

func main() {
	// p := createPipeline("large.in", 800000000, 4)
	// writeToFile(p, "large.out")
	// printFile("large.out")

	p := createNetworkPipeline("small.in", 512, 4)
	writeToFile(p, "small.out")
	printFile("small.out")
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.ReaderSource(file, -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count > 100 {
			break
		}
	}
}

func createPipeline(
	filename string,
	filesize, chunkCount int) <-chan int {
	chunkSize := filesize / chunkCount
	pipeline.Init()
	sortResult := []<-chan int{}

	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)

		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)
		sortResult = append(sortResult, pipeline.InMemSort(source))
	}

	return pipeline.MergeN(sortResult...)
}

func writeToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriterSink(writer, p)
}

func createNetworkPipeline(
	filename string,
	filesize, chunkCount int) <-chan int {
	chunkSize := filesize / chunkCount
	pipeline.Init()
	sortAddr := []string{}

	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)

		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)
		addr := ":" + strconv.Itoa(7000+i)
		pipeline.NetworkSink(addr, pipeline.InMemSort(source))
		sortAddr = append(sortAddr, addr)
	}

	sortResult := []<-chan int{}
	for _, addr := range sortAddr {
		sortResult = append(sortResult, pipeline.NetworkSource(addr))
	}

	return pipeline.MergeN(sortResult...)
}
