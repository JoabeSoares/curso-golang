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

	x[`Fleming_Ian`] = []string{`steaks`, `espionage`, `cigars`}

	for k, v := range x {
		fmt.Println("This is the record for ", k)
		for i, v2 := range v {
			fmt.Println("\t", i ,v2)
		}
	}
	fmt.Println(x)
}