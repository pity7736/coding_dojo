package main

import (
	playcontroller "github.com/pity7736/snake_go/src/application/play_controller"
	"github.com/pity7736/snake_go/src/presentation/console"
)

func main() {
	ui := console.New()
	playController := playcontroller.New(ui)
	playController.Play()
}
