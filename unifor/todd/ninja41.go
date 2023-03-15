package main

import "fmt"

func main()  {
	x := [5]string{"Jow", "Sáskia", "Cleide", "Eliúde", "Ubiratan"}
	for i, v := range x {
		fmt.Println(i, v)	
	}
	fmt.Printf("%T\n", x)
}