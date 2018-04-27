package main

import (
	"fmt"
	"log"
	"os/exec"
)

const (
	// The width and height of the board.
	width  = 20
	height = 10

	// The maximum number of steps to allow the
	// character to move before starting a new
	// round.
	steps = 40

	// The time to wait in-between each step
	// the character moves.
	tick = 100 // Milliseconds
)

// position represents an x and y
// coordinate on a grid.
type position struct {
	x int
	y int
}

// clearScreen runs the "clear" command and
// outputs the results in the terminal.
func clearScreen() {
	bs, err := exec.Command("clear").Output()
	if err != nil {
		log.Fatalf("Unable to clear screen: %v", err)
	}

	fmt.Print(string(bs))
}

func main() {
	// Create a new board with dots so it's
	// easy to see the differences between the
	// background and other symbols
	b := newBoard('.')
}
