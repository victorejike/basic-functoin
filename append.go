package main

import "fmt"

func main() {
	// i add append to my slice i now understand how to use append
	myslice := []int{1, 30, 40, 90, 100, 34}
	// add append to 6 to see what it will turn to be
	myslice = append(myslice, 9, 7)
	//to output my result
	fmt.Println(myslice)

}
