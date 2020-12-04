package main

import (
	"fmt"
	"strconv"
)

// Print prints the board
func (b *Board) Print() {
	// Print top
	fmt.Println("-------------------------------------------")

	// Nums
	fmt.Print("|")
	for i := 0; i < 12; i++ {
		char := strconv.Itoa(i + 1)
		if len(char) == 1 {
			char += " "
		}
		if i == 6 {
			fmt.Print("| | ")
		}
		fmt.Print(char + " ")
	}
	fmt.Println("|")

	// Rows for first half
	for i := 0; i < 5; i++ {
		fmt.Print("|")
		for j := 0; j < 12; j++ {
			if j == 6 {
				fmt.Print("|")
				if i < b.Bar.Black {
					fmt.Print("b")
				} else {
					fmt.Print(" ")
				}
				fmt.Print("| ")
			}
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

	// Middle
	fmt.Println("-------------------| |---------------------")

	// Rows for second half
	for i := 5; i > -1; i-- {
		fmt.Print("|")
		for j := 23; j > 11; j-- {
			if j == 17 {
				fmt.Print("|")
				if i < b.Bar.White {
					fmt.Print("w")
				} else {
					fmt.Print(" ")
				}
				fmt.Print("| ")
			}
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

	// Nums for second half
	fmt.Print("|")
	for i := 23; i > 11; i-- {
		char := strconv.Itoa(i + 1)
		if len(char) == 1 {
			char += " "
		}
		if i == 17 {
			fmt.Print("| | ")
		}
		fmt.Print(char + " ")
	}
	fmt.Println("|")

	// End
	fmt.Println("-------------------------------------------")
}
