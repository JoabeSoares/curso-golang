package main

import "fmt"

func main() {
	defer joabe()
	soares()
	teixeira()
	souza()
}

func joabe() {
	fmt.Println("Joabe")
}

func soares() {
	fmt.Println("Soares")
}

func teixeira() {
	fmt.Println("Teixeira")
}

func souza() {
	fmt.Println("de Souza")
}
