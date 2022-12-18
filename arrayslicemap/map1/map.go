package main

import "fmt"

func main() {
	aprovados := make(map[int]string)

	aprovados[122344555] = "Pedro"
	aprovados[545434567] = "Maria"
	aprovados[909876656] = "Joabe"
	
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[909876656])
	delete(aprovados, 909876656)
	fmt.Println(aprovados[909876656])
}