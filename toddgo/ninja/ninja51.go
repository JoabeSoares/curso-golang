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
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favorite {
		fmt.Println(i, v)
	}
	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favorite {
		fmt.Println(i, v)
	}

}
