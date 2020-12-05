package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var reader *bufio.Reader

func main() {
	fmt.Println(`Welcome to BackGammon! Type "exit" to exit. To move a piece from 1 to 3, you would input "1 3". To bear off a piece, type "bear".`)
	fmt.Println()
	rand.Seed(time.Now().UnixNano())   // Seed so that it is actually random
	reader = bufio.NewReader(os.Stdin) // For some reason fmt.Scanln wasnt working so using this

	b := NewBoard()
	b.Start()
	running := true
	for running {
		fmt.Println()
		running = b.GameTurn()
	}
	fmt.Println("Thanks for playing!")
}
