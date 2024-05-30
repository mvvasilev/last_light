package state

import (
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/input"
)

type GameState interface {
	InputContext() input.Context
	OnTick(dt int64) GameState
	CollectDrawables() []engine.Drawable
}
