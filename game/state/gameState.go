package state

import (
	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
)

type GameState interface {
	OnInput(e *tcell.EventKey)
	OnTick(dt int64) GameState
	OnDraw(c views.View)
}

type PausableState interface {
	Pause()
	Unpause()
	SetPaused(paused bool)

	GameState
}
