package state

import (
	"mvvasilev/last_light/engine"

	"github.com/gdamore/tcell/v2"
)

type GameState interface {
	OnInput(e *tcell.EventKey)
	OnTick(dt int64) GameState
	CollectDrawables() []engine.Drawable
}

type PausableState interface {
	Pause()
	Unpause()
	SetPaused(paused bool)

	GameState
}
