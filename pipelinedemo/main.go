package main

import (
	"fmt"
	"goDemo/pipeline"
)

func main() {
	p := pipeline.InMemSort(pipeline.ArraySource(3, 2, 6, 7, 4))
	for num := range p {
		fmt.Println(num)
	}
}
