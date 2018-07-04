package main

import (
	"fmt"
	"unicode/utf8"
)

func runeDemo() {
	fmt.Println("==========runeDemo=========")
	str := "seed你好!"
	fmt.Printf("str len:%d\n", len(str))

	bytes := []byte(str)
	fmt.Println("range bytes:")
	for i, ch := range bytes {
		fmt.Printf("%d %X ", i, ch)
	}
	fmt.Println()

	fmt.Println("range str: ")
	for i, ch := range str {
		fmt.Printf("%d %X ", i, ch)
	}
	fmt.Println()

	fmt.Printf("rune count: %d\n", utf8.RuneCountInString(str))

	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		fmt.Printf("%c %d\n", ch, size)
		bytes = bytes[size:]
	}

	for i, ch := range []rune(str) {
		fmt.Printf("%d, %c", i, ch)
	}

	fmt.Println("\n")
}
