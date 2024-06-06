package state

import (
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/systems"
	"mvvasilev/last_light/game/ui"

	"github.com/gdamore/tcell/v2"
)

type MainMenuState struct {
	turnSystem  *systems.TurnSystem
	inputSystem *systems.InputSystem

	menuTitle          *engine.Raw
	buttons            []*ui.UISimpleButton
	currButtonSelected int

	quitGame     bool
	startNewGame bool
}

func CreateMainMenuState(turnSystem *systems.TurnSystem, inputSystem *systems.InputSystem) *MainMenuState {
	turnSystem.Clear()

	state := new(MainMenuState)

	highlightStyle := tcell.StyleDefault.Attributes(tcell.AttrBold)

	state.turnSystem = turnSystem
	state.inputSystem = inputSystem

	state.menuTitle = engine.CreateRawDrawable(
		11, 1, tcell.StyleDefault.Attributes(tcell.AttrBold).Foreground(tcell.ColorYellow),
		" |                   |         |     _)         |      |  ",
		" |       _` |   __|  __|       |      |   _` |  __ \\   __|",
		" |      (   | \\__ \\  |         |      |  (   |  | | |  |  ",
		"_____| \\__,_| ____/ \\__|      _____| _| \\__, | _| |_| \\__|",
		"                                        |___/             ",
	)
	state.buttons = make([]*ui.UISimpleButton, 0)
	state.buttons = append(state.buttons, ui.CreateSimpleButton(11, 7, "New Game", tcell.StyleDefault, highlightStyle, func() {
		state.startNewGame = true
	}))
	state.buttons = append(state.buttons, ui.CreateSimpleButton(11, 9, "Key Bindings // TODO", tcell.StyleDefault, highlightStyle, func() {

	}))
	state.buttons = append(state.buttons, ui.CreateSimpleButton(11, 11, "Quit", tcell.StyleDefault, highlightStyle, func() {
		state.quitGame = true
	}))

	state.currButtonSelected = 0
	state.buttons[state.currButtonSelected].Highlight()

	return state
}

func (s *MainMenuState) InputContext() systems.InputContext {
	return systems.InputContext_Menu
}

func (mms *MainMenuState) OnTick(dt int64) GameState {
	nextAction := mms.inputSystem.NextAction()

	if nextAction == systems.InputAction_Menu_HighlightDown {
		mms.buttons[mms.currButtonSelected].Unhighlight()
		mms.currButtonSelected = engine.LimitIncrement(mms.currButtonSelected, 2)
		mms.buttons[mms.currButtonSelected].Highlight()
	}

	if nextAction == systems.InputAction_Menu_HighlightUp {
		mms.buttons[mms.currButtonSelected].Unhighlight()
		mms.currButtonSelected = engine.LimitDecrement(mms.currButtonSelected, 0)
		mms.buttons[mms.currButtonSelected].Highlight()
	}

	if nextAction == systems.InputAction_Menu_Select {
		mms.buttons[mms.currButtonSelected].Select()
	}

	if mms.quitGame {
		return &QuitState{}
	}

	if mms.startNewGame {
		return CreateCharacterCreationState(mms.turnSystem, mms.inputSystem)
	}

	return mms
}

func (mms *MainMenuState) CollectDrawables() []engine.Drawable {
	arr := make([]engine.Drawable, 0)

	arr = append(arr, mms.menuTitle)

	for _, b := range mms.buttons {
		arr = append(arr, b)
	}

	return arr
}
