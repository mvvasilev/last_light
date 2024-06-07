package main

import (
	"mvvasilev/last_light/game"
)

func main() {
	gc := game.CreateGameContext()
	gc.Run()
}
