package state

import (
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/systems"
)

type GameState interface {
	InputContext() systems.InputContext
	OnTick(dt int64) GameState
	CollectDrawables() []engine.Drawable
}
