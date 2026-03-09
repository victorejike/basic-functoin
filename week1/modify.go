package main

import "fmt"

func main() {
	//this was calling the slice from the index of o to 6
	price := []int{10, 20, 30, 40}

	fmt.Println(price[0])
	//to print out the second slice
	fmt.Println(price[2])
}
