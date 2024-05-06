package state

import (
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/model"
	"mvvasilev/last_light/game/ui/menu"
	"mvvasilev/last_light/util"

	"github.com/gdamore/tcell/v2"
)

type InventoryScreenState struct {
	prevState PausableState
	exitMenu  bool

	inventoryMenu         *menu.PlayerInventoryMenu
	selectedInventorySlot util.Position

	player *model.Player

	moveInventorySlotDirection model.Direction
	dropSelectedInventorySlot  bool
}

func CreateInventoryScreenState(player *model.Player, prevState PausableState) *InventoryScreenState {
	iss := new(InventoryScreenState)

	iss.prevState = prevState
	iss.player = player
	iss.selectedInventorySlot = util.PositionAt(0, 0)
	iss.exitMenu = false
	iss.inventoryMenu = menu.CreatePlayerInventoryMenu(43, 0, player.Inventory(), tcell.StyleDefault, tcell.StyleDefault.Background(tcell.ColorDarkSlateGray))

	return iss
}

func (iss *InventoryScreenState) OnInput(e *tcell.EventKey) {
	if e.Key() == tcell.KeyEsc || (e.Key() == tcell.KeyRune && e.Rune() == 'i') {
		iss.exitMenu = true
	}

	if e.Key() == tcell.KeyRune && e.Rune() == 'x' {
		iss.dropSelectedInventorySlot = true
	}

	if e.Key() != tcell.KeyRune {
		return
	}

	switch e.Rune() {
	case 'k':
		iss.moveInventorySlotDirection = model.DirectionUp
	case 'j':
		iss.moveInventorySlotDirection = model.DirectionDown
	case 'h':
		iss.moveInventorySlotDirection = model.DirectionLeft
	case 'l':
		iss.moveInventorySlotDirection = model.DirectionRight
	}
}

func (iss *InventoryScreenState) OnTick(dt int64) GameState {
	if iss.exitMenu {
		iss.prevState.Unpause()
		return iss.prevState
	}

	if iss.dropSelectedInventorySlot {
		iss.player.Inventory().Drop(iss.selectedInventorySlot.XY())
		iss.dropSelectedInventorySlot = false
	}

	if iss.moveInventorySlotDirection != model.DirectionNone {

		switch iss.moveInventorySlotDirection {
		case model.DirectionUp:
			if iss.selectedInventorySlot.Y() == 0 {
				break
			}

			iss.selectedInventorySlot = iss.selectedInventorySlot.WithOffset(0, -1)
		case model.DirectionDown:
			if iss.selectedInventorySlot.Y() == iss.player.Inventory().Shape().Height()-1 {
				break
			}

			iss.selectedInventorySlot = iss.selectedInventorySlot.WithOffset(0, +1)
		case model.DirectionLeft:
			if iss.selectedInventorySlot.X() == 0 {
				break
			}

			iss.selectedInventorySlot = iss.selectedInventorySlot.WithOffset(-1, 0)
		case model.DirectionRight:
			if iss.selectedInventorySlot.X() == iss.player.Inventory().Shape().Width()-1 {
				break
			}

			iss.selectedInventorySlot = iss.selectedInventorySlot.WithOffset(+1, 0)
		}

		iss.inventoryMenu.SelectSlot(iss.selectedInventorySlot.XY())
		iss.moveInventorySlotDirection = model.DirectionNone
	}

	return iss
}

func (iss *InventoryScreenState) CollectDrawables() []engine.Drawable {
	drawables := append(
		iss.prevState.CollectDrawables(),
		iss.inventoryMenu,
	)

	return drawables
}
