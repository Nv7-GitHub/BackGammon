package main

import (
	"fmt"
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
	Turn     bool
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
	b.Tris[23] = Triangle{Black: 2}
	return b
}

// GameTurn has a single turn of the game, returns whether the game is over or not
func (b *Board) GameTurn() bool {
	fmt.Println()
	if b.Turn {
		fmt.Println("White's Turn")
	} else {
		fmt.Println("Black's Turn")
	}

	// Rolling
	rolls := Roll()
	fmt.Printf("You rolled a %d and a %d\n", rolls[0], rolls[1])
	if rolls[0] == rolls[1] {
		fmt.Println("Doubles!")
		rolls = append(rolls, rolls...)
	}
	fmt.Println()

	b.Print()
	fmt.Println()
	for len(rolls) > 0 {
		// Input
		s, e, err := b.GetInput()
		if err {
			return false
		}

		success := false
		// Is valid?
		for i, val := range rolls {
			if b.IsValid(val, s, e) {
				b.Tris[s] = SetPieces(b.Turn, b.Tris[s], GetPieces(b.Turn, b.Tris[s])-1)
				b.Tris[e] = SetPieces(b.Turn, b.Tris[e], GetPieces(b.Turn, b.Tris[e])+1)
				if i == len(rolls)-1 {
					rolls = rolls[:i]
				} else {
					rolls = append(rolls[:i], rolls[i+1:]...)
				}
				b.Print()
				fmt.Println()
				success = true
			}
		}
		if !success {
			fmt.Println("Invalid Move!")
			fmt.Println()
		}
	}
	b.Turn = !b.Turn // Next turn
	return true
}

// GetInput gets input
func (b *Board) GetInput() (int, int, bool) {
	// Input
	fmt.Print("Your Move: ")
	s, e, err := ParseMove(Input())
	for err != "" {
		if err == "exit" {
			return 0, 0, true
		}
		if err == "bear" { // Bear off piece
			if !b.Bear() {
				continue
			}
		}
		fmt.Println(err)
		fmt.Print("Your Move: ")
		s, e, err = ParseMove(Input())
	}
	return s, e, false
}

// Start starts the game
func (b *Board) Start() {
	roll := Roll()
	fmt.Printf("White rolled %d. Black rolled %d\n", roll[0], roll[1])
	for roll[0] == roll[1] {
		roll = Roll()
		fmt.Printf("White rolled %d. Black rolled %d\n", roll[0], roll[1])
	}
	if roll[0] > roll[1] {
		b.Turn = true
		fmt.Println("White goes first!")
	} else {
		b.Turn = false
		fmt.Println("Black goes first!")
	}
}
