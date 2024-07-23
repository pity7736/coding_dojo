package domain_test

import (
	"testing"

	boardpackage "github.com/pity7736/snake_go/src/domain/board"
	directionpackage "github.com/pity7736/snake_go/src/domain/direction"
	positionpackage "github.com/pity7736/snake_go/src/domain/position"
	"github.com/stretchr/testify/assert"
)

func TestInitialPosition(t *testing.T) {
	board := boardpackage.New()
	snakePosition := positionpackage.New(
		int8(board.Width()/2),
		int8(board.Height()/2),
	)

	assert.True(t, isSnakeHeadInPosition(board.Cells(), snakePosition))
}

func TestMoveSnakeRight(t *testing.T) {
	board := boardpackage.New()
	initialSnakePosition := positionpackage.New(
		int8(board.Width()/2),
		int8(board.Height()/2),
	)

	board.MoveSnake(directionpackage.RIGHT())

	snakePosition := positionpackage.New(
		initialSnakePosition.Row(),
		initialSnakePosition.Column()+1,
	)
	cells := board.Cells()
	assert.True(t, isSnakeHeadInPosition(cells, snakePosition))
	assert.True(t, isCellEmpty(cells, initialSnakePosition))
}

func TestMoveSnakeDown(t *testing.T) {
	board := boardpackage.New()
	initialSnakePosition := positionpackage.New(
		int8(board.Width()/2),
		int8(board.Height()/2),
	)

	board.MoveSnake(directionpackage.DOWN())

	snakePosition := positionpackage.New(
		initialSnakePosition.Row()+1,
		initialSnakePosition.Column(),
	)
	cells := board.Cells()
	assert.True(t, isSnakeHeadInPosition(cells, snakePosition))
	assert.True(t, isCellEmpty(cells, initialSnakePosition))
}

func TestMoveSnakeUp(t *testing.T) {
	board := boardpackage.New()
	initialSnakePosition := positionpackage.New(
		int8(board.Width()/2),
		int8(board.Height()/2),
	)

	board.MoveSnake(directionpackage.DOWN())
	board.MoveSnake(directionpackage.DOWN())
	board.MoveSnake(directionpackage.RIGHT())
	board.MoveSnake(directionpackage.UP())

	snakePosition := positionpackage.New(
		initialSnakePosition.Row()+1,
		initialSnakePosition.Column()+1,
	)
	cells := board.Cells()
	assert.True(t, isSnakeHeadInPosition(cells, snakePosition))
	assert.True(t, isCellEmpty(cells, initialSnakePosition))
}

func TestMoveSnakeLeft(t *testing.T) {
	board := boardpackage.New()
	initialSnakePosition := positionpackage.New(
		int8(board.Width()/2),
		int8(board.Height()/2),
	)

	board.MoveSnake(directionpackage.DOWN())
	board.MoveSnake(directionpackage.RIGHT())
	board.MoveSnake(directionpackage.RIGHT())
	board.MoveSnake(directionpackage.LEFT())

	snakePosition := positionpackage.New(
		initialSnakePosition.Row()+1,
		initialSnakePosition.Column()+1,
	)
	cells := board.Cells()
	assert.True(t, isSnakeHeadInPosition(cells, snakePosition))
	assert.True(t, isCellEmpty(cells, initialSnakePosition))
}

func isSnakeHeadInPosition(cells [][]rune, position positionpackage.Position) bool {
	return cells[position.Row()][position.Column()] == '$'
}

func isCellEmpty(cells [][]rune, position positionpackage.Position) bool {
	return cells[position.Row()][position.Column()] == ' '
}
