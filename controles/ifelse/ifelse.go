package main

import "fmt"

func imprimirresultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com a nota", nota)
	} else {
		fmt.Println("Reprovado com a nota", nota)
	}
}

func main() {
	imprimirresultado(2.9)	
}