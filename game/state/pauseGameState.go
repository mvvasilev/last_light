package state

import (
	"mvvasilev/last_light/render"
	"mvvasilev/last_light/ui"
	"mvvasilev/last_light/util"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
)

type PauseGameState struct {
	prevState PausableState

	unpauseGame      bool
	returnToMainMenu bool

	pauseMenuWindow    *ui.UIWindow
	buttons            []*ui.UISimpleButton
	currButtonSelected int
}

func PauseGame(prevState PausableState) *PauseGameState {
	s := new(PauseGameState)

	s.prevState = prevState

	highlightStyle := tcell.StyleDefault.Attributes(tcell.AttrBold)

	s.pauseMenuWindow = ui.CreateWindow(uint16(render.TERMINAL_SIZE_WIDTH)/2-15, uint16(render.TERMINAL_SIZE_HEIGHT)/2-7, 30, 14, "PAUSED", tcell.StyleDefault)
	s.buttons = make([]*ui.UISimpleButton, 0)
	s.buttons = append(
		s.buttons,
		ui.CreateSimpleButton(
			uint16(s.pauseMenuWindow.Position().X())+3,
			uint16(s.pauseMenuWindow.Position().Y())+1,
			"Resume",
			tcell.StyleDefault,
			highlightStyle,
			func() {
				s.unpauseGame = true
			},
		),
	)
	s.buttons = append(
		s.buttons,
		ui.CreateSimpleButton(
			uint16(s.pauseMenuWindow.Position().X())+3,
			uint16(s.pauseMenuWindow.Position().Y())+3,
			"Exit To Main Menu",
			tcell.StyleDefault,
			highlightStyle,
			func() {
				s.returnToMainMenu = true
			},
		),
	)

	s.currButtonSelected = 0
	s.buttons[s.currButtonSelected].Highlight()

	return s
}

func (pg *PauseGameState) OnInput(e *tcell.EventKey) {
	if e.Key() == tcell.KeyEsc {
		pg.unpauseGame = true
	}

	if e.Key() == tcell.KeyDown {
		pg.buttons[pg.currButtonSelected].Unhighlight()
		pg.currButtonSelected = util.LimitIncrement(pg.currButtonSelected, 1)
		pg.buttons[pg.currButtonSelected].Highlight()
	}

	if e.Key() == tcell.KeyUp {
		pg.buttons[pg.currButtonSelected].Unhighlight()
		pg.currButtonSelected = util.LimitDecrement(pg.currButtonSelected, 0)
		pg.buttons[pg.currButtonSelected].Highlight()
	}

	if e.Key() == tcell.KeyEnter {
		pg.buttons[pg.currButtonSelected].Select()
	}
}

func (pg *PauseGameState) OnTick(dt int64) GameState {
	if pg.unpauseGame {
		pg.prevState.Unpause()
		return pg.prevState
	}

	if pg.returnToMainMenu {
		return NewMainMenuState()
	}

	return pg
}

func (pg *PauseGameState) OnDraw(c views.View) {
	pg.prevState.OnDraw(c)

	pg.pauseMenuWindow.Draw(c)

	for _, b := range pg.buttons {
		b.Draw(c)
	}
}
