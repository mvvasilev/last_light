package state

import (
	"mvvasilev/last_light/game/model"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
)

type PlayingState struct {
	player *model.Player

	pauseGame bool
}

func BeginPlayingState() *PlayingState {
	s := new(PlayingState)

	s.player = model.CreatePlayer(10, 10, tcell.StyleDefault)

	return s
}

func (ps *PlayingState) Pause() {
	ps.pauseGame = true
}

func (ps *PlayingState) Unpause() {
	ps.pauseGame = false
}

func (ps *PlayingState) SetPaused(paused bool) {
	ps.pauseGame = paused
}

func (ps *PlayingState) OnInput(e *tcell.EventKey) {
	ps.player.Input(e)

	if e.Key() == tcell.KeyEsc {
		ps.pauseGame = true
	}
}

func (ps *PlayingState) OnTick(dt int64) GameState {
	ps.player.Tick(dt)

	if ps.pauseGame {
		return PauseGame(ps)
	}

	return ps
}

func (ps *PlayingState) OnDraw(c views.View) {
	ps.player.Draw(c)
}
