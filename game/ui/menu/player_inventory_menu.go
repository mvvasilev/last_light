package menu

import (
	"fmt"
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/model"
	"mvvasilev/last_light/game/ui"
	"mvvasilev/last_light/util"

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
	help           *ui.UILabel

	selectedInventorySlot util.Position
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

				style := item.Type().Style()

				if isHighlighted {
					style = highlightStyle
				}

				ui.CreateSingleLineUILabel(
					menu.inventoryGrid.Position().X()+1+x*4,
					menu.inventoryGrid.Position().Y()+y*2,
					fmt.Sprintf("%03d", item.Quantity()),
					style,
				).Draw(v)

				ui.CreateSingleLineUILabel(
					menu.inventoryGrid.Position().X()+1+x*4,
					menu.inventoryGrid.Position().Y()+1+y*2,
					item.Type().Icon(),
					style,
				).Draw(v)
			}
		}
	})

	menu.selectedItem = engine.CreateDrawingInstructions(func(v views.View) {
		ui.CreateWindow(x+2, y+14, 33, 8, "ITEM", style).Draw(v)

		item := playerInventory.ItemAt(menu.selectedInventorySlot.XY())

		if item == nil {
			return
		}

		ui.CreateSingleLineUILabel(x+3, y+15, fmt.Sprintf("Name: %v", item.Name()), style).Draw(v)
		ui.CreateSingleLineUILabel(x+3, y+16, fmt.Sprintf("Desc: %v", item.Description()), style).Draw(v)
	})

	menu.help = ui.CreateSingleLineUILabel(x+2, y+22, "hjkl - move, x - drop, e - equip", style)

	return menu
}

func (pim *PlayerInventoryMenu) MoveTo(x int, y int) {

}

func (pim *PlayerInventoryMenu) Position() util.Position {
	return pim.inventoryMenu.Position()
}

func (pim *PlayerInventoryMenu) Size() util.Size {
	return pim.inventoryMenu.Size()
}

func (pim *PlayerInventoryMenu) Input(e *tcell.EventKey) {

}

func (pim *PlayerInventoryMenu) UniqueId() uuid.UUID {
	return pim.inventoryMenu.UniqueId()
}

func (pim *PlayerInventoryMenu) SelectSlot(x, y int) {
	pim.inventoryGrid.Unhighlight()
	pim.selectedInventorySlot = util.PositionAt(x, y)
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
	pim.help.Draw(v)
}
