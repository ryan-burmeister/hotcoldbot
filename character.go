package main

import "math/rand"

// direction represents a direction in 2D
// space.
type direction int

const (
	DIRECTION_UP = iota
	DIRECTION_DOWN
	DIRECTION_LEFT
	DIRECTION_RIGHT
)

// bounds represents the boundaries a
// character's position should remain
// inside.
type bounds struct {
	top    int
	bottom int
	left   int
	right  int
}

// step represents an instruction on where
// to move next, as well as if the
// direction should be changed. It also
// has a memories of directions it has not
// yet tried to move in.
type step struct {
	dir        direction
	marked     bool
	unusedDirs []direction
}

// step.newStep creates a new instance of
// step with a random direction and adds
// the other directions into
// step.unusedDirs.
func newStep() step {
	var s step
	s.dir = direction(rand.Intn(4))
	s.marked = false
	for i := 0; i < 4; i++ {
		if i != int(s.dir) {
			s.unusedDirs = append(s.unusedDirs, direction(i))
		}
	}

	return s
}

// character represents a symbol with a
// position that can move along a a slice
// of steps.
type character struct {
	steps  [maxSteps]step
	bounds bounds
	pos    position
	symb   rune
}

// newCharacter creates a new instance of
// character with random steps and a '+'
// symbol.
func newCharacter() character {
	var c character
	for i, _ := range c.steps {
		c.steps[i] = newStep()
	}

	// Bounds
	c.bounds.top = 0
	c.bounds.bottom = height - 1
	c.bounds.left = 0
	c.bounds.right = width - 1

	// Position
	c.pos.x = 0
	c.pos.y = 0

	// Symbol
	c.symb = '+'

	return c
}

// character.unmarkSteps sets the marked
// status of all steps to false.
func (c *character) unmarkSteps() {
	for i, _ := range c.steps {
		c.steps[i].marked = false
	}
}

// character.learn iterates through steps
// and picks a new random direction for
// each one if it thinks it's neccessary.
func (c *character) learn() {
	for i, _ := range c.steps {
		if c.steps[i].marked && len(c.steps[i].unusedDirs) > 0 {
			// Unmark this step
			c.steps[i].marked = false

			// Pick a random new direction and apply
			// then delete it from the unusedDirs slice
			unused := c.steps[i].unusedDirs
			num := rand.Intn(len(c.steps[i].unusedDirs))
			c.steps[i].dir = c.steps[i].unusedDirs[num]
			c.steps[i].unusedDirs = append(unused[:num], unused[num+1:]...)
		}
	}
}

// character.move will shift a character's
// position by one unit in the given
// direction.
func (c *character) move(i int) {
	// Move in specified direction
	switch c.steps[i].dir {
	case DIRECTION_UP:
		c.pos.y -= 1
	case DIRECTION_DOWN:
		c.pos.y += 1
	case DIRECTION_LEFT:
		c.pos.x -= 1
	case DIRECTION_RIGHT:
		c.pos.x += 1
	}

	// Stay within bounds
	var outOfBounds bool
	if c.pos.y < c.bounds.top {
		c.pos.y = c.bounds.top
		outOfBounds = true
	} else if c.pos.y > c.bounds.bottom {
		c.pos.y = c.bounds.bottom
		outOfBounds = true
	} else if c.pos.x < c.bounds.left {
		c.pos.x = c.bounds.left
		outOfBounds = true
	} else if c.pos.x > c.bounds.right {
		c.pos.x = c.bounds.right
		outOfBounds = true
	}

	// Mark step for change if it's out of
	// bounds
	if outOfBounds {
		c.steps[i].marked = true
	}
}
