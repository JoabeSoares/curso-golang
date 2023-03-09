package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var numbers [6]int
	var count int

	for count < 6 {
		num := rand.Intn(60) + 1
		if !contains(numbers, num) {
			numbers[count] = num
			count++
		}
	}
	
	sort.Ints(numbers[:])
	fmt.Printf("Os 6 números sorteados foram: %d ", numbers)
}

func contains(numbers [6]int, num int) bool {
	for _, n := range numbers {
		if n == num {
			return true
		}
	}
	return false
}


