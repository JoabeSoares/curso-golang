package main

import (
	"fmt"
)

func main()  {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Hellooooo, James!"}
	fmt.Println(x, y)

	xy := [][]string{x, y}
	fmt.Println(xy)

	for i, xs := range xy {
		fmt.Println("Record: ", i)
		for j, val := range xs {
			fmt.Printf("\t Index position: %v \t Value: \t %v \n", j, val)
		}
	}
}