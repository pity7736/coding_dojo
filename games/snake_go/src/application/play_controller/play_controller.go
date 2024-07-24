package playcontroller

import (
	"github.com/pity7736/snake_go/src/application/ui"
	boardpackage "github.com/pity7736/snake_go/src/domain/board"
)

type PlayController struct {
	board *boardpackage.Board
	ui    ui.UI
}

func New(ui ui.UI) *PlayController {
	return &PlayController{board: boardpackage.New(), ui: ui}
}

func (self *PlayController) Play() {
	self.ui.Show(self.board)
	direction := self.ui.AskDirection()
	directionChannel := self.ui.Start()
	for {
		if self.isPlaying() {
			self.board.MoveSnake(direction)
			self.ui.Show(self.board)
			select {
			case direction = <-directionChannel:
			default:
			}
		} else {
			self.ui.ShowLostMessage()
			break
		}
	}
}

func (self *PlayController) isPlaying() bool {
	return !self.board.SnakeHasCrashed()
}
