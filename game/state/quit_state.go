package state

import (
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/input"
)

type QuitState struct {
}

func (s *QuitState) InputContext() input.Context {
	return input.InputContext_Menu
}

func (q *QuitState) OnTick(dt int64) GameState {
	return q
}

func (q *QuitState) CollectDrawables() []engine.Drawable {
	return engine.Multidraw(nil)
}
