package gogo

import "fmt"

// PerformMove accepts an intent to perform a move, validates it, and returns
// the corresponding change in match as a new match struct.
func (gameboard GameBoard) performMove(move Move) (outBoard GameBoard, err error) {
	//outBoard = NewBoard(cap(gameboard.Positions))
	outBoard = gameboard.copy()

	if gameboard.Positions[move.Position.X][move.Position.Y] != EmptyPosition {
		return outBoard, fmt.Errorf(RulesFailureSpaceOccupied, move)
	}
	outBoard.Positions[move.Position.X][move.Position.Y] = move.Player
	return outBoard, err
}

func (position Coordinate) String() string {
	return fmt.Sprintf("(%d,%d)", position.X, position.Y)
}

func (move Move) String() string {
	return fmt.Sprintf("%d/%s", move.Player, move.Position)
}
