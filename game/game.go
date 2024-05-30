package game

import (
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/input"
	"mvvasilev/last_light/game/state"
	"mvvasilev/last_light/game/turns"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	turnSystem  *turns.TurnSystem
	inputSystem *input.InputSystem

	state state.GameState

	quitGame bool
}

func CreateGame() *Game {
	game := new(Game)

	game.turnSystem = turns.CreateTurnSystem()

	game.inputSystem = input.CreateInputSystemWithDefaultBindings()

	game.state = state.CreateMainMenuState(game.turnSystem, game.inputSystem)

	return game
}

func (g *Game) Input(ev *tcell.EventKey) {
	if ev.Key() == tcell.KeyCtrlC {
		g.quitGame = true
	}

	g.inputSystem.Input(g.state.InputContext(), ev)
}

func (g *Game) Tick(dt int64) (continueGame bool) {
	continueGame = !g.quitGame

	s := g.state.OnTick(dt)

	switch s.(type) {
	case *state.QuitState:
		g.quitGame = true
	}

	g.state = s

	return
}

func (g *Game) CollectDrawables() []engine.Drawable {
	return g.state.CollectDrawables()
}
