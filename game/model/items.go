package model

import (
	"fmt"
	"mvvasilev/last_light/engine"

	"github.com/gdamore/tcell/v2"
)

type ItemType int

const (
	// Consumables
	ItemType_Fish ItemType = iota
	ItemType_SmallHealthPotion
	ItemType_HealthPotion
	ItemType_LargeHealthPotion

	// Weapons
	ItemType_Bow
	ItemType_Longsword
	ItemType_Club
	ItemType_Dagger
	ItemType_Handaxe
	ItemType_Javelin
	ItemType_LightHammer
	ItemType_Mace
	ItemType_Sickle
	ItemType_Spear
	ItemType_Quarterstaff

	// Armour

	// Special
)

func Item_Fish() Item {
	return createBaseItem(
		ItemType_Fish,
		'>',
		"»o>",
		tcell.StyleDefault,
		item_WithQuantity(1, 32),
		item_WithName("Fish", tcell.StyleDefault),
		item_WithDescription("On use heals for 1d4", tcell.StyleDefault),
		item_WithUsable(
			func(e Entity) bool {
				return e.HealthData() != nil
			},
			func(log *engine.GameEventLog, d *Dungeon, e Entity) {
				damageable := e.HealthData()

				if damageable != nil {
					healAmt := RollD4(1)
					damageable.Health = engine.LimitAdd(damageable.Health, healAmt, damageable.MaxHealth)

					name := "Entity"

					if e.Named() != nil {
						name = e.Named().Name
					}

					log.Log(fmt.Sprintf("%s heals for %d HP", name, healAmt))
				}
			},
		),
	)
}

func Item_SmallHealthPotion() Item {
	return createBaseItem(
		ItemType_SmallHealthPotion,
		'ó',
		" Ô ",
		tcell.StyleDefault.Foreground(tcell.ColorRed),
		item_WithQuantity(1, 3),
		item_WithName("Small Health Potion", tcell.StyleDefault),
		item_WithDescription("On use heals for 2d6", tcell.StyleDefault),
		item_WithUsable(
			func(e Entity) bool {
				return e.HealthData() != nil
			},
			func(log *engine.GameEventLog, d *Dungeon, e Entity) {
				damageable := e.HealthData()

				if damageable != nil {
					healAmt := RollD6(2)
					damageable.Health = engine.LimitAdd(damageable.Health, healAmt, damageable.MaxHealth)

					name := "Entity"

					if e.Named() != nil {
						name = e.Named().Name
					}

					log.Log(fmt.Sprintf("%s heals for %d HP", name, healAmt))
				}
			},
		),
	)
}

func Item_HealthPotion() Item {
	return createBaseItem(
		ItemType_HealthPotion,
		'ó',
		" Ô ",
		tcell.StyleDefault.Foreground(tcell.ColorRed),
		item_WithQuantity(1, 2),
		item_WithName("Health Potion", tcell.StyleDefault),
		item_WithDescription("On use heals for 3d6", tcell.StyleDefault),
		item_WithUsable(
			func(e Entity) bool {
				return e.HealthData() != nil
			},
			func(log *engine.GameEventLog, d *Dungeon, e Entity) {
				damageable := e.HealthData()

				if damageable != nil {
					healAmt := RollD6(3)
					damageable.Health = engine.LimitAdd(damageable.Health, healAmt, damageable.MaxHealth)

					name := "Entity"

					if e.Named() != nil {
						name = e.Named().Name
					}

					log.Log(fmt.Sprintf("%s heals for %d HP", name, healAmt))
				}
			},
		),
	)
}

func Item_LargeHealthPotion() Item {
	return createBaseItem(
		ItemType_LargeHealthPotion,
		'ó',
		" Ô ",
		tcell.StyleDefault.Foreground(tcell.ColorRed),
		item_WithQuantity(1, 1),
		item_WithName("Large Health Potion", tcell.StyleDefault),
		item_WithDescription("On use heals for 4d6", tcell.StyleDefault),
		item_WithUsable(
			func(e Entity) bool {
				return e.HealthData() != nil
			},
			func(log *engine.GameEventLog, d *Dungeon, e Entity) {
				damageable := e.HealthData()

				if damageable != nil {
					healAmt := RollD6(4)
					damageable.Health = engine.LimitAdd(damageable.Health, healAmt, damageable.MaxHealth)

					name := "Entity"

					if e.Named() != nil {
						name = e.Named().Name
					}

					log.Log(fmt.Sprintf("%s heals for %d HP", name, healAmt))
				}
			},
		),
	)
}

