package main

import (
	"fmt"
	"log"
	"math"
	"os/exec"
)

const (
	// The width and height of the board.
	width  = 20
	height = 10

	// The maximum number of steps to allow the
	// character to move before starting a new
	// round.
	maxSteps = 40

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

// position.distanceTo computes the
// distance between two points using
// Pathagoras' Theorem.
func (p position) distanceTo(p2 position) float64 {
	xDist := p.x - p2.x
	yDist := p.y - p2.y
	return math.Sqrt(math.Pow(float64(xDist), 2) + math.Pow(float64(yDist), 2))
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
	board := newBoard('.')
	board.draw()
}
