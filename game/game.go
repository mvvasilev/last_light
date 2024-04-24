package game

import (
	"mvvasilev/last_light/game/state"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
)

type Game struct {
	state state.GameState
}

func CreateGame() *Game {
	game := new(Game)

	game.state = state.NewMainMenuState()

	return game
}

func (g *Game) Input(ev *tcell.EventKey) {
	g.state.OnInput(ev)
}

func (g *Game) Tick(dt int64) bool {
	s := g.state.OnTick(dt)

	switch s.(type) {
	case *state.QuitState:
		return false
	}

	g.state = s

	return true
}

func (g *Game) Draw(v views.View) {
	g.state.OnDraw(v)
}
