package src

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	s := "babad"
	fmt.Printf("LongestPalindrome of %s is %s\n", s, longestPalindrome(s))
}
