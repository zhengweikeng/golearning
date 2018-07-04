package main

import "fmt"

/**
给定一个字符串，找出不含有重复字符的最长子串的长度。

示例：

给定 "abcabcbb" ，没有重复字符的最长子串是 "abc" ，那么长度就是3。

给定 "bbbbb" ，最长的子串就是 "b" ，长度是1。

给定 "pwwkew" ，最长子串是 "wke" ，长度是3。请注意答案必须是一个子串，"pwke" 是 子序列  而不是子串。
*/
func getLengthOfLongestSubstring(str string) (maxLength int, subStr string) {
	lastOccurred := make(map[byte]int)
	// 最长子串的起始位置
	start := 0

	for i, ch := range []byte(str) {
		lastI, ok := lastOccurred[ch]

		//fmt.Printf("i:%d lastI:%d start:%d maxLength:%d\n", i, lastI, start, maxLength)
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	fmt.Println(start, maxLength)
	if start <= maxLength {
		subStr = str[start : maxLength+start]
	} else {
		subStr = str[start-maxLength : start]
	}

	return
}

func main() {
	length, subStr := getLengthOfLongestSubstring("abcdbkefoi")
	fmt.Printf("abcdbkefoi:%d subStr:%s\n", length, subStr)
	//
	length, subStr = getLengthOfLongestSubstring("bbbbb")
	fmt.Printf("bbbbb:%d subStr:%s\n", length, subStr)

	length, subStr = getLengthOfLongestSubstring("pwwkew")
	fmt.Printf("pwwkew:%d subStr:%s\n", length, subStr)

	length, subStr = getLengthOfLongestSubstring("pwwkeew")
	fmt.Printf("pwwkeew:%d subStr:%s\n", length, subStr)

	length, subStr = getLengthOfLongestSubstring("")
	fmt.Printf("\"\":%d subStr:%s\n", length, subStr)
	//
	length, subStr = getLengthOfLongestSubstring("abcdefg")
	fmt.Printf("abcdefg:%d subStr:%s\n", length, subStr)
}
