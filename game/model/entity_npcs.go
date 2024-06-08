package model

import (
	"mvvasilev/last_light/engine"

	"github.com/gdamore/tcell/v2"
)

type specialItemType ItemType

const (
	ImpClaws specialItemType = 100_000 + iota
)

func Entity_ArrowProjectile(startX, startY int, targetX, targetY int) Entity {
	return CreateEntity(
		WithName("Arrow"),
		WithPosition(engine.PositionAt(startX, startY)),
	)
}

func Entity_Imp(x, y int, behavior func(npc Entity) (complete, requeue bool)) Entity {
	return CreateEntity(
		WithName("Imp"),
		WithDescription("A fiery little creature"),
		WithHealthData(15, 15, false),
		WithPosition(engine.PositionAt(x, y)),
		WithPresentation('i', tcell.StyleDefault.Foreground(tcell.ColorDarkRed)),
		WithBehavior(110, behavior),
		WithStats(map[Stat]int{
			Stat_Attributes_Constitution: 5,
			Stat_Attributes_Dexterity:    10,
			Stat_Attributes_Strength:     5,
			Stat_Attributes_Intelligence: 7,

			Stat_ResistanceBonus_Magic_Fire: 5,
		}),
		WithInventory(BuildInventory(
			Inv_WithDominantHand(createBaseItem(
				ItemType(ImpClaws),
				'v', "|||",
				tcell.StyleDefault,
				item_WithName("Claws", tcell.StyleDefault),
				item_WithDamaging(false, func() (damage int, dmgType DamageType) {
					return RollD4(1), DamageType_Physical_Slashing
				}),
			)),
		)),
	)
}

func Entity_SkeletalKnight(x, y int, behavior func(npc Entity) (complete, requeue bool)) Entity {
	return CreateEntity(
		WithName("Skeletal Knight"),
		WithDescription("Rattling in the dark..."),
		WithHealthData(25, 25, false),
		WithPosition(engine.PositionAt(x, y)),
		WithPresentation('S', tcell.StyleDefault.Foreground(tcell.ColorAntiqueWhite)),
		WithBehavior(150, behavior),
		WithStats(map[Stat]int{
			Stat_Attributes_Constitution: 10,
			Stat_Attributes_Dexterity:    6,
			Stat_Attributes_Strength:     12,
			Stat_Attributes_Intelligence: 5,

			Stat_ResistanceBonus_Physical_Bludgeoning: -2,
		}),
		WithInventory(BuildInventory(
			Inv_WithDominantHand(Item_Longsword()),
		)),
		WithDropTable(map[int]ItemSupplier{
			9: ItemSupplierOf(Item_Longsword()),
			1: ItemSupplierOfGeneratedPrototype(Item_Longsword(), map[int]ItemRarity{1: ItemRarity_Legendary}),
		}),
	)
}

func Entity_SkeletalWarrior(x, y int, behavior func(npc Entity) (complete, requeue bool)) Entity {
	return CreateEntity(
		WithName("Skeletal Warrior"),
		WithDescription("Rattling in the dark..."),
		WithHealthData(25, 25, false),
		WithPosition(engine.PositionAt(x, y)),
		WithPresentation('S', tcell.StyleDefault.Foreground(tcell.ColorAntiqueWhite)),
		WithBehavior(150, behavior),
		WithStats(map[Stat]int{
			Stat_Attributes_Constitution: 10,
			Stat_Attributes_Dexterity:    6,
			Stat_Attributes_Strength:     12,
			Stat_Attributes_Intelligence: 5,

			Stat_ResistanceBonus_Physical_Bludgeoning: -2,
		}),
		WithInventory(BuildInventory(
			Inv_WithDominantHand(Item_Mace()),
		)),
		WithDropTable(map[int]ItemSupplier{
			9: ItemSupplierOf(Item_Mace()),
			1: ItemSupplierOfGeneratedPrototype(Item_Mace(), map[int]ItemRarity{1: ItemRarity_Legendary}),
		}),
	)
}
