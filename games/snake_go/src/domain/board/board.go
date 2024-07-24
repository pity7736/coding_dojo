package boardpackage

import (
	"math/rand/v2"

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
	board.init()
	return board
}

func (self *Board) init() {
	self.setValueInCells(self.snake.Head(), '$')
	self.createCookie()
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
	if self.isPositionWithinBoard(self.snake.Head()) {
		if self.isCookieInPosition(self.snake.Head()) {
			self.snake.Eat()
			self.createCookie()
		} else {
			self.setValueInCells(tailSnakePosition, ' ')
		}
		self.setSnake()
	}
}

func (self *Board) isCookieInPosition(position positionpackage.Position) bool {
	return self.cells[position.Row()][position.Column()] == '#'
}

func (self *Board) isPositionWithinBoard(position positionpackage.Position) bool {
	return position.Column() >= 0 &&
		position.Column() < int8(self.width) &&
		position.Row() < int8(self.height) &&
		position.Row() >= 0
}

func (self *Board) createCookie() {
	position := positionpackage.New(
		int8(rand.IntN(int(self.width))),
		int8(rand.IntN(int(self.height))),
	)
	if self.isCellEmpty(position) {
		self.setValueInCells(position, '#')
	} else {
		self.createCookie()
	}
}

func (self *Board) isCellEmpty(position positionpackage.Position) bool {
	return self.cells[position.Row()][position.Column()] == ' '
}

func (self *Board) setSnake() {
	if !self.SnakeHasCrashed() {
		self.setValueInCells(self.snake.Head(), '$')
		for _, position := range self.snake.Body() {
			self.setValueInCells(position, '-')
		}
	}
}

func (self *Board) setValueInCells(position positionpackage.Position, value rune) {
	self.cells[position.Row()][position.Column()] = value
}

func (self *Board) SnakeHasCrashed() bool {
	return self.snake.Head().Column() > int8(self.width) ||
		self.snake.Head().Column() < 0 ||
		self.snake.Head().Row() > int8(self.height) ||
		self.snake.Head().Row() < 0
}
