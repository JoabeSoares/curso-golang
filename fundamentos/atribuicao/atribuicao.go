package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	x, y := 1, 2
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)
}