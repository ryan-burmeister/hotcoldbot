package main

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
// direction should be changed.
type step struct {
	dir    direction
	marked bool
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

// character.move will shift a character's
// position by one unit in the given
// direction.
func (c *character) move(d direction) {
	// Move in specified direction
	switch d {
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
	if c.pos.y < c.bounds.top {
		c.pos.y = c.bounds.top
	} else if c.pos.y > c.bounds.bottom {
		c.pos.y = c.bounds.bottom
	} else if c.pos.x < c.bounds.left {
		c.pos.x = c.bounds.left
	} else if c.pos.x > c.bounds.right {
		c.pos.x = c.bounds.right
	}
}
