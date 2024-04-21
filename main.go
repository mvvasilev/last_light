package main

import (
	"fmt"
	"log"
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

	c.HandleInput(func(ev *tcell.EventKey) {
		if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
			c.Stop()
			os.Exit(0)
		}
	})

	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	rect := render.CreateRectangle(
		0, 0, 80, 24,
		'┌', '─', '┐',
		'│', '#', '│',
		'└', '─', '┘',
		false, true, defStyle,
	)

	// text := render.CreateText(1, 2, 8, 8, "Hello World! How are you today?", defStyle)

	// grid := render.CreateGrid(
	// 	11, 1, 3, 3, 3, 3,
	// 	'┌', '─', '┬', '┐',
	// 	'│', '#', '│', '│',
	// 	'├', '─', '┼', '┤',
	// 	'└', '─', '┴', '┘',
	// 	defStyle,
	// )

	layers := render.CreateLayeredDrawContainer()

	layers.Insert(0, rect)
	// layers.Insert(1, text)
	// layers.Insert(0, grid)

	c.HandleRender(func(view views.View, deltaTime int64) {
		fps := 1_000_000 / deltaTime

		fpsText := render.CreateText(0, 0, 16, 1, fmt.Sprintf("%v FPS", fps), defStyle)

		layers.Draw(view)
		fpsText.Draw(view)
	})

	c.BeginRendering()
}
