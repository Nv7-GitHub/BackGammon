package main

import (
	"fmt"
	"strconv"
)

// Triangle stores the amount of white and black pieces
type Triangle struct {
	White int
	Black int
}

// Board represents the board
type Board struct {
	Tris     [24]Triangle
	Bar      Triangle
	Finished Triangle
}

// NewBoard creates a board
func NewBoard() Board {
	b := Board{}
	b.Tris[0] = Triangle{White: 2}
	b.Tris[5] = Triangle{Black: 5}
	b.Tris[7] = Triangle{Black: 3}
	b.Tris[11] = Triangle{White: 5}
	b.Tris[12] = Triangle{Black: 5}
	b.Tris[16] = Triangle{White: 3}
	b.Tris[18] = Triangle{White: 5}
	b.Tris[23] = Triangle{Black: 3}
	return b
}

// Print prints the board
func (c *Board) Print() {
	fmt.Print("--")
	for i := 0; i < 12; i++ {
		fmt.Print("---")
	}
	fmt.Print("\n|")
	for i := 0; i < 12; i++ {
		char := strconv.Itoa(i + 1)
		if len(char) == 1 {
			char += " "
		}
		fmt.Print(char + " ")
	}
	fmt.Println("|")
	fmt.Println("|a")
}
