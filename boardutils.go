package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ParseMove parses a move
func ParseMove(input string) (int, int, string) {
	if input == "exit\n" { // Exit
		return 0, 0, "exit"
	}
	if input == "bear\n" { // Exit
		return 0, 0, "bear"
	}
	parts := strings.Split(input[:len(input)-1], " ")
	if len(parts) != 2 {
		return 0, 0, "You need to provide 2 inputs!"
	}
	var val1 int
	if parts[0] == "b" {
		val1 = -1
	} else {
		var err error
		val1, err = strconv.Atoi(parts[0])
		if err != nil {
			return 0, 0, "Error During Parsing: " + err.Error()
		}
	}
	val2, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, "Error During Parsing: " + err.Error()
	}
	if (val1 > 24) || (val1 < 1) || (val2 > 24) || (val2 < 1) {
		return 0, 0, "Input out of bounds!"
	}
	return val1 - 1, val2 - 1, ""
}

// IsValid checks if a move is valid
func (b *Board) IsValid(d int, s int, e int) bool { // s = start, e = end
	if (GetPieces(b.Turn, b.Bar) != 0) && (s >= 0) { // If trying to move from the bar, put a neg num as input
		return false // Has pieces in the Bar but trying to move pieces not in the bar
	}
	absChange := (e - s) * Dir(b.Turn)
	if absChange != d { // Handles Dice AND going counterclockwise/clockwise
		return false
	}
	if GetPieces(!b.Turn, b.Tris[e]) >= 2 { // Spot is not open
		return false
	}
	if GetPieces(b.Turn, b.Tris[s]) <= 0 { // No Pieces to move
		return false
	}
	if GetPieces(!b.Turn, b.Tris[e]) == 1 { // Blot! - Increase in bar
		b.Tris[e] = SetPieces(!b.Turn, b.Tris[e], 0)
		b.Bar = SetPieces(!b.Turn, b.Bar, GetPieces(!b.Turn, b.Bar)+1)
	}
	return true
}

// Bear bears a piece off of the board, returns true if invalid
func (b *Board) Bear() bool {
	// Check if ready
	var spots []Triangle
	var poses [2]int
	if b.Turn {
		poses = [2]int{19, 24}
		spots = b.Tris[18:23]
	} else {
		poses = [2]int{1, 6}
		spots = b.Tris[0:5]
	}
	count := 0
	for _, spot := range spots {
		count += GetPieces(b.Turn, spot)
	}
	count += GetPieces(b.Turn, b.Finished)
	if count != 15 { // Not ready
		fmt.Println("You don't have all your pieces in your home!")
		return true
	}

	// Get Input
	var val int
	err := errors.New("")
	for err != nil {
		fmt.Println(err.Error())
		fmt.Print("Bear From: ")
		val, err = strconv.Atoi(Input())
		if err != nil {
			// Got Input
			if (val > poses[1]) || (val < poses[0]) {
				err = errors.New("")
				fmt.Println("You need to input a value in your home!")
			}
			spotPos := -1
			for (val > poses[0]) && (spotPos > 0) {
				if GetPieces(b.Turn, b.Tris[val-1]) > 0 {
					spotPos = val
				}
			}
			if spotPos > 0 {
				b.Tris[spotPos] = SetPieces(b.Turn, b.Tris[spotPos], GetPieces(b.Turn, b.Tris[spotPos])-1)
				b.Finished = SetPieces(b.Turn, b.Finished, GetPieces(b.Turn, b.Finished)+1)
			}
		} // No input
	}
	return false
}
