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
	s := "abccd  efghijabc"
	l := findLongestSubstringLength(s)
	fmt.Printf("The longest substring with unique characters is: %v.\n", l)
}

func findLongestSubstringLength(s string) int {
	sub := ""
	sublen := 0
	for i := 0; i < len(s); i++ {
		if strings.Count(sub, s[i:i+1]) == 0 {
			sub += s[i : i+1]
		} else {
			sub = s[i : i+1]
		}
		if sublen < len(sub) {
			sublen = len(sub)
		}
	}
	return sublen
}
