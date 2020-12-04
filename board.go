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
	Turn     bool
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

// IsValid checks if a move is valid
func (b *Board) IsValid(d int, s int, e int) bool { // s = start, e = end
	if (GetPieces(b.Turn, b.Bar) != 0) && (s >= 0) { // If trying to move from the bar, put a neg num as input
		return false // Has pieces in the Bar but trying to move pieces not in the bar
	}
	absChange := (e - s) * Dir(b.Turn)
	if absChange != d { // Handles Dice AND going counterclockwise/clockwise
		return false
	}
	if GetPieces(!b.Turn, b.Tris[e]) >= 2 { // Spot is not open
		return false
	}
	if GetPieces(!b.Turn, b.Tris[e]) == 1 { // Blot! - Increase in bar
		b.Tris[e] = SetPieces(!b.Turn, b.Tris[e], 0)
		b.Bar = SetPieces(!b.Turn, b.Bar, GetPieces(!b.Turn, b.Bar)+1)
	}
	return true
}
