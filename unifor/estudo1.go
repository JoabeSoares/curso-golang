package main

import "fmt"

func soma(a int, b int)int {
	var totalPreco = a + b
	return totalPreco
}

func mult(a int, b int)int {
	var multiPreco = a * b
	return multiPreco
}

func main()  {
	a, b := 40, 60
	totalPreco := soma(a, b)
	multiPreco := mult(a, b)
	fmt.Println(totalPreco)
	fmt.Println(multiPreco)
}

