package main

import "mvvasilev/last_light/game"

func main() {
	gc := game.CreateGameContext()
	gc.Run()
}

func runGame() {
	gc := game.CreateGameContext()
	gc.Run()
}
