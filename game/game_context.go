package game

import (
	"fmt"
	"log"
	"mvvasilev/last_light/engine"
	"time"

	"github.com/gdamore/tcell/v2"
)

const TickRate int64 = 1 // tick every 50ms ( 20 ticks per second )

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

	tickRateText := engine.CreateText(0, 1, 16, 1, "0ms", tcell.StyleDefault)

	for {
		deltaTime := 1 + time.Since(lastLoop).Microseconds()
		lastLoop = time.Now()

		for _, e := range gc.renderContext.CollectInputEvents() {
			gc.game.Input(e)
		}

		deltaTickTime := time.Since(lastTick).Milliseconds()
		if deltaTickTime >= TickRate {
			tickRateText = engine.CreateText(0, 1, 16, 1, fmt.Sprintf("%vms", deltaTickTime), tcell.StyleDefault)

			stop := !gc.game.Tick(deltaTickTime)

			if stop {
				gc.renderContext.Stop()
				break
			}

			lastTick = time.Now()
		}

		drawables := gc.game.CollectDrawables()
		drawables = append(drawables, tickRateText)

		gc.renderContext.Draw(deltaTime, drawables)
	}
}
