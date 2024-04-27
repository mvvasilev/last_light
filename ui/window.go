package ui

import (
	"mvvasilev/last_light/render"
	"mvvasilev/last_light/util"
	"unicode/utf8"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	"github.com/google/uuid"
)

type UIWindow struct {
	id uuid.UUID

	title *render.Text
	box   render.Rectangle
}

func CreateWindow(x, y, width, height int, title string, style tcell.Style) *UIWindow {
	w := new(UIWindow)

	titleLen := utf8.RuneCountInString(title)

	titlePos := (width / 2) - int(titleLen/2)

	w.title = render.CreateText(x+titlePos, y, int(titleLen), 1, title, style)

	w.box = render.CreateRectangle(
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

func (w *UIWindow) Position() util.Position {
	return w.box.Position()
}

func (w *UIWindow) Size() util.Size {
	return util.SizeOf(0, 0)
}

func (w *UIWindow) Draw(v views.View) {
	w.box.Draw(v)
	w.title.Draw(v)
}

func (w *UIWindow) Input(e *tcell.EventKey) {
}
