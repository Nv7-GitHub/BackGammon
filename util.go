package main

import "math/rand"

// Roll rolls 2 dice
func Roll() []int {
	return []int{rand.Intn(5) + 1, rand.Intn(5) + 1}
}

// Dir gets the direction a player should be going based on the turn
func Dir(turn bool) int {
	if turn {
		return 1
	}
	return -1
}

// GetPieces gets the amount of pieces of a certain color given the turn and a triangle
func GetPieces(turn bool, tri Triangle) int {
	if turn {
		return tri.White
	}
	return tri.Black
}

// SetPieces sets the pieces given a turn and a triangle
func SetPieces(turn bool, tri Triangle, val int) Triangle {
	if turn {
		tri.White = val
	} else {
		tri.Black = val
	}
	return tri
}
