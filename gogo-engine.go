package gogo

import (
	"fmt"
	"time"

	"code.google.com/p/go-uuid/uuid"
)

func (position Coordinate) String() string {
	return fmt.Sprintf("(%d,%d)", position.X, position.Y)
}

func (move Move) String() string {
	return fmt.Sprintf("%d/%s", move.Player, move.Position)
}

// NewMatch creates a new match
func NewMatch(size int) Match {
	result := Match{}
	result.ID = uuid.New()
	result.StartTime = time.Now()
	result.GameBoard = newBoard(size)
	result.TurnCount = 0
	result.GridSize = size
	return result
}
