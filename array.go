package main

import (
	"fmt"
)

func arraymain() {
	var arr [3]int 
	fmt.Println(arr)
	arr[0] = 1
	arr[1] = 2 
	arr[2] = 3
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	slice := arr[:]
	fmt.Println(slice)
	slice = append(slice, 4,5,6)
	fmt.Println(slice)
	
	s2 := slice[1:]
	s3 := slice[:2]
	s4 := slice[1:2]

	fmt.Println(s2,s3,s4)
}
