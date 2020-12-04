package main

// Triangle stores the amount of white and black pieces
type Triangle struct {
	White int
	Black int
}

// Board represents the board
type Board struct {
	Tris     [24]Triangle
	Bar      Triangle
	Finished Triangle
}

// NewBoard creates a board
func NewBoard() Board {
	b := Board{}
	b.Tris[0] = Triangle{White: 2}
	b.Tris[5] = Triangle{Black: 5}
	b.Tris[7] = Triangle{Black: 3}
	b.Tris[11] = Triangle{White: 5}
	b.Tris[12] = Triangle{Black: 5}
	b.Tris[16] = Triangle{White: 3}
	b.Tris[18] = Triangle{White: 5}
	b.Tris[23] = Triangle{Black: 3}
	return b
}
