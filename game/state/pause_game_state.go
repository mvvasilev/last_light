package state

import (
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/systems"
	"mvvasilev/last_light/game/ui"

	"github.com/gdamore/tcell/v2"
)

type PauseGameState struct {
	turnSystem  *systems.TurnSystem
	inputSystem *systems.InputSystem

	prevState GameState

	unpauseGame      bool
	returnToMainMenu bool

	pauseMenuWindow    *ui.UIWindow
	buttons            []*ui.UISimpleButton
	currButtonSelected int
}

func PauseGame(prevState GameState, turnSystem *systems.TurnSystem, inputSystem *systems.InputSystem) *PauseGameState {
	s := new(PauseGameState)

	s.turnSystem = turnSystem
	s.inputSystem = inputSystem
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

func (s *PauseGameState) InputContext() systems.InputContext {
	return systems.InputContext_Menu
}

func (pg *PauseGameState) OnTick(dt int64) GameState {
	switch pg.inputSystem.NextAction() {
	case systems.InputAction_Menu_Exit:
		pg.unpauseGame = true
	case systems.InputAction_Menu_HighlightDown:
		pg.buttons[pg.currButtonSelected].Unhighlight()
		pg.currButtonSelected = engine.LimitIncrement(pg.currButtonSelected, 1)
		pg.buttons[pg.currButtonSelected].Highlight()
	case systems.InputAction_Menu_HighlightUp:
		pg.buttons[pg.currButtonSelected].Unhighlight()
		pg.currButtonSelected = engine.LimitDecrement(pg.currButtonSelected, 0)
		pg.buttons[pg.currButtonSelected].Highlight()
	case systems.InputAction_Menu_Select:
		pg.buttons[pg.currButtonSelected].Select()
	}

	if pg.unpauseGame {
		return pg.prevState
	}

	if pg.returnToMainMenu {
		return CreateMainMenuState(pg.turnSystem, pg.inputSystem)
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
