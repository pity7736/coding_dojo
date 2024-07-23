package boardpackage

import (
	directionpackage "github.com/pity7736/snake_go/src/domain/direction"
	positionpackage "github.com/pity7736/snake_go/src/domain/position"
	snakepackage "github.com/pity7736/snake_go/src/domain/snake"
)

type Board struct {
	cells  [][]rune
	width  uint8
	height uint8
	snake  *snakepackage.Snake
}

func New() *Board {
	width := int8(30)
	height := int8(30)
	cells := make([][]rune, height)
	for i := int8(0); i < width; i++ {
		cells[i] = make([]rune, width)
		for j := int8(0); j < height; j++ {
			cells[i][j] = ' '
		}
	}
	snake := snakepackage.New(positionpackage.New(width/2, height/2))
	board := &Board{cells: cells, width: uint8(width), height: uint8(height), snake: snake}
	board.setValueInCells(snake.Head(), '$')
	return board
}

func (self *Board) Cells() [][]rune {
	return self.cells
}

func (self *Board) Width() uint8 {
	return self.width
}

func (self *Board) Height() uint8 {
	return self.width
}

func (self *Board) MoveSnake(direction directionpackage.Direction) {
	tailSnakePosition := self.snake.Tail()
	self.snake.Move(direction)
	self.setValueInCells(tailSnakePosition, ' ')
	self.setValueInCells(self.snake.Head(), '$')
}

func (self *Board) setValueInCells(position positionpackage.Position, value rune) {
	self.cells[position.Row()][position.Column()] = value
}
