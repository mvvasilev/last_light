package ui

import (
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/input"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	"github.com/google/uuid"
)

type UIEventLog struct {
	id uuid.UUID

	eventLogger *engine.GameEventLog

	window *UIWindow

	style tcell.Style
}

func CreateUIEventLog(x, y, width, height int, eventLogger *engine.GameEventLog, style tcell.Style) *UIEventLog {
	return &UIEventLog{
		id:          uuid.New(),
		eventLogger: eventLogger,
		window:      CreateWindow(x, y, width, height, "", style),
		style:       style,
	}
}

func (uie *UIEventLog) MoveTo(x int, y int) {
	uie.window.MoveTo(x, y)
}

func (uie *UIEventLog) Position() engine.Position {
	return uie.window.Position()
}

func (uie *UIEventLog) Size() engine.Size {
	return uie.window.Size()
}

func (uie *UIEventLog) Input(inputAction input.InputAction) {

}

func (uie *UIEventLog) UniqueId() uuid.UUID {
	return uie.id
}

func (uie *UIEventLog) Draw(v views.View) {
	uie.window.Draw(v)

	x, y := uie.Position().XY()
	height := uie.Size().Height()

	textHeight := height - 2

	for i, ge := range uie.eventLogger.Tail(textHeight) {
		engine.DrawText(x+1, y+i+1, ge.Contents(), uie.style, v)
	}
}
