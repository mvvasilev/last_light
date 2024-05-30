package ui

import (
	"fmt"
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/input"
	"mvvasilev/last_light/game/item"
	"mvvasilev/last_light/game/rpg"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	"github.com/google/uuid"
)

type UIBasicItem struct {
	id uuid.UUID

	item item.Item

	window   UIWindow
	itemName UILabel

	engine.Positioned
	engine.Sized
}

func CreateUIBasicItem(x, y int, item item.Item, style tcell.Style) *UIBasicItem {

	name, nameStyle := item.Name()

	return &UIBasicItem{
		id:         uuid.New(),
		item:       item,
		window:     *CreateWindow(x, y, 33, 8, "Item", style),
		itemName:   *CreateSingleLineUILabel(x+1, y+1, name, nameStyle),
		Positioned: engine.WithPosition(engine.PositionAt(x, y)),
		Sized:      engine.WithSize(engine.SizeOf(33, 8)),
	}
}

func (uibi *UIBasicItem) Input(e *tcell.EventKey) {
}

func (uibi *UIBasicItem) UniqueId() uuid.UUID {
	return uibi.id
}

func (uibi *UIBasicItem) Draw(v views.View) {
	uibi.window.Draw(v)
	uibi.itemName.Draw(v)
}

type UIRPGItem struct {
	id uuid.UUID

	item rpg.RPGItem

	window   UIWindow
	itemName UILabel

	engine.Positioned
	engine.Sized
}

func CreateUIRPGItem(x, y int, item rpg.RPGItem, style tcell.Style) *UIRPGItem {

	name, nameStyle := item.Name()

	return &UIRPGItem{
		id:         uuid.New(),
		item:       item,
		window:     *CreateWindow(x, y, 33, 8, "Item", style),
		itemName:   *CreateSingleLineUILabel(x+1, y+1, name, nameStyle),
		Positioned: engine.WithPosition(engine.PositionAt(x, y)),
		Sized:      engine.WithSize(engine.SizeOf(33, 8)),
	}
}

func (uiri *UIRPGItem) Input(inputAction input.InputAction) {
}

func (uiri *UIRPGItem) UniqueId() uuid.UUID {
	return uiri.id
}

func (uiri *UIRPGItem) Draw(v views.View) {
	uiri.window.Draw(v)
	uiri.itemName.Draw(v)

	statModifiers := uiri.item.Modifiers()

	x, y := uiri.itemName.Position().XY()
	y++

	for i, sm := range statModifiers {

		drawRPGItemStatModifier(x, y, tcell.StyleDefault, v, &sm)

		x += 9 + 2 // each stat is 9 characters long, with 2 characters separating the stats

		// Only 3 stats per line
		if i > 0 && (i+1)%3 == 0 {
			x = uiri.itemName.Position().X()
			y++
		}
	}
}

func drawRPGItemStatModifier(x, y int, style tcell.Style, view views.View, sm *rpg.StatModifier) {

	// 5 characters per stat name
	// 1 separating character
	// 3 characters for bonus ( including sign, modifiers are limited to -99 and +99)

	const SEPARATING_CHARACTER rune = ':'

	switch sm.Stat {
	case rpg.Stat_Attributes_Strength:
		engine.DrawText(x, y, "STR", style, view)
	case rpg.Stat_Attributes_Dexterity:
		engine.DrawText(x, y, "DEX", style, view)
	case rpg.Stat_Attributes_Intelligence:
		engine.DrawText(x, y, "INT", style, view)
	case rpg.Stat_Attributes_Constitution:
		engine.DrawText(x, y, "CON", style, view)
	case rpg.Stat_PhysicalPrecisionBonus:
		engine.DrawText(x, y, "pPrcs", style, view)
	case rpg.Stat_EvasionBonus:
		engine.DrawText(x, y, "Evasn", style, view)
	case rpg.Stat_MagicPrecisionBonus:
		engine.DrawText(x, y, "mPrcs", style, view)
	case rpg.Stat_TotalPrecisionBonus:
		engine.DrawText(x, y, "tPrcs", style, view)
	case rpg.Stat_DamageBonus_Physical_Unarmed:
		engine.DrawText(x, y, "Unrmd", style, view)
	case rpg.Stat_DamageBonus_Physical_Slashing:
		engine.DrawText(x, y, "Slshn", style, view)
	case rpg.Stat_DamageBonus_Physical_Piercing:
		engine.DrawText(x, y, "Prcng", style, view)
	case rpg.Stat_DamageBonus_Physical_Bludgeoning:
		engine.DrawText(x, y, "Bldgn", style, view)
	case rpg.Stat_DamageBonus_Magic_Fire:
		engine.DrawText(x, y, "Fire", style, view)
	case rpg.Stat_DamageBonus_Magic_Cold:
		engine.DrawText(x, y, "Cold", style, view)
	case rpg.Stat_DamageBonus_Magic_Necrotic:
		engine.DrawText(x, y, "Ncrtc", style, view)
	case rpg.Stat_DamageBonus_Magic_Thunder:
		engine.DrawText(x, y, "Thndr", style, view)
	case rpg.Stat_DamageBonus_Magic_Acid:
		engine.DrawText(x, y, "Acid", style, view)
	case rpg.Stat_DamageBonus_Magic_Poison:
		engine.DrawText(x, y, "Poisn", style, view)
	case rpg.Stat_MaxHealthBonus:
		engine.DrawText(x, y, "maxHP", style, view)
	default:
	}

	view.SetContent(x+5, y, SEPARATING_CHARACTER, nil, style)

	if sm.Bonus < 0 {
		engine.DrawText(x+6, y, fmt.Sprintf("-%02d", -sm.Bonus), tcell.StyleDefault.Foreground(tcell.ColorIndianRed), view)
	} else {
		engine.DrawText(x+6, y, fmt.Sprintf("+%02d", sm.Bonus), tcell.StyleDefault.Foreground(tcell.ColorLime), view)
	}
}
