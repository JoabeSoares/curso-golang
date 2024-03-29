package main

import "math"

type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy Ponto) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - b.x)
	return
}

func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}