package main

import "fmt"

func soma(a int, b int) {
	var totalPreco = a + b
	fmt.Println(totalPreco)
}

func mult(a int, b int) {
	var multiPreco = a * b
	fmt.Println(multiPreco)
}

func main() {
	a, b := 40, 60
	soma(a, b)
	mult(a, b)

}