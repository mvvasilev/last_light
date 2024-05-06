package state

import (
	"mvvasilev/last_light/engine"

	"github.com/gdamore/tcell/v2"
)

type QuitState struct {
}

func (q *QuitState) OnInput(e *tcell.EventKey) {

}

func (q *QuitState) OnTick(dt int64) GameState {
	return q
}

func (q *QuitState) CollectDrawables() []engine.Drawable {
	return engine.Multidraw(nil)
}
