package main

import "fmt"

func obterresultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado!"
	}
	return "Reprovado."
}

func main() {
	fmt.Println(obterresultado(6.2))
}