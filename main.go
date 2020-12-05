package main

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

var reader *bufio.Reader

func main() {
	rand.Seed(time.Now().UnixNano())   // Seed so that it is actually random
	reader = bufio.NewReader(os.Stdin) // For some reason fmt.Scanln wasnt working so using this

	b := NewBoard()
	b.Print()
	b.GameTurn()
}
