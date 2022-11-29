package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LongestSubstringLenght(t *testing.T) {
	tableTests := []struct {
		testString string
		expected   int
		msg        string
	}{
		{"abcabcbb", 3, "The substring length must be 3."},
		{"bbbbb", 1, "The substring length must be 1."},
		{"pwwkew", 3, "The substring length must be 3."},
	}
	for _, tt := range tableTests {
		r := findLongestSubstringLength(tt.testString)
		assert.Equal(t, tt.expected, r, tt.msg)
	}
}
