package main

import "fmt"

func main()  {
	intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, k := range intSlice {
		fmt.Println(i, k)
	}
	fmt.Printf("%T", intSlice)
}