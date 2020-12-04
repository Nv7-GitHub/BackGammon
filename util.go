package main

import "math/rand"

// Roll rolls 2 dice
func Roll() [2]int {
	return [2]int{rand.Intn(5) + 1, rand.Intn(5) + 1}
}
