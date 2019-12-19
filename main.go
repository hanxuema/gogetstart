package main

import (
	"fmt"
)

const (
	first  = iota //0
	second = iota //1
	third  = iota //2
	fifth         //3
)

const (
	forth = iota //0 reset
	six          //1
)

func main() {
	const c int = 3
	fmt.Println(c)
	fmt.Println(first, second, third, fifth, forth, six)
	arraymain()
	mapmain()
	structmain()
}
