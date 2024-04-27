package state

import (
	"mvvasilev/last_light/game/model"
	"mvvasilev/last_light/render"
	"mvvasilev/last_light/ui"

	"github.com/gdamore/tcell/v2"
)

type InventoryScreenState struct {
	prevState GameState
	exitMenu  bool

	inventoryMenu *ui.UIWindow

	player *model.Player
}

func CreateInventoryScreenState(player *model.Player, prevState GameState) *InventoryScreenState {
	iss := new(InventoryScreenState)

	iss.prevState = prevState
	iss.player = player
	iss.exitMenu = false

	iss.inventoryMenu = ui.CreateWindow(40, 0, 40, 24, "INVENTORY", tcell.StyleDefault)

	return iss
}

func (iss *InventoryScreenState) OnInput(e *tcell.EventKey) {
	if e.Key() == tcell.KeyEsc || (e.Key() == tcell.KeyRune && e.Rune() == 'i') {
		iss.exitMenu = true
	}
}

func (iss *InventoryScreenState) OnTick(dt int64) GameState {
	if iss.exitMenu {
		return iss.prevState
	}

	return iss
}

func (iss *InventoryScreenState) CollectDrawables() []render.Drawable {
	return append(
		iss.prevState.CollectDrawables(),
		iss.inventoryMenu,
	)
}
