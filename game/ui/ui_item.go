package ui

import (
	"fmt"
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/model"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	"github.com/google/uuid"
)

type UIItem struct {
	id uuid.UUID

	item model.Item

	window UIWindow

	engine.Positioned
	engine.Sized
}

func CreateUIItem(x, y int, item model.Item, style tcell.Style) *UIItem {
	return &UIItem{
		id:         uuid.New(),
		item:       item,
		window:     *CreateWindow(x, y, 33, 8, "Item", style),
		Positioned: engine.WithPosition(engine.PositionAt(x, y)),
		Sized:      engine.WithSize(engine.SizeOf(33, 8)),
	}
}

func (uibi *UIItem) Input(e *tcell.EventKey) {
}

func (uibi *UIItem) UniqueId() uuid.UUID {
	return uibi.id
}

func (uibi *UIItem) Draw(v views.View) {
	uibi.window.Draw(v)

	if uibi.item.Named() != nil {
		engine.DrawText(uibi.Position().X()+1, uibi.Position().Y()+1, uibi.item.Named().Name, uibi.item.Named().Style, v)
	}

	if uibi.item.Described() != nil {
		engine.DrawText(uibi.Position().X()+1, uibi.Position().Y()+2, uibi.item.Described().Description, uibi.item.Described().Style, v)
	}

	if uibi.item.StatModifier() == nil {
		return
	}

	statModifiers := uibi.item.StatModifier().StatModifiers

	originalX, y := uibi.Position().XY()
	x := originalX + 1
	y += 3

	for i, sm := range statModifiers {

		drawRPGItemStatModifier(x, y, tcell.StyleDefault, v, &sm)

		x += 9 + 2 // each stat is 9 characters long, with 2 characters separating the stats

		// Only 3 stats per line
		if i > 0 && (i+1)%3 == 0 {
			x = originalX + 1
			y++
		}
	}
}

func drawRPGItemStatModifier(x, y int, style tcell.Style, view views.View, sm *model.StatModifier) {

	// 5 characters per stat name
	// 1 separating character
	// 3 characters for bonus ( including sign, modifiers are limited to -99 and +99)

	const SEPARATING_CHARACTER rune = ':'

	switch sm.Stat {
	case model.Stat_Attributes_Strength:
		engine.DrawText(x, y, "STR", style, view)
	case model.Stat_Attributes_Dexterity:
		engine.DrawText(x, y, "DEX", style, view)
	case model.Stat_Attributes_Intelligence:
		engine.DrawText(x, y, "INT", style, view)
	case model.Stat_Attributes_Constitution:
		engine.DrawText(x, y, "CON", style, view)
	case model.Stat_PhysicalPrecisionBonus:
		engine.DrawText(x, y, "pPrcs", style, view)
	case model.Stat_EvasionBonus:
		engine.DrawText(x, y, "Evasn", style, view)
	case model.Stat_MagicPrecisionBonus:
		engine.DrawText(x, y, "mPrcs", style, view)
	case model.Stat_TotalPrecisionBonus:
		engine.DrawText(x, y, "tPrcs", style, view)
	case model.Stat_DamageBonus_Physical_Unarmed:
		engine.DrawText(x, y, "Unrmd", style, view)
	case model.Stat_DamageBonus_Physical_Slashing:
		engine.DrawText(x, y, "Slshn", style, view)
	case model.Stat_DamageBonus_Physical_Piercing:
		engine.DrawText(x, y, "Prcng", style, view)
	case model.Stat_DamageBonus_Physical_Bludgeoning:
		engine.DrawText(x, y, "Bldgn", style, view)
	case model.Stat_DamageBonus_Magic_Fire:
		engine.DrawText(x, y, "Fire", style, view)
	case model.Stat_DamageBonus_Magic_Cold:
		engine.DrawText(x, y, "Cold", style, view)
	case model.Stat_DamageBonus_Magic_Necrotic:
		engine.DrawText(x, y, "Ncrtc", style, view)
	case model.Stat_DamageBonus_Magic_Thunder:
		engine.DrawText(x, y, "Thndr", style, view)
	case model.Stat_DamageBonus_Magic_Acid:
		engine.DrawText(x, y, "Acid", style, view)
	case model.Stat_DamageBonus_Magic_Poison:
		engine.DrawText(x, y, "Poisn", style, view)
	case model.Stat_MaxHealthBonus:
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
