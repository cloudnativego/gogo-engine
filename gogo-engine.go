package gogo

import "fmt"

func (position Coordinate) String() string {
	return fmt.Sprintf("(%d,%d)", position.X, position.Y)
}

func (move Move) String() string {
	return fmt.Sprintf("%d/%s", move.Player, move.Position)
}
