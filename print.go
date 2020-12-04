package main

import (
	"fmt"
	"strconv"
)

// Print prints the board
func (b *Board) Print() {
	// Print top
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

	for i := 0; i < 5; i++ {
		fmt.Print("|")
		for j := 0; j < 12; j++ {
			if i < b.Tris[j].White {
				fmt.Print("w  ")
			} else if i < (b.Tris[j].Black + b.Tris[j].White) {
				fmt.Print("b  ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println("|")
	}

	fmt.Print("--")
	for i := 0; i < 12; i++ {
		fmt.Print("---")
	}
	fmt.Println()

	for i := 5; i > 0; i-- {
		fmt.Print("|")
		for j := 23; j > 11; j-- {
			if i < b.Tris[j].White {
				fmt.Print("w  ")
			} else if i < (b.Tris[j].Black + b.Tris[j].White) {
				fmt.Print("b  ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println("|")
	}

	fmt.Print("|")
	for i := 23; i > 11; i-- {
		char := strconv.Itoa(i + 1)
		if len(char) == 1 {
			char += " "
		}
		fmt.Print(char + " ")
	}
	fmt.Println("|")

	fmt.Print("--")
	for i := 0; i < 12; i++ {
		fmt.Print("---")
	}
	fmt.Println()
}
