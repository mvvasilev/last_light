package model

import (
	"mvvasilev/last_light/engine"

	"github.com/gdamore/tcell/v2"
)

type specialItemType ItemType

const (
	ImpClaws specialItemType = 100_000 + iota
)

func Entity_Imp(x, y int) Entity {
	return CreateEntity(
		WithName("Imp"),
		WithDescription("A fiery little creature"),
		WithHealthData(15, 15, false),
		WithPosition(engine.PositionAt(x, y)),
		WithPresentation('i', tcell.StyleDefault.Foreground(tcell.ColorDarkRed)),
		WithSpeed(11),
		WithStats(map[Stat]int{
			Stat_Attributes_Constitution: 5,
			Stat_Attributes_Dexterity:    10,
			Stat_Attributes_Strength:     5,
			Stat_Attributes_Intelligence: 7,
		}),
		WithInventory(BuildInventory(
			Inv_WithDominantHand(createBaseItem(
				ItemType(ImpClaws),
				'v', "|||",
				tcell.StyleDefault,
				item_WithName("Claws", tcell.StyleDefault),
				item_WithDamaging(func() (damage int, dmgType DamageType) {
					return RollD4(1), DamageType_Physical_Slashing
				}),
			)),
		)),
	)
}
