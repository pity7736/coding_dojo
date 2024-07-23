package positionpackage

import directionpackage "github.com/pity7736/snake_go/src/domain/direction"

type Position struct {
	row    int8
	column int8
}

func New(row, column int8) Position {
	return Position{row, column}
}

func (self Position) Move(direction directionpackage.Direction) Position {
	switch direction {
	case directionpackage.DOWN():
		return New(self.row+1, self.column)
	case directionpackage.UP():
		return New(self.row-1, self.column)
	case directionpackage.RIGHT():
		return New(self.row, self.column+1)
	default:
		return New(self.row, self.column-1)
	}
}

func (self Position) Row() int8 {
	return self.row
}

func (self Position) Column() int8 {
	return self.column
}
