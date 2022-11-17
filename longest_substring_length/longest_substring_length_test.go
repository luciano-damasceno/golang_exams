package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	s := "abcabcbb"
	l := findLongestSubstringLength(s)
	assert.Equal(t, 3, l, "The substring length must be 3.")
}

func TestExample2(t *testing.T) {
	s := "bbbbb"
	l := findLongestSubstringLength(s)
	assert.Equal(t, 1, l, "The substring length must be 1.")
}

func TestExample3(t *testing.T) {
	s := "pwwkew"
	l := findLongestSubstringLength(s)
	assert.Equal(t, 3, l, "The substring length must be 3.")
}
