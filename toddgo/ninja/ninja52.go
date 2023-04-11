package main

import "fmt"

type person struct {
	first    string
	last     string
	favorite []string
}

func main() {
	p1 := person{
		first:    "Joabe",
		last:     "Soares",
		favorite: []string{"Flocos", "Chocolate"},
	}
	p2 := person{
		first:    "Caroline",
		last:     "Cleophas",
		favorite: []string{"Morango", "Creme Brulet"},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(m)

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favorite {
			fmt.Println(i, val)
		}
		fmt.Println("------")
	}

}
