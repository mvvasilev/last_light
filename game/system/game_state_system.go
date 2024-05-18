package system

import (
	"math"
	"mvvasilev/last_light/engine/ecs"
	"mvvasilev/last_light/game/component"
)

type GameStateSystem struct {
}

func CreateGameStateSystem() *GameStateSystem {
	return &GameStateSystem{}
}

func (gss *GameStateSystem) Name() string {
	return "GameStateSystem"
}

func (gss *GameStateSystem) Order() int {
	return math.MinInt + 100
}

func (gss *GameStateSystem) Tick(world *ecs.World, deltaTime int64) {
	comp, err := world.FetchSingletonComponent(component.ComponentType_GameStateComponent)

}
