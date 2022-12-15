package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough
	case 9:
		return "Aluno nota: A!"
	case 8, 7:
		return "Aluno nota: B!"
	case 6, 5:
		return "Aluno nota: C!"
	case 4, 3:
		return "Aluno nota: D!"
	case 2, 1, 0:
		return "Aluno nota: E!"
	default:
		return "Nota inválida..."
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
}