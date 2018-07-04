package main

import "fmt"

func main() {
	b := 10
	a := &b
	*a++
	fmt.Println(*a, b)
}
