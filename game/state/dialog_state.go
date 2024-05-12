package state

import (
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/ui"

	"github.com/gdamore/tcell/v2"
)

type DialogState struct {
	prevState GameState

	dialog *ui.UIDialog

	selectDialog          bool
	returnToPreviousState bool
}

func CreateDialogState(dialog *ui.UIDialog, prevState GameState) *DialogState {
	return &DialogState{
		prevState:             prevState,
		dialog:                dialog,
		returnToPreviousState: false,
	}
}

func (ds *DialogState) OnInput(e *tcell.EventKey) {
	if e.Key() == tcell.KeyEnter {
		ds.selectDialog = true
		return
	}

	ds.dialog.Input(e)
}

func (ds *DialogState) OnTick(dt int64) GameState {
	if ds.selectDialog {
		ds.selectDialog = false
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
