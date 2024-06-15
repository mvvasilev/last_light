package state

import (
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/model"
	"mvvasilev/last_light/game/systems"
	"mvvasilev/last_light/game/ui/menu"

	"github.com/gdamore/tcell/v2"
)

type InventoryScreenState struct {
	eventLog    *engine.GameEventLog
	inputSystem *systems.InputSystem
	turnSystem  *systems.TurnSystem
	dungeon     *model.Dungeon

	prevState GameState
	exitMenu  bool

	inventoryMenu         *menu.PlayerInventoryMenu
	selectedInventorySlot engine.Position

	player *model.Player
}

func CreateInventoryScreenState(eventLog *engine.GameEventLog, dungeon *model.Dungeon, inputSystem *systems.InputSystem, turnSystem *systems.TurnSystem, player *model.Player, prevState GameState) *InventoryScreenState {
	iss := new(InventoryScreenState)

	iss.eventLog = eventLog
	iss.inputSystem = inputSystem
	iss.turnSystem = turnSystem
	iss.prevState = prevState
	iss.player = player
	iss.selectedInventorySlot = engine.PositionAt(0, 0)
	iss.exitMenu = false
	iss.inventoryMenu = menu.CreatePlayerInventoryMenu(43, 0, player.Inventory(), tcell.StyleDefault, tcell.StyleDefault.Background(tcell.ColorDarkSlateGray))
	iss.dungeon = dungeon

	return iss
}

func (s *InventoryScreenState) InputContext() systems.InputContext {
	return systems.InputContext_Inventory
}

func (iss *InventoryScreenState) OnTick(dt int64) (nextState GameState) {
	nextAction := iss.inputSystem.NextAction()
	nextState = iss

	switch nextAction {
	case systems.InputAction_Menu_Exit:
		nextState = iss.prevState
	case systems.InputAction_InteractItem:
		item := iss.player.Inventory().ItemAt(iss.selectedInventorySlot.XY())

		if item == nil {
			break
		}

		if item.Usable() != nil {
			item.Usable().Use(iss.eventLog, iss.dungeon, iss.player)

			iss.player.Inventory().ReduceQuantityAt(iss.selectedInventorySlot.X(), iss.selectedInventorySlot.Y(), 1)
			return
		}

		if item.Equippable() != nil {
			if iss.player.Inventory().AtSlot(item.Equippable().Slot) != nil {
				iss.player.Inventory().Push(iss.player.Inventory().AtSlot(item.Equippable().Slot))
			}

			iss.player.Inventory().Equip(item, item.Equippable().Slot)
		}

		iss.player.Inventory().Drop(iss.selectedInventorySlot.X(), iss.selectedInventorySlot.Y())

	case systems.InputAction_DropItem:
		iss.player.Inventory().Drop(iss.selectedInventorySlot.XY())
	case systems.InputAction_Menu_HighlightUp:
		if iss.selectedInventorySlot.Y() == 0 {
			break
		}

		iss.selectedInventorySlot = iss.selectedInventorySlot.WithOffset(0, -1)
		iss.inventoryMenu.SelectSlot(iss.selectedInventorySlot.XY())
	case systems.InputAction_Menu_HighlightDown:
		if iss.selectedInventorySlot.Y() == iss.player.Inventory().Shape().Height()-1 {
			break
		}

		iss.selectedInventorySlot = iss.selectedInventorySlot.WithOffset(0, +1)
		iss.inventoryMenu.SelectSlot(iss.selectedInventorySlot.XY())
	case systems.InputAction_Menu_HighlightLeft:
		if iss.selectedInventorySlot.X() == 0 {
			break
		}

		iss.selectedInventorySlot = iss.selectedInventorySlot.WithOffset(-1, 0)
		iss.inventoryMenu.SelectSlot(iss.selectedInventorySlot.XY())
	case systems.InputAction_Menu_HighlightRight:
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
