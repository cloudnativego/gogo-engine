package gogo

func genMove(x int, y int, player byte) Move {
	return Move{Player: player, Position: Coordinate{X: x, Y: y}}
}

// Copy copies the state of an existing board into a new board.
func (gameboard GameBoard) copy() GameBoard {
	outBoard := NewBoard(cap(gameboard.Positions))
	copy(outBoard.Positions, gameboard.Positions)
	return outBoard
}
