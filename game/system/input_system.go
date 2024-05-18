package system

import (
	"math"
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/engine/ecs"
	"mvvasilev/last_light/game/component"
)

type InputSystem struct {
	engineContext *engine.EngineContext
}

func CreateInputSystem(ec *engine.EngineContext) *InputSystem {
	return &InputSystem{
		engineContext: ec,
	}
}

func (is *InputSystem) Name() string {
	return "InputSystem"
}

func (is *InputSystem) Order() int {
	return math.MinInt
}

func (is *InputSystem) Tick(world *ecs.World, deltaTime int64) {
	world.AddSingletonComponent(&component.InputComponent{
		KeyEvents: is.engineContext.CollectInputEvents(),
	})
}
