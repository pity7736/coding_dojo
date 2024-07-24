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
	body := make([]positionpackage.Position, len(self.body))
	copy(body, self.body)
	self.previousTailPosition = self.Tail()
	self.body[0] = self.Head().Move(direction)
	for i := 1; i < len(self.body); i++ {
		self.body[i] = body[i-1]
	}
}

func (self *Snake) Eat() {
	self.body = append(self.body, self.previousTailPosition)
}

func (self *Snake) Tail() positionpackage.Position {
	return self.body[len(self.body)-1]
}

func (self *Snake) Head() positionpackage.Position {
	return self.body[0]
}

func (self *Snake) Body() []positionpackage.Position {
	return self.body[1:]
}
