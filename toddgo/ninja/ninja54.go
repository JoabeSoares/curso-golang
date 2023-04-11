package main

import "fmt"

func main() {
	p1 := struct {
		first     string
		last      string
		age       int
		friends   map[string]int
		favDrinks []string
		degree    bool
	}{
		first: "Joabe",
		last:  "Soares",
		age:   33,
		friends: map[string]int{
			"Gustavo": 29,
			"Pedro":   30,
			"Matheus": 25,
		},
		favDrinks: []string{
			"Água",
			"Cerveja",
		},
		degree: true,
	}
	fmt.Println(p1)
	for k, v := range p1.friends {
		fmt.Println(k, v)
	}
	for i, v := range p1.favDrinks {
		fmt.Println(i, v)
	}
}
