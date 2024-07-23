package snakepackage

import (
	directionpackage "github.com/pity7736/snake_go/src/domain/direction"
	positionpackage "github.com/pity7736/snake_go/src/domain/position"
)

type Snake struct {
	body                 []positionpackage.Position
	previousTailPosition positionpackage.Position
}

func New(headPosition positionpackage.Position) *Snake {
	body := make([]positionpackage.Position, 0, 15)
	body = append(body, headPosition)
	return &Snake{body: body, previousTailPosition: headPosition}
}

func (self *Snake) Move(direction directionpackage.Direction) {
	self.previousTailPosition = self.Tail()
	self.body[0] = self.Head().Move(direction)
}

func (self *Snake) Tail() positionpackage.Position {
	return self.body[len(self.body)-1]
}

func (self *Snake) Head() positionpackage.Position {
	return self.body[0]
}
