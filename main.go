package main

import (
	"fmt"
)

func main() {
	var first *string = new(string)
	*first = "archur" //first is the pointer
	fmt.Println(*first) //string s
}
