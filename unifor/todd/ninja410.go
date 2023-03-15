package main

import (
	"fmt"
)

func main()  {
	x := map[string][]string{
		"Bond_James": []string{"Shaken, not stirred", "Martinis", "Woman"},
		"Moneypenny_ Miss": []string{"James Bond", "Literature", "Computer Science"},
		"No_Dr": []string{"Ice cream", "Being Evil", "Sunset"},
	}

	x["Joabe"] = []string{"carol", "michelle", "maura"} 

	for k, v := range x {
		fmt.Println("This is the record for ", k)
		for i, v2 := range v {
			fmt.Println("\t", i ,v2)
		}
		delete(x, "No_Dr")
	}

	fmt.Println(x)
}