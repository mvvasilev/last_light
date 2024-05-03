package state

import (
	"fmt"
	"mvvasilev/last_light/game/model"
	"mvvasilev/last_light/render"
	"mvvasilev/last_light/ui"
	"mvvasilev/last_light/util"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
)

type InventoryScreenState struct {
	prevState PausableState
	exitMenu  bool

	inventoryMenu  *ui.UIWindow
	armourLabel    *ui.UILabel
	armourGrid     *render.Grid
	leftHandLabel  *ui.UILabel
	leftHandBox    render.Rectangle
	rightHandLabel *ui.UILabel
	rightHandBox   render.Rectangle
	inventoryGrid  *render.Grid
	playerItems    *render.ArbitraryDrawable
	selectedItem   *render.ArbitraryDrawable
	help           *ui.UILabel

	player *model.Player

	moveInventorySlotDirection model.Direction
	selectedInventorySlot      util.Position
	dropSelectedInventorySlot  bool
}

func CreateInventoryScreenState(player *model.Player, prevState PausableState) *InventoryScreenState {
	iss := new(InventoryScreenState)

	iss.prevState = prevState
	iss.player = player
	iss.exitMenu = false
	iss.selectedInventorySlot = util.PositionAt(0, 0)

	iss.inventoryMenu = ui.CreateWindow(43, 0, 37, 24, "INVENTORY", tcell.StyleDefault)

	iss.armourLabel = ui.CreateSingleLineUILabel(58, 1, "ARMOUR", tcell.StyleDefault)

	iss.armourGrid = render.CreateGrid(
		53, 2, 3, 1, 4, 1, '┌', '─', '┬', '┐', '│', ' ', '│', '│', '├', '─', '┼', '┤', '└', '─', '┴', '┘', tcell.StyleDefault, tcell.StyleDefault.Background(tcell.ColorDarkSlateGray),
	)

	iss.leftHandLabel = ui.CreateUILabel(
		46, 1, 5, 1, "OFF", tcell.StyleDefault,
	)

	iss.leftHandBox = render.CreateRectangle(
		45, 2, 5, 3,
		'┌', '─', '┐',
		'│', ' ', '│',
		'└', '─', '┘',
		false, true,
		tcell.StyleDefault,
	)

	iss.rightHandLabel = ui.CreateUILabel(
		74, 1, 5, 1, "DOM", tcell.StyleDefault,
	)

	iss.rightHandBox = render.CreateRectangle(
		73, 2, 5, 3,
		'┌', '─', '┐',
		'│', ' ', '│',
		'└', '─', '┘',
		false, true,
		tcell.StyleDefault,
	)

	iss.inventoryGrid = render.CreateGrid(
		45, 5, 3, 1, 8, 4, '┌', '─', '┬', '┐', '│', ' ', '│', '│', '├', '─', '┼', '┤', '└', '─', '┴', '┘', tcell.StyleDefault, tcell.StyleDefault.Background(tcell.ColorDarkSlateGray),
	)

	iss.playerItems = render.CreateDrawingInstructions(func(v views.View) {
		for y := range player.Inventory().Shape().Height() {
			for x := range player.Inventory().Shape().Width() {
				item := player.Inventory().ItemAt(x, y)
				isHighlighted := x == iss.selectedInventorySlot.X() && y == iss.selectedInventorySlot.Y()

				if item == nil {

					if isHighlighted {
						ui.CreateSingleLineUILabel(45+1+x*4, 5+1+y*2, "   ", tcell.StyleDefault.Background(tcell.ColorDarkSlateGray)).Draw(v)
					}

					continue
				}

				style := item.Type().Style()

				if isHighlighted {
					style = style.Background(tcell.ColorDarkSlateGray)
				}

				ui.CreateSingleLineUILabel(45+1+x*4, 5+y*2, fmt.Sprintf("%03d", item.Quantity()), style).Draw(v)
				ui.CreateSingleLineUILabel(45+1+x*4, 5+1+y*2, item.Type().Icon(), style).Draw(v)
			}
		}
	})

	iss.selectedItem = render.CreateDrawingInstructions(func(v views.View) {
		ui.CreateWindow(45, 14, 33, 8, "ITEM", tcell.StyleDefault).Draw(v)

		item := player.Inventory().ItemAt(iss.selectedInventorySlot.XY())

		if item == nil {
			return
		}

		ui.CreateSingleLineUILabel(46, 15, fmt.Sprintf("Name: %v", item.Name()), tcell.StyleDefault).Draw(v)
		ui.CreateSingleLineUILabel(46, 16, fmt.Sprintf("Desc: %v", item.Description()), tcell.StyleDefault).Draw(v)
	})

	iss.help = ui.CreateSingleLineUILabel(45, 22, "hjkl - move, x - drop, e - equip", tcell.StyleDefault)

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

		iss.inventoryGrid.Highlight(iss.selectedInventorySlot)
		iss.moveInventorySlotDirection = model.DirectionNone
	}

	return iss
}

func (iss *InventoryScreenState) CollectDrawables() []render.Drawable {
	drawables := append(
		iss.prevState.CollectDrawables(),
		iss.inventoryMenu,
		iss.armourLabel,
		iss.armourGrid,
		iss.leftHandLabel,
		iss.leftHandBox,
		iss.rightHandLabel,
		iss.rightHandBox,
		iss.inventoryGrid,
		iss.playerItems,
		iss.selectedItem,
		iss.help,
	)

	return drawables
}
