package console

import (
	"fmt"
	"log"
	"time"

	boardpackage "github.com/pity7736/snake_go/src/domain/board"
	"github.com/pity7736/snake_go/src/domain/constants"
	directionpackage "github.com/pity7736/snake_go/src/domain/direction"
)

type ConsoleUI struct{}

func New() *ConsoleUI {
	return &ConsoleUI{}
}

func (self *ConsoleUI) Show(board *boardpackage.Board) {
	time.Sleep(time.Millisecond * 100)
	fmt.Printf("\033[H\033[2J")
	fmt.Println("aoeu")
	self.printHorizontalLine(board)
	for _, row := range board.Cells() {
		fmt.Print("|")
		for _, cell := range row {
			if cell == constants.EMPTY_VALUE_CHARACTER {
				fmt.Print("  ")
			} else {
				fmt.Printf("%c ", cell)
			}
		}
		fmt.Print("|")
		fmt.Println("")
	}
	self.printHorizontalLine(board)
	fmt.Println("")
}

func (self *ConsoleUI) printHorizontalLine(board *boardpackage.Board) {
	for range board.Cells() {
		fmt.Print("--")
	}
	fmt.Println("-")
}

func (self *ConsoleUI) Start() chan directionpackage.Direction {
	channel := make(chan directionpackage.Direction)
	go func() {
		for {
			direction := self.askDirection()
			if direction != nil {
				channel <- *direction
			}
		}
	}()
	return channel
}

func (self *ConsoleUI) AskDirection() directionpackage.Direction {
	direction := self.askDirection()
	if direction != nil {
		return *direction
	}
	return self.AskDirection()
}

func (self *ConsoleUI) askDirection() *directionpackage.Direction {
	var option string
	_, err := fmt.Scanln(&option)
	if err != nil {
		log.Fatal(err)
	}
	var direction directionpackage.Direction
	switch option {
	case "u":
		direction = directionpackage.RIGHT()
	case "e":
		direction = directionpackage.DOWN()
	case "a":
		direction = directionpackage.LEFT()
	case ".":
		direction = directionpackage.UP()
	}
	return &direction
}

func (self *ConsoleUI) ShowLostMessage() {
	fmt.Println("¡Perdiste!")
}
