package ui

import (
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/input"
	"unicode/utf8"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	"github.com/google/uuid"
)

type UIWindow struct {
	id uuid.UUID

	title *engine.Text
	box   engine.Rectangle
}

func CreateWindow(x, y, width, height int, title string, style tcell.Style) *UIWindow {
	w := new(UIWindow)

	titleLen := utf8.RuneCountInString(title)

	titlePos := (width / 2) - int(titleLen/2)

	if title != "" {
		w.title = engine.CreateText(x+titlePos, y, int(titleLen), 1, title, style)
	}

	w.box = engine.CreateRectangle(
		x, y, width, height,
		'┌', '─', '┐',
		'│', ' ', '│',
		'└', '─', '┘',
		false, true, style,
	)

	return w
}

func (w *UIWindow) UniqueId() uuid.UUID {
	return w.id
}

func (w *UIWindow) MoveTo(x int, y int) {

}

func (w *UIWindow) Position() engine.Position {
	return w.box.Position()
}

func (w *UIWindow) Size() engine.Size {
	return w.box.Size()
}

func (w *UIWindow) Draw(v views.View) {
	w.box.Draw(v)

	if w.title != nil {
		w.title.Draw(v)
	}
}

func (w *UIWindow) Input(inputAction input.InputAction) {
}
