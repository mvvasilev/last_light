package state

import (
	"mvvasilev/last_light/render"

	"github.com/gdamore/tcell/v2"
)

type QuitState struct {
}

func (q *QuitState) OnInput(e *tcell.EventKey) {

}

func (q *QuitState) OnTick(dt int64) GameState {
	return q
}

func (q *QuitState) CollectDrawables() []render.Drawable {
	return render.Multidraw(nil)
}
