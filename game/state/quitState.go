package state

import (
	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
)

type QuitState struct {
}

func (q *QuitState) OnInput(e *tcell.EventKey) {

}

func (q *QuitState) OnTick(dt int64) GameState {
	return q
}

func (q *QuitState) OnDraw(c views.View) {

}
