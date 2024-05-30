package input

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type Context string

const (
	InputContext_Play      = "play"
	InputContext_Menu      = "menu"
	InputContext_Inventory = "inventory"
)

type InputKey string

func InputKeyOf(context Context, mod tcell.ModMask, key tcell.Key, r rune) InputKey {
	return InputKey(fmt.Sprintf("%v-%v-%v-%v", context, mod, key, r))
}

type InputAction int

const (
	InputAction_None InputAction = iota

	InputAction_MovePlayer_North
	InputAction_MovePlayer_South
	InputAction_MovePlayer_East
	InputAction_MovePlayer_West

	InputAction_Interact
	InputAction_OpenInventory
	InputAction_PickUpItem
	InputAction_OpenLogs
	InputAction_DropItem
	InputAction_EquipItem
	InputAction_UnequipItem

	InputAction_PauseGame

	InputAction_Menu_HighlightDown
	InputAction_Menu_HighlightUp
	InputAction_Menu_HighlightLeft
	InputAction_Menu_HighlightRight

	InputAction_Menu_Select

	InputAction_Menu_Exit
)

type InputSystem struct {
	keyBindings map[InputKey]InputAction

	nextAction InputAction
}

func CreateInputSystemWithDefaultBindings() *InputSystem {
	return &InputSystem{
		keyBindings: map[InputKey]InputAction{
			InputKeyOf(InputContext_Play, 0, tcell.KeyUp, 0):          InputAction_MovePlayer_North,
			InputKeyOf(InputContext_Play, 0, tcell.KeyDown, 0):        InputAction_MovePlayer_South,
			InputKeyOf(InputContext_Play, 0, tcell.KeyLeft, 0):        InputAction_MovePlayer_West,
			InputKeyOf(InputContext_Play, 0, tcell.KeyRight, 0):       InputAction_MovePlayer_East,
			InputKeyOf(InputContext_Play, 0, tcell.KeyEsc, 0):         InputAction_PauseGame,
			InputKeyOf(InputContext_Play, 0, tcell.KeyRune, 'i'):      InputAction_OpenInventory,
			InputKeyOf(InputContext_Play, 0, tcell.KeyRune, 'l'):      InputAction_OpenLogs,
			InputKeyOf(InputContext_Play, 0, tcell.KeyRune, 'e'):      InputAction_Interact,
			InputKeyOf(InputContext_Play, 0, tcell.KeyRune, 'p'):      InputAction_PickUpItem,
			InputKeyOf(InputContext_Menu, 0, tcell.KeyESC, 0):         InputAction_Menu_Exit,
			InputKeyOf(InputContext_Menu, 0, tcell.KeyLeft, 0):        InputAction_Menu_HighlightLeft,
			InputKeyOf(InputContext_Menu, 0, tcell.KeyRight, 0):       InputAction_Menu_HighlightRight,
			InputKeyOf(InputContext_Menu, 0, tcell.KeyUp, 0):          InputAction_Menu_HighlightUp,
			InputKeyOf(InputContext_Menu, 0, tcell.KeyDown, 0):        InputAction_Menu_HighlightDown,
			InputKeyOf(InputContext_Menu, 0, tcell.KeyCR, 13):         InputAction_Menu_Select,
			InputKeyOf(InputContext_Inventory, 0, tcell.KeyESC, 0):    InputAction_Menu_Exit,
			InputKeyOf(InputContext_Inventory, 0, tcell.KeyRune, 'i'): InputAction_Menu_Exit,
			InputKeyOf(InputContext_Inventory, 0, tcell.KeyRune, 'e'): InputAction_EquipItem,
			InputKeyOf(InputContext_Inventory, 0, tcell.KeyRune, 'd'): InputAction_DropItem,
			InputKeyOf(InputContext_Inventory, 0, tcell.KeyLeft, 0):   InputAction_Menu_HighlightLeft,
			InputKeyOf(InputContext_Inventory, 0, tcell.KeyRight, 0):  InputAction_Menu_HighlightRight,
			InputKeyOf(InputContext_Inventory, 0, tcell.KeyUp, 0):     InputAction_Menu_HighlightUp,
			InputKeyOf(InputContext_Inventory, 0, tcell.KeyDown, 0):   InputAction_Menu_HighlightDown,
		},
	}
}

func (kb *InputSystem) ImportBindings(imports map[InputKey]InputAction) {
	kb.keyBindings = imports
}

func (kb *InputSystem) ExportBindings() map[InputKey]InputAction {
	return kb.keyBindings
}

func (kb *InputSystem) Bind(key InputKey, action InputAction) {
	kb.keyBindings[key] = action
}

func (kb *InputSystem) Input(context Context, ev *tcell.EventKey) {
	kb.nextAction = kb.keyBindings[InputKeyOf(context, ev.Modifiers(), ev.Key(), ev.Rune())]
}

func (kb *InputSystem) NextAction() (nextAction InputAction) {
	nextAction = kb.nextAction

	kb.nextAction = InputAction_None

	return
}
