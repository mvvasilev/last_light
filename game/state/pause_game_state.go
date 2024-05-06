package state

import (
	"mvvasilev/last_light/engine"
	engine1 "mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/ui"

	"github.com/gdamore/tcell/v2"
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

	s.pauseMenuWindow = ui.CreateWindow(int(engine.TERMINAL_SIZE_WIDTH)/2-15, int(engine.TERMINAL_SIZE_HEIGHT)/2-7, 30, 14, "PAUSED", tcell.StyleDefault)
	s.buttons = make([]*ui.UISimpleButton, 0)
	s.buttons = append(
		s.buttons,
		ui.CreateSimpleButton(
			int(s.pauseMenuWindow.Position().X())+3,
			int(s.pauseMenuWindow.Position().Y())+1,
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
			int(s.pauseMenuWindow.Position().X())+3,
			int(s.pauseMenuWindow.Position().Y())+3,
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
		pg.currButtonSelected = engine1.LimitIncrement(pg.currButtonSelected, 1)
		pg.buttons[pg.currButtonSelected].Highlight()
	}

	if e.Key() == tcell.KeyUp {
		pg.buttons[pg.currButtonSelected].Unhighlight()
		pg.currButtonSelected = engine1.LimitDecrement(pg.currButtonSelected, 0)
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

func (pg *PauseGameState) CollectDrawables() []engine.Drawable {
	arr := make([]engine.Drawable, 0)

	arr = append(arr, pg.prevState.CollectDrawables()...)

	arr = append(arr, pg.pauseMenuWindow)

	for _, b := range pg.buttons {
		arr = append(arr, b)
	}

	return arr
}
