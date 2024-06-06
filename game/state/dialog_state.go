package state

import (
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/systems"
	"mvvasilev/last_light/game/ui"
)

type DialogState struct {
	inputSystem *systems.InputSystem
	turnSystem  *systems.TurnSystem

	prevState GameState

	dialog *ui.UIDialog

	returnToPreviousState bool
}

func CreateDialogState(inputSystem *systems.InputSystem, turnSystem *systems.TurnSystem, dialog *ui.UIDialog, prevState GameState) *DialogState {
	return &DialogState{
		inputSystem:           inputSystem,
		turnSystem:            turnSystem,
		prevState:             prevState,
		dialog:                dialog,
		returnToPreviousState: false,
	}
}

func (s *DialogState) InputContext() systems.InputContext {
	return systems.InputContext_Menu
}

func (ds *DialogState) OnTick(dt int64) GameState {
	if ds.inputSystem.NextAction() == systems.InputAction_Menu_Select {
		ds.returnToPreviousState = true
		ds.dialog.Select()
	}

	if ds.returnToPreviousState {
		return ds.prevState
	}

	return ds
}

func (ds *DialogState) CollectDrawables() []engine.Drawable {
	return append(ds.prevState.CollectDrawables(), ds.dialog)
}
