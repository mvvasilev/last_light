package game

import (
	"log"
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/engine/ecs"
	"mvvasilev/last_light/game/component"
	"mvvasilev/last_light/game/system"
)

type GameWorld struct {
	ecs *ecs.World
}

func CreateGameWorld() *GameWorld {
	world := &GameWorld{
		ecs: ecs.CreateWorld(),
	}

	engineContext, err := engine.InitEngine()

	if err != nil {
		// TODO: error logs
		log.Fatalf("%~v", err)
		return nil
	}

	world.ecs.RegisterComponentType(component.ComponentType_RenderableComponent, "RenderableComponent")
	world.ecs.RegisterComponentType(component.ComponentType_InputComponent, "InputComponent")
	world.ecs.RegisterComponentType(component.ComponentType_GameStateComponent, "GameStateComponent")

	world.ecs.AddSystem(system.CreateRenderingSystem(engineContext))
	world.ecs.AddSystem(system.CreateInputSystem(engineContext))
	world.ecs.AddSystem(system.CreateGameStateSystem())

	world.ecs.AddSingletonComponent(&component.InputComponent{})

	world.ecs.CreateEntity(&component.DrawablesComponent{
		Priority:  0,
		Drawables: []engine.Drawable{},
	})

	return world
}

func (gw *GameWorld) World() *ecs.World {
	return gw.ecs
}

func (gw *GameWorld) Tick(dt int64) {
	gw.ecs.Tick(dt)
}
