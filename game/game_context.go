package game

import (
	"log"
	"mvvasilev/last_light/engine"
	"os"
	"time"
)

const TICK_RATE int64 = 50 // tick every 50ms ( 20 ticks per second )

type GameContext struct {
	renderContext *engine.RenderContext

	game *Game
}

func CreateGameContext() *GameContext {
	gc := new(GameContext)

	rc, err := engine.CreateRenderContext()

	if err != nil {
		log.Fatalf("%~v", err)
	}

	gc.renderContext = rc
	gc.game = CreateGame()

	return gc
}

func (gc *GameContext) Run() {
	lastLoop := time.Now()
	lastTick := time.Now()

	for {
		deltaTime := 1 + time.Since(lastLoop).Microseconds()
		lastLoop = time.Now()

		for _, e := range gc.renderContext.CollectInputEvents() {
			gc.game.Input(e)
		}

		if time.Since(lastTick).Milliseconds() >= TICK_RATE {
			stop := !gc.game.Tick(deltaTime)

			if stop {
				gc.renderContext.Stop()
				os.Exit(0)
				break
			}

			lastTick = time.Now()
		}

		drawables := gc.game.CollectDrawables()
		gc.renderContext.Draw(deltaTime, drawables)
	}
}
