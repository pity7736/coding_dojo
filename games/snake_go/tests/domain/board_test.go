package domain_test

import (
	"testing"

	boardpackage "github.com/pity7736/snake_go/src/domain/board"
	"github.com/pity7736/snake_go/src/domain/constants"
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

func TestMoveSnakeUntilCrashRight(t *testing.T) {
	board := boardpackage.New()
	for i := uint8(0); i < board.Width()/2+1; i++ {
		board.MoveSnake(directionpackage.RIGHT())
	}

	assert.True(t, board.SnakeHasCrashed())
}

func TestMoveSnakeUntilCrashLeft(t *testing.T) {
	board := boardpackage.New()
	for i := uint8(0); i < board.Width()/2+1; i++ {
		board.MoveSnake(directionpackage.LEFT())
	}

	assert.True(t, board.SnakeHasCrashed())
}

func TestMoveSnakeUntilCrashDown(t *testing.T) {
	board := boardpackage.New()
	for i := uint8(0); i < board.Height()/2+1; i++ {
		board.MoveSnake(directionpackage.DOWN())
	}

	assert.True(t, board.SnakeHasCrashed())
}

func TestMoveSnakeUntilCrashUp(t *testing.T) {
	board := boardpackage.New()
	for i := uint8(0); i < board.Height()/2+1; i++ {
		board.MoveSnake(directionpackage.UP())
	}

	assert.True(t, board.SnakeHasCrashed())
}

func TestCookieIsCreated(t *testing.T) {
	board := boardpackage.New()
	count := 0
	for _, row := range board.Cells() {
		for _, cell := range row {
			if cell == constants.COOKIE_CHARACTER {
				count++
			}
		}
	}
	assert.Equal(t, 1, count)
}

func TestEatCookie(t *testing.T) {
	board := boardpackage.New()
	cookiePosition := getCookiePosition(board)
	_, opositeSnakePosition := eatCookie(board)
	oldSnakePosition := cookiePosition.Move(opositeSnakePosition)

	assert.True(t, isSnakeHeadInPosition(board.Cells(), cookiePosition))
	assert.True(t, isSnakeBodyInPosition(board.Cells(), oldSnakePosition))
	assert.True(t, checkOneSnakeHead(board.Cells()))
}

func TestEatCookieAndMove(t *testing.T) {
	board := boardpackage.New()
	cookiePosition := getCookiePosition(board)
	currentSnakeDirection, opositeSnakePosition := eatCookie(board)
	oldSnakePosition := cookiePosition.Move(opositeSnakePosition)
	board.MoveSnake(currentSnakeDirection)

	assert.True(t, isSnakeHeadInPosition(board.Cells(), cookiePosition.Move(currentSnakeDirection)))
	assert.True(t, isSnakeBodyInPosition(board.Cells(), cookiePosition))
	assert.True(t, isCellEmpty(board.Cells(), oldSnakePosition))
	assert.True(t, checkOneSnakeHead(board.Cells()))
}

func TestEatCookieTwiceAndMove(t *testing.T) {
	board := boardpackage.New()
	eatCookie(board)
	cookiePosition := getCookiePosition(board)
	currentSnakeDirection, opositeSnakePosition := eatCookie(board)
	oldSnakePosition := cookiePosition.Move(opositeSnakePosition).Move(opositeSnakePosition)
	board.MoveSnake(currentSnakeDirection)

	assert.True(t, isSnakeHeadInPosition(board.Cells(), cookiePosition.Move(currentSnakeDirection)))
	assert.True(t, isSnakeBodyInPosition(board.Cells(), cookiePosition))
	assert.True(t, isCellEmpty(board.Cells(), oldSnakePosition))
	assert.True(t, checkOneSnakeHead(board.Cells()))
}

func eatCookie(board *boardpackage.Board) (directionpackage.Direction, directionpackage.Direction) {
	var opositeSnakeDirection directionpackage.Direction
	var currentSnakeDirection directionpackage.Direction
	cookiePosition := getCookiePosition(board)
	snakePosition := getSnakePosition(board)
	for cookiePosition.Row() > snakePosition.Row() {
		snakePosition = snakePosition.Move(directionpackage.DOWN())
		board.MoveSnake(directionpackage.DOWN())
		opositeSnakeDirection = directionpackage.UP()
		currentSnakeDirection = directionpackage.DOWN()
	}

	for cookiePosition.Row() < snakePosition.Row() {
		snakePosition = snakePosition.Move(directionpackage.UP())
		board.MoveSnake(directionpackage.UP())
		currentSnakeDirection = directionpackage.UP()
		opositeSnakeDirection = directionpackage.DOWN()
	}

	for cookiePosition.Column() > snakePosition.Column() {
		snakePosition = snakePosition.Move(directionpackage.RIGHT())
		board.MoveSnake(directionpackage.RIGHT())
		currentSnakeDirection = directionpackage.RIGHT()
		opositeSnakeDirection = directionpackage.LEFT()
	}

	for cookiePosition.Column() < snakePosition.Column() {
		snakePosition = snakePosition.Move(directionpackage.LEFT())
		board.MoveSnake(directionpackage.LEFT())
		currentSnakeDirection = directionpackage.LEFT()
		opositeSnakeDirection = directionpackage.RIGHT()
	}
	return currentSnakeDirection, opositeSnakeDirection
}

func getCookiePosition(board *boardpackage.Board) positionpackage.Position {
	var position positionpackage.Position
	for i, row := range board.Cells() {
		for j, cell := range row {
			if cell == constants.COOKIE_CHARACTER {
				position = positionpackage.New(int8(i), int8(j))
			}
		}
	}
	return position
}

func getSnakePosition(board *boardpackage.Board) positionpackage.Position {
	var position positionpackage.Position
	for i, row := range board.Cells() {
		for j, cell := range row {
			if cell == constants.SNAKE_HEAD_CHARACTER {
				position = positionpackage.New(int8(i), int8(j))
			}
		}
	}
	return position
}

func checkOneSnakeHead(cells [][]rune) bool {
	count := 0
	for _, row := range cells {
		for _, cell := range row {
			if cell == constants.SNAKE_HEAD_CHARACTER {
				count++
			}
		}
	}
	return count == 1
}

func isSnakeHeadInPosition(cells [][]rune, position positionpackage.Position) bool {
	return cells[position.Row()][position.Column()] == constants.SNAKE_HEAD_CHARACTER
}

func isSnakeBodyInPosition(cells [][]rune, position positionpackage.Position) bool {
	return cells[position.Row()][position.Column()] == constants.SNAKE_BODY_CHARACTER
}

func isCellEmpty(cells [][]rune, position positionpackage.Position) bool {
	return cells[position.Row()][position.Column()] == constants.EMPTY_VALUE_CHARACTER
}
