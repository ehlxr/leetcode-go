package main

import "fmt"

func main() {
	s := "Hello 你好"
	fmt.Println(s[2], s[6])
	rs := []rune(s)
	println(string(rs[6]))

}
