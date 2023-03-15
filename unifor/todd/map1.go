package main

import "fmt"

func main()  {
	m := map[string]int{
		"James":32,
		"Miss Moneypenny":27,
		
	}
	fmt.Println(m)
	fmt.Println(m["Miss Moneypenny"])
	fmt.Println(m["Joabe"])

	v, ok := m["Joabe"]
	fmt.Println(v)
	fmt.Println(ok)

	m["Joabe"] = 33

	if v, ok := m["James"]; ok {
		fmt.Println("THIS IS THE IF PRINT", v)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	xi := []int{1, 2, 3, 4, 12, 42}

	for i, v := range xi {
		fmt.Println(i, v)
	}

	delete(m, "James")
	fmt.Println(m)
}