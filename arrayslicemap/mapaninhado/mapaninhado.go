package main

import (
	"fmt"
)

func main() {
	funcsPorLetra := map[string]map[string]float64 {
		"G": {
			"Gabriela Silva": 15000.00,
			"Guga Pereira": 12000.00,
		},
		"J": {
			"José João": 11321.22,
		},
		"P": {
			"Pedro Júnior": 12344.00,
		},
	}
	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}