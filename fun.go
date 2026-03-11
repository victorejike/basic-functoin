package main

import (
	"fmt"
	"strconv"
)

func main() {
	hex := "1A"
	num, _ := strconv.ParseInt(hex, 16, 64)
	fmt.Println(num)
}