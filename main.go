package main

import (
	"fmt"
	"log"
	"mvvasilev/last_light/render"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

func main() {

	s, err := tcell.NewScreen()

	if err != nil {
		log.Fatalf("%~v", err)
	}

	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	// width, height := s.Size()

	// if width < 50 || height < 50 {
	// 	log.Fatalf("Your terminal must be at least 50x50")
	// }

	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	rect := render.CreateRectangle(
		0, 1, 10, 10,
		'┌', '─', '┐',
		'│', '#', '│',
		'└', '─', '┘',
		defStyle,
	)

	text := render.CreateText(1, 2, 8, 8, "Hello World! How are you today?", defStyle)

	layers := render.CreateLayeredDrawContainer()

	layers.Insert(0, rect)
	layers.Insert(1, text)

	layers.Remove(text.UniqueId())

	events := make(chan tcell.Event)
	quit := make(chan struct{})

	go s.ChannelEvents(events, quit)

	lastTime := time.Now()

	for {
		deltaTime := 1 + time.Since(lastTime).Microseconds()
		lastTime = time.Now()

		s.Clear()

		fps := 1_000_000 / deltaTime

		fpsText := render.CreateText(0, 0, 16, 1, fmt.Sprintf("%v FPS", fps), defStyle)

		layers.Draw(s)
		fpsText.Draw(s)

		s.Show()

		select {
		case ev, ok := <-events:

			if !ok {
				break
			}

			switch ev := ev.(type) {
			case *tcell.EventResize:
				s.Sync()
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					s.Fini()
					os.Exit(0)
				}
			}
		default:
		}
	}

}
