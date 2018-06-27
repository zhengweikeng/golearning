package main

import (
	"fmt"
	"golearning/pipeline"
)

func main() {
	p := pipeline.Merge(
		pipeline.InMemSort(pipeline.ArraySource(3, 2, 6, 7, 4)),
		pipeline.InMemSort(pipeline.ArraySource(7, 4, 0, 3, 2, 13, 8)),
	)
	for num := range p {
		fmt.Println(num)
	}
}
