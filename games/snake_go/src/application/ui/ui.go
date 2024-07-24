package ui

import (
	boardpackage "github.com/pity7736/snake_go/src/domain/board"
	directionpackage "github.com/pity7736/snake_go/src/domain/direction"
)

type UI interface {
	Show(board *boardpackage.Board)
	Start() chan directionpackage.Direction
	AskDirection() directionpackage.Direction
	ShowLostMessage()
}