func Item_Bow() Item {
	return createBaseItem(
		ItemType_Bow,
		')',
		" |)",
		tcell.StyleDefault.Foreground(tcell.ColorBrown),
		item_WithQuantity(1, 1),
		item_WithName("Bow", tcell.StyleDefault),
		item_WithDescription("Deals 1d8 Piercing damage", tcell.StyleDefault),
		item_WithDamaging(true, func() (damage int, dmgType DamageType) {
			return RollD8(1), DamageType_Physical_Piercing
		}),
		item_WithMetaTypes([]ItemMetaType{MetaItemType_Physical_Weapon, MetaItemType_Weapon}),
		item_WithEquippable(EquippedSlotDominantHand),
	)
}

func Item_Longsword() Item {
	return createBaseItem(
		ItemType_Longsword,
		'/',
		"╪══",
		tcell.StyleDefault.Foreground(tcell.ColorSilver),
		item_WithQuantity(1, 1),
		item_WithName("Longsword", tcell.StyleDefault),
		item_WithDescription("Deals 1d8 Slashing damage", tcell.StyleDefault),
		item_WithDamaging(false, func() (damage int, dmgType DamageType) {
			return RollD8(1), DamageType_Physical_Slashing
		}),
		item_WithMetaTypes([]ItemMetaType{MetaItemType_Physical_Weapon, MetaItemType_Weapon}),
		item_WithEquippable(EquippedSlotDominantHand),
	)
}

func Item_Club() Item {
	return createBaseItem(
		ItemType_Club,
		'!',
		"-══",
		tcell.StyleDefault.Foreground(tcell.ColorSaddleBrown),
		item_WithQuantity(1, 1),
		item_WithName("Club", tcell.StyleDefault),
		item_WithDescription("Deals 1d8 Bludgeoning damage", tcell.StyleDefault),
		item_WithDamaging(false, func() (damage int, dmgType DamageType) {
			return RollD8(1), DamageType_Physical_Bludgeoning
		}),
		item_WithMetaTypes([]ItemMetaType{MetaItemType_Physical_Weapon, MetaItemType_Weapon}),
		item_WithEquippable(EquippedSlotDominantHand),
	)
}

func Item_Dagger() Item {
	return createBaseItem(
		ItemType_Dagger,
		'-',
		" +─",
		tcell.StyleDefault.Foreground(tcell.ColorSilver),
		item_WithQuantity(1, 1),
		item_WithName("Dagger", tcell.StyleDefault),
		item_WithDescription("Deals 1d6 Piercing damage", tcell.StyleDefault),
		item_WithDamaging(false, func() (damage int, dmgType DamageType) {
			return RollD6(1), DamageType_Physical_Piercing
		}),
		item_WithMetaTypes([]ItemMetaType{MetaItemType_Physical_Weapon, MetaItemType_Weapon}),
		item_WithEquippable(EquippedSlotDominantHand),
	)
}

func Item_Handaxe() Item {
	return createBaseItem(
		ItemType_Handaxe,
		'¶',
		" ─╗",
		tcell.StyleDefault.Foreground(tcell.ColorSilver),
		item_WithQuantity(1, 1),
		item_WithName("Handaxe", tcell.StyleDefault),
		item_WithDescription("Deals 1d6 Slashing damage", tcell.StyleDefault),
		item_WithDamaging(false, func() (damage int, dmgType DamageType) {
			return RollD6(1), DamageType_Physical_Piercing
		}),
		item_WithMetaTypes([]ItemMetaType{MetaItemType_Physical_Weapon, MetaItemType_Weapon}),
		item_WithEquippable(EquippedSlotDominantHand),
	)
}

func Item_Javelin() Item {
	return createBaseItem(
		ItemType_Javelin,
		'Î',
		" ─>",
		tcell.StyleDefault.Foreground(tcell.ColorSilver),
		item_WithQuantity(1, 1),
		item_WithName("Javelin", tcell.StyleDefault),
		item_WithDescription("Deals 1d6 Piercing damage", tcell.StyleDefault),
		item_WithDamaging(false, func() (damage int, dmgType DamageType) {
			return RollD6(1), DamageType_Physical_Piercing
		}),
		item_WithMetaTypes([]ItemMetaType{MetaItemType_Physical_Weapon, MetaItemType_Weapon}),
		item_WithEquippable(EquippedSlotDominantHand),
	)
}

func Item_LightHammer() Item {
	return createBaseItem(
		ItemType_LightHammer,
		'i',
		" ─0",
		tcell.StyleDefault.Foreground(tcell.ColorSilver),
		item_WithQuantity(1, 1),
		item_WithName("Light Hammer", tcell.StyleDefault),
		item_WithDescription("Deals 1d6 Bludgeoning damage", tcell.StyleDefault),
		item_WithDamaging(false, func() (damage int, dmgType DamageType) {
			return RollD6(1), DamageType_Physical_Bludgeoning
		}),
		item_WithMetaTypes([]ItemMetaType{MetaItemType_Physical_Weapon, MetaItemType_Weapon}),
		item_WithEquippable(EquippedSlotDominantHand),
	)
}

