package state

import (
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/input"
	"mvvasilev/last_light/game/player"
	"mvvasilev/last_light/game/turns"
	"mvvasilev/last_light/game/ui/menu"

	"github.com/gdamore/tcell/v2"
)

type InventoryScreenState struct {
	inputSystem *input.InputSystem
	turnSystem  *turns.TurnSystem

	prevState GameState
	exitMenu  bool

	inventoryMenu         *menu.PlayerInventoryMenu
	selectedInventorySlot engine.Position

	player *player.Player
}

func CreateInventoryScreenState(inputSystem *input.InputSystem, turnSystem *turns.TurnSystem, player *player.Player, prevState GameState) *InventoryScreenState {
	iss := new(InventoryScreenState)

	iss.inputSystem = inputSystem
	iss.turnSystem = turnSystem
	iss.prevState = prevState
	iss.player = player
	iss.selectedInventorySlot = engine.PositionAt(0, 0)
	iss.exitMenu = false
	iss.inventoryMenu = menu.CreatePlayerInventoryMenu(43, 0, player.Inventory(), tcell.StyleDefault, tcell.StyleDefault.Background(tcell.ColorDarkSlateGray))

	return iss
}

func (s *InventoryScreenState) InputContext() input.Context {
	return input.InputContext_Inventory
}

func (iss *InventoryScreenState) OnTick(dt int64) (nextState GameState) {
	nextAction := iss.inputSystem.NextAction()
	nextState = iss

	switch nextAction {
	case input.InputAction_Menu_Exit:
		nextState = iss.prevState
	case input.InputAction_DropItem:
		iss.player.Inventory().Drop(iss.selectedInventorySlot.XY())
	case input.InputAction_Menu_HighlightUp:
		if iss.selectedInventorySlot.Y() == 0 {
			break
		}

		iss.selectedInventorySlot = iss.selectedInventorySlot.WithOffset(0, -1)
		iss.inventoryMenu.SelectSlot(iss.selectedInventorySlot.XY())
	case input.InputAction_Menu_HighlightDown:
		if iss.selectedInventorySlot.Y() == iss.player.Inventory().Shape().Height()-1 {
			break
		}

		iss.selectedInventorySlot = iss.selectedInventorySlot.WithOffset(0, +1)
		iss.inventoryMenu.SelectSlot(iss.selectedInventorySlot.XY())
	case input.InputAction_Menu_HighlightLeft:
		if iss.selectedInventorySlot.X() == 0 {
			break
		}

		iss.selectedInventorySlot = iss.selectedInventorySlot.WithOffset(-1, 0)
		iss.inventoryMenu.SelectSlot(iss.selectedInventorySlot.XY())
	case input.InputAction_Menu_HighlightRight:
		if iss.selectedInventorySlot.X() == iss.player.Inventory().Shape().Width()-1 {
			break
		}

		iss.selectedInventorySlot = iss.selectedInventorySlot.WithOffset(+1, 0)
		iss.inventoryMenu.SelectSlot(iss.selectedInventorySlot.XY())
	}

	return
}

func (iss *InventoryScreenState) CollectDrawables() []engine.Drawable {
	drawables := append(
		iss.prevState.CollectDrawables(),
		iss.inventoryMenu,
	)

	return drawables
}
