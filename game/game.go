package game

import (
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/state"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	state state.GameState

	quitGame bool
}

func CreateGame() *Game {
	game := new(Game)

	game.state = state.NewMainMenuState()

	return game
}

func (g *Game) Input(ev *tcell.EventKey) {
	if ev.Key() == tcell.KeyCtrlC {
		g.quitGame = true
	}

	g.state.OnInput(ev)
}

func (g *Game) Tick(dt int64) bool {
	if g.quitGame {
		return false
	}

	s := g.state.OnTick(dt)

	switch s.(type) {
	case *state.QuitState:
		return false
	}

	g.state = s

	return true
}

func (g *Game) CollectDrawables() []engine.Drawable {
	return g.state.CollectDrawables()
}
