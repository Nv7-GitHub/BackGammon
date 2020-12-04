package main

func (b *Board) moveValid(d int, s int, e int) bool { // s = start, e = end
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
