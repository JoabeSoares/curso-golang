package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I'm", s.first, s.last)
}

func main() {
	s1 := secretAgent{
		person: person{
			"Caroline",
			"Cleophas",
		},
		ltk: true,
	}
	s2 := secretAgent{
		person: person{
			"Joabe",
			"Soares",
		},
		ltk: false,
	}
	fmt.Println(s1)
	s1.speak()
	fmt.Println(s2)
	s2.speak()
}
