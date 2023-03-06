package src

import (
	"fmt"
	"testing"
)

func TestBf(t *testing.T) {
	s := "bacbababaabcbab"
	p := "abababca"

	fmt.Printf("bf %s match %s index is: %+v\n", p, s, bf(s, p))
	// t.Logf("%s match %s index is: %+v\n", p, s, bf(s, p))
}

func TestGetNext(t *testing.T) {
	p := "abababca"
	fmt.Printf("getNext %s next is: %+v\n", p, getNext(p))
	// t.Logf("%s next is: %+v\n", p, next)
}

func TestKmp(t *testing.T) {
	s := "bacbababaabcbab"
	p := "abababca"

	fmt.Printf("kmp %s match %s index is: %+v\n", p, s, kmp(s, p))
	// t.Logf("%s match %s index is: %+v\n", p, s, kmp(s, p))
}
