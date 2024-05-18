package game

import (
	"time"
)

const TICK_RATE int64 = 50 // tick every 50ms ( 20 ticks per second )

type GameContext struct {
	world *GameWorld
}

func CreateGameContext() *GameContext {
	gc := new(GameContext)

	gc.world = CreateGameWorld()

	return gc
}

func (gc *GameContext) Run() {
	lastLoop := time.Now()

	for {
		deltaTime := 1 + time.Since(lastLoop).Milliseconds()
		lastLoop = time.Now()

		gc.world.Tick(deltaTime)
	}
}
