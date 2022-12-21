package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64 {
		"Joabe": 11900.00,
		"Carol": 30500.21,
		"Lara": 45233.12,
	}

	funcsESalarios["Rafael"] = 12433.90
	funcsESalarios["Cleide"] = 32122.32

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}