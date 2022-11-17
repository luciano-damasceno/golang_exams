package main

import (
	"fmt"
	"strings"
)

/*
 * The objective of this exercise is to find the length of the longest substring with unique characters of a given string.
 * Youtube link: https://www.youtube.com/watch?v=Pe64ee24MLY&t=492s
 */
func main() {
	s := "pwwkewaaaa"
	l := findLongestSubstringLength(s)
	fmt.Printf("The longest substring with unique characters is: %v.\n", l)
}

func findLongestSubstringLength(s string) int {
	sub := ""
	sublen := 0
	for leftIndex := 0; leftIndex < len(s); leftIndex++ {
		rightIndex := leftIndex
		for ; rightIndex < len(s); rightIndex++ {
			if strings.Count(sub, s[rightIndex:rightIndex+1]) == 0 {
				sub = s[leftIndex : rightIndex+1]
				if sublen < len(sub) {
					sublen = len(sub)
				}
			} else {
				sub = ""
				break
			}
		}
	}
	return sublen
}
