package main

import "fmt"

// board represents a 2D plane comprised
// of runes.
type board [height][width]rune

// newBoard returns a new instance of
// a board.
func newBoard(r rune) board {
	var b board
	b.clear(r)
	return b
}

// board.clear replaces all values of the
// board with fresh cells.
func (b *board) clear(r rune) {
	for row, _ := range b {
		for col, _ := range b[row] {
			b[row][col] = r
		}
	}
}

// board.placeCell will change the value of
// the rune at position row, col.
func (b *board) setRune(row, col int, r rune) {
	b[row][col] = r
}

// board.draw will draw draw a
// representation of the board to the
// console.
func (b board) draw() {
	for _, r := range b {
		for _, c := range r {
			fmt.Print(string(c))
		}

		fmt.Println()
	}
}
