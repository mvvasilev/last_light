package state

import (
	"mvvasilev/last_light/render"

	"github.com/gdamore/tcell/v2"
)

type LookState struct {
}

func (ls *LookState) OnInput(e *tcell.EventKey) {
	panic("not implemented") // TODO: Implement
}

func (ls *LookState) OnTick(dt int64) GameState {
	panic("not implemented") // TODO: Implement
}

func (ls *LookState) CollectDrawables() []render.Drawable {
	panic("not implemented") // TODO: Implement
}
