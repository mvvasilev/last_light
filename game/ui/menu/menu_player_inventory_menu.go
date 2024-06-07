package menu

import (
	"fmt"
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/model"
	"mvvasilev/last_light/game/systems"
	"mvvasilev/last_light/game/ui"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	"github.com/google/uuid"
)

type PlayerInventoryMenu struct {
	inventory *model.EquippedInventory

	inventoryMenu  *ui.UIWindow
	armourLabel    *ui.UILabel
	armourGrid     *engine.Grid
	leftHandLabel  *ui.UILabel
	leftHandBox    engine.Rectangle
	rightHandLabel *ui.UILabel
	rightHandBox   engine.Rectangle
	inventoryGrid  *engine.Grid
	playerItems    *engine.ArbitraryDrawable
	selectedItem   *engine.ArbitraryDrawable

	equipmentSlots *engine.ArbitraryDrawable

	selectedInventorySlot engine.Position

	highlightStyle tcell.Style
}

func CreatePlayerInventoryMenu(x, y int, playerInventory *model.EquippedInventory, style tcell.Style, highlightStyle tcell.Style) *PlayerInventoryMenu {
	menu := new(PlayerInventoryMenu)

	menu.inventory = playerInventory

	menu.inventoryMenu = ui.CreateWindow(x, y, 37, 24, "INVENTORY", style)

	menu.armourLabel = ui.CreateSingleLineUILabel(x+15, y+1, "ARMOUR", style)

	menu.armourGrid = engine.CreateGrid(
		x+10, y+2, 3, 1, 4, 1, '┌', '─', '┬', '┐', '│', ' ', '│', '│', '├', '─', '┼', '┤', '└', '─', '┴', '┘', style, highlightStyle, //style.Background(tcell.ColorDarkSlateGray),
	)

	menu.leftHandLabel = ui.CreateUILabel(
		x+3, y+1, 5, 1, "OFF", style,
	)

	menu.leftHandBox = engine.CreateRectangle(
		x+2, y+2, 5, 3,
		'┌', '─', '┐',
		'│', ' ', '│',
		'└', '─', '┘',
		false, true,
		style,
	)

	menu.rightHandLabel = ui.CreateUILabel(
		x+31, y+1, 5, 1, "DOM", style,
	)

	menu.rightHandBox = engine.CreateRectangle(
		x+30, y+2, 5, 3,
		'┌', '─', '┐',
		'│', ' ', '│',
		'└', '─', '┘',
		false, true,
		style,
	)

	menu.inventoryGrid = engine.CreateGrid(
		x+2, y+5, 3, 1, 8, 4, '┌', '─', '┬', '┐', '│', ' ', '│', '│', '├', '─', '┼', '┤', '└', '─', '┴', '┘', style, highlightStyle,
	)

	menu.highlightStyle = highlightStyle

	menu.playerItems = engine.CreateDrawingInstructions(func(v views.View) {
		for y := range menu.inventory.Shape().Height() {
			for x := range menu.inventory.Shape().Width() {
				item := menu.inventory.ItemAt(x, y)
				isHighlighted := x == menu.selectedInventorySlot.X() && y == menu.selectedInventorySlot.Y()

				if item == nil {

					if isHighlighted {
						ui.CreateSingleLineUILabel(
							menu.inventoryGrid.Position().X()+1+x*4,
							menu.inventoryGrid.Position().Y()+1+y*2,
							"   ",
							menu.highlightStyle,
						).Draw(v)
					}

					continue
				}

				style := item.Style()

				if isHighlighted {
					style = menu.highlightStyle
				}

				menu.drawItemSlot(
					menu.inventoryGrid.Position().X()+1+x*4,
					menu.inventoryGrid.Position().Y()+y*2,
					item,
					style,
					v,
				)
			}
		}
	})

	menu.selectedItem = engine.CreateDrawingInstructions(func(v views.View) {
		item := menu.inventory.ItemAt(menu.selectedInventorySlot.XY())

		if item == nil {
			return
		}

		ui.CreateUIItem(x+2, y+14, item, style).Draw(v)
	})

	menu.equipmentSlots = engine.CreateDrawingInstructions(func(v views.View) {
		drawEquipmentSlot(menu.rightHandBox.Position().X()+1, menu.rightHandBox.Position().Y(), menu.inventory.AtSlot(model.EquippedSlotDominantHand), false, menu.highlightStyle, v)
		drawEquipmentSlot(menu.leftHandBox.Position().X()+1, menu.leftHandBox.Position().Y(), menu.inventory.AtSlot(model.EquippedSlotOffhand), false, menu.highlightStyle, v)
		drawEquipmentSlot(x+10+1, y+3, menu.inventory.AtSlot(model.EquippedSlotHead), false, menu.highlightStyle, v)
		drawEquipmentSlot(x+10+4, y+3, menu.inventory.AtSlot(model.EquippedSlotChestplate), false, menu.highlightStyle, v)
		drawEquipmentSlot(x+10+7, y+3, menu.inventory.AtSlot(model.EquippedSlotLeggings), false, menu.highlightStyle, v)
		drawEquipmentSlot(x+10+10, y+3, menu.inventory.AtSlot(model.EquippedSlotShoes), false, menu.highlightStyle, v)
	})

	return menu
}

func drawEquipmentSlot(screenX, screenY int, item model.Item, highlighted bool, highlightStyle tcell.Style, v views.View) {
	if item == nil {
		return
	}

	style := item.Style()

	if highlighted {
		style = highlightStyle
	}

	ui.CreateSingleLineUILabel(
		screenX,
		screenY+1,
		item.Icon(),
		style,
	).Draw(v)
}

func (pim *PlayerInventoryMenu) drawItemSlot(screenX, screenY int, item model.Item, style tcell.Style, v views.View) {
	if item.Quantifiable() != nil {
		ui.CreateSingleLineUILabel(
			screenX,
			screenY,
			fmt.Sprintf("%03d", item.Quantifiable().CurrentQuantity),
			style,
		).Draw(v)
	}

	ui.CreateSingleLineUILabel(
		screenX,
		screenY+1,
		item.Icon(),
		style,
	).Draw(v)
}

func (pim *PlayerInventoryMenu) MoveTo(x int, y int) {

}

func (pim *PlayerInventoryMenu) Position() engine.Position {
	return pim.inventoryMenu.Position()
}

func (pim *PlayerInventoryMenu) Size() engine.Size {
	return pim.inventoryMenu.Size()
}

func (pim *PlayerInventoryMenu) Input(inputAction systems.InputAction) {

}

func (pim *PlayerInventoryMenu) UniqueId() uuid.UUID {
	return pim.inventoryMenu.UniqueId()
}

func (pim *PlayerInventoryMenu) SelectSlot(x, y int) {
	if pim.selectedInventorySlot.X() == x && pim.selectedInventorySlot.Y() == y {
		return
	}

	pim.inventoryGrid.Unhighlight()
	pim.selectedInventorySlot = engine.PositionAt(x, y)
	pim.inventoryGrid.Highlight(pim.selectedInventorySlot)
}

func (pim *PlayerInventoryMenu) Draw(v views.View) {
	pim.inventoryMenu.Draw(v)
	pim.armourLabel.Draw(v)
	pim.armourGrid.Draw(v)
	pim.leftHandLabel.Draw(v)
	pim.leftHandBox.Draw(v)
	pim.rightHandLabel.Draw(v)
	pim.rightHandBox.Draw(v)
	pim.inventoryGrid.Draw(v)
	pim.playerItems.Draw(v)
	pim.selectedItem.Draw(v)
	pim.equipmentSlots.Draw(v)
}