func Item_Mace() Item {
	return createBaseItem(
		ItemType_Mace,
		'i',
		" ─¤",
		tcell.StyleDefault.Foreground(tcell.ColorSilver),
		item_WithQuantity(1, 1),
		item_WithName("Mace", tcell.StyleDefault),
		item_WithDescription("Deals 1d6 Bludgeoning damage", tcell.StyleDefault),
		item_WithDamaging(false, func() (damage int, dmgType DamageType) {
			return RollD6(1), DamageType_Physical_Bludgeoning
		}),
		item_WithMetaTypes([]ItemMetaType{MetaItemType_Physical_Weapon, MetaItemType_Weapon}),
		item_WithEquippable(EquippedSlotDominantHand),
	)
}

func Item_Quarterstaff() Item {
	return createBaseItem(
		ItemType_Quarterstaff,
		'|',
		"───",
		tcell.StyleDefault.Foreground(tcell.ColorSilver),
		item_WithQuantity(1, 1),
		item_WithName("Quarterstaff", tcell.StyleDefault),
		item_WithDescription("Deals 1d6 Bludgeoning damage", tcell.StyleDefault),
		item_WithDamaging(false, func() (damage int, dmgType DamageType) {
			return RollD6(1), DamageType_Physical_Bludgeoning
		}),
		item_WithMetaTypes([]ItemMetaType{MetaItemType_Physical_Weapon, MetaItemType_Weapon}),
		item_WithEquippable(EquippedSlotDominantHand),
	)
}

func Item_Sickle() Item {
	return createBaseItem(
		ItemType_Sickle,
		'?',
		" ─U",
		tcell.StyleDefault.Foreground(tcell.ColorSilver),
		item_WithQuantity(1, 1),
		item_WithName("Sickle", tcell.StyleDefault),
		item_WithDescription("Deals 1d6 Slashing damage", tcell.StyleDefault),
		item_WithDamaging(false, func() (damage int, dmgType DamageType) {
			return RollD6(1), DamageType_Physical_Slashing
		}),
		item_WithMetaTypes([]ItemMetaType{MetaItemType_Physical_Weapon, MetaItemType_Weapon}),
		item_WithEquippable(EquippedSlotDominantHand),
	)
}

func Item_Spear() Item {
	return createBaseItem(
		ItemType_Spear,
		'Î',
		"──>",
		tcell.StyleDefault.Foreground(tcell.ColorSilver),
		item_WithQuantity(1, 1),
		item_WithName("Spear", tcell.StyleDefault),
		item_WithDescription("Deals 1d8 Piercing damage", tcell.StyleDefault),
		item_WithDamaging(false, func() (damage int, dmgType DamageType) {
			return RollD8(1), DamageType_Physical_Piercing
		}),
		item_WithMetaTypes([]ItemMetaType{MetaItemType_Physical_Weapon, MetaItemType_Weapon}),
		item_WithEquippable(EquippedSlotDominantHand),
	)
}

// func ItemTypeGold() ItemType {
// 	return &BasicItemType{
// 		id:             1,
// 		name:           "Gold",
// 		description:    "Not all those who wander are lost",
// 		tileIcon:       '¤',
// 		itemIcon:       " ¤ ",
// 		equippableSlot: EquippedSlotNone,
// 		style:          tcell.StyleDefault.Foreground(tcell.ColorGoldenrod),
// 		maxStack:       255,
// 	}
// }

// func ItemTypeArrow() ItemType {
// 	return &BasicItemType{
// 		id:             2,
// 		name:           "Arrow",
// 		description:    "Ammunition for a bow",
// 		tileIcon:       '-',
// 		itemIcon:       "»->",
// 		equippableSlot: EquippedSlotNone,
// 		style:          tcell.StyleDefault.Foreground(tcell.ColorGoldenrod),
// 		maxStack:       32,
// 	}
// }

// func ItemTypeKey() ItemType {
// 	return &BasicItemType{
// 		id:             3,
// 		name:           "Key",
// 		description:    "Indispensable for unlocking things",
// 		tileIcon:       '¬',
// 		itemIcon:       " o╖",
// 		equippableSlot: EquippedSlotNone,
// 		style:          tcell.StyleDefault.Foreground(tcell.ColorDarkGoldenrod),
// 		maxStack:       1,
// 	}
// }
