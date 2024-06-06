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

	selectedInventorySlot engine.Position
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

	menu.playerItems = engine.CreateDrawingInstructions(func(v views.View) {
		for y := range playerInventory.Shape().Height() {
			for x := range playerInventory.Shape().Width() {
				item := playerInventory.ItemAt(x, y)
				isHighlighted := x == menu.selectedInventorySlot.X() && y == menu.selectedInventorySlot.Y()

				if item == nil {

					if isHighlighted {
						ui.CreateSingleLineUILabel(
							menu.inventoryGrid.Position().X()+1+x*4,
							menu.inventoryGrid.Position().Y()+1+y*2,
							"   ",
							highlightStyle,
						).Draw(v)
					}

					continue
				}

				style := item.Style()

				if isHighlighted {
					style = highlightStyle
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
		item := playerInventory.ItemAt(menu.selectedInventorySlot.XY())

		if item == nil {
			return
		}

		ui.CreateUIItem(x+2, y+14, item, style).Draw(v)
	})

	return menu
}

func (pim *PlayerInventoryMenu) drawItemSlot(screenX, screenY int, item model.Item_V2, style tcell.Style, v views.View) {
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
}
