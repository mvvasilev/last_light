package main

import (
	"fmt"
	"log"
	"mvvasilev/last_light/game"
	"mvvasilev/last_light/render"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
)

func main() {

	c, err := render.CreateRenderContext()

	if err != nil {
		log.Fatalf("%~v", err)
	}

	g := game.CreateGame()

	c.HandleInput(func(ev *tcell.EventKey) {
		if ev.Key() == tcell.KeyCtrlC {
			c.Stop()
			os.Exit(0)
		}

		g.Input(ev)
	})

	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	c.HandleRender(func(view views.View, deltaTime int64) {
		fps := 1_000_000 / deltaTime

		fpsText := render.CreateText(0, 0, 16, 1, fmt.Sprintf("%v FPS", fps), defStyle)

		keepGoing := g.Tick(deltaTime)

		if !keepGoing {
			c.Stop()
			os.Exit(0)
		}

		g.Draw(view)

		fpsText.Draw(view)
	})

	c.BeginRendering()
}
