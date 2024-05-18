package component

import (
	"mvvasilev/last_light/engine/ecs"
	"mvvasilev/last_light/game/state"
)

const ComponentType_GameStateComponent = ecs.ComponentType_2

type GameStateComponent struct {
	GameState *state.GameState
}

func (gsc *GameStateComponent) Type() ecs.ComponentType {
	return ComponentType_GameStateComponent
}
