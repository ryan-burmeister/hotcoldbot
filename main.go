package main

import (
	"fmt"
	"log"
	"math"
	"os/exec"
	"time"
)

const (
	// The width and height of the board.
	width  = 20
	height = 10

	// The maximum number of steps to allow the
	// character to move before starting a new
	// round.
	maxSteps = 28

	// The time to wait in-between each step
	// the character moves.
	tick = 100 // Milliseconds

	// The time to wait in-between each round.
	roundTick = 500
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

// drawCharacterAndGoal clears the screen
// and board, adds the character and goal
// runes, and draws the board.
func drawCharacterAndGoal(b *board, c character, g position) {
	// Clear terminal to start fresh
	clearScreen()

	// Clear board, set runes, and draw
	b.clear('.')
	b.setRune(g.y, g.x, '*')
	b.setRune(c.pos.y, c.pos.x, c.symb)
	b.draw()
}

// drawInformation draws statistics related
// to the character and board.
func drawInformation(round, steps int, distance float64) {
	// Draw information
	fmt.Printf("Round: %v\n", round)
	fmt.Printf("Step: %v/%v\n", steps, maxSteps)
	fmt.Printf("Current Distance: %v\n", float32(distance))
}

func main() {
	// Create a new board with dots so it's
	// easy to see the differences between the
	// background and other symbols
	b := newBoard('.')

	// Create character and goal
	c := newCharacter()
	g := position{width - 1, height - 1}

	// Counters
	round := 1

	for {
		var prevPositions []position

		// Reset
		c.pos.x = 0
		c.pos.y = 0
		c.unmarkSteps()

		// Show first position and tick
		drawCharacterAndGoal(&b, c, g)
		drawInformation(round, 0, c.pos.distanceTo(g))
		time.Sleep(100 * time.Millisecond)

		// Handle steps
		for i, _ := range c.steps {
			// Handle updates
			c.move(i)

			// Mark the step for change if the last
			// position was closer than the current
			// one
			lastPos := position{0, 0}
			if i >= 1 {
				lastPos = prevPositions[i-1]
			}

			if lastPos.distanceTo(g) < c.pos.distanceTo(g) {
				c.steps[i].marked = true
			}

			// Record current position
			prevPositions = append(prevPositions, c.pos)

			// Draw board
			drawCharacterAndGoal(&b, c, g)
			drawInformation(round, i+1, c.pos.distanceTo(g))

			// Tick
			time.Sleep(100 * time.Millisecond)
		}

		// Allow character to learn from it's
		// mistakes
		c.learn()

		// Increment round counter and tick
		round++
		time.Sleep(roundTick * time.Millisecond)
	}
}
