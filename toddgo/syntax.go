package main

import "fmt"

func main() {

	carol("Joabe")
	soma(3, 10)
	x := mouse("Joabe", "Soares")
	y := false
	fmt.Println(x)
	fmt.Println(y)

}

func carol(s string) {
	fmt.Println("Hello,", s)
}

func soma(x, y int) {
	fmt.Println(x + y)
}

func mouse(fn string, ln string) (string bool) {
	a := fmt.Sprint(fn, ln, `, says "Hello"`)
	
	fmt.Println(b)
	return b, a

}
