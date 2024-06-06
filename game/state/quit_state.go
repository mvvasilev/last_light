package state

import (
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/systems"
)

type QuitState struct {
}

func (s *QuitState) InputContext() systems.InputContext {
	return systems.InputContext_Menu
}

func (q *QuitState) OnTick(dt int64) GameState {
	return q
}

func (q *QuitState) CollectDrawables() []engine.Drawable {
	return engine.Multidraw(nil)
}
