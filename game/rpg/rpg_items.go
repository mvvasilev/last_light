package rpg

import (
	"mvvasilev/last_light/game/item"

	"github.com/gdamore/tcell/v2"
)

type RPGItemMetaType int

const (
	MetaItemType_Physical_Weapon RPGItemMetaType = iota
	MetaItemType_Magic_Weapon
	MetaItemType_Weapon
	MetaItemType_Physical_Armour
	MetaItemType_Magic_Armour
	MetaItemType_Armour
)

type RPGItemType interface {
	RollDamage() func(victim, attacker RPGEntity) (damage int, dmgType DamageType)
	MetaTypes() []RPGItemMetaType

	item.ItemType
}

type RPGItem interface {
	Modifiers() []StatModifier

	item.Item
}

type BasicRPGItemType struct {
	damageRollFunc func(victim, attacker RPGEntity) (damage int, dmgType DamageType)

	metaTypes []RPGItemMetaType

	*item.BasicItemType
}

func (it *BasicRPGItemType) RollDamage() func(victim, attacker RPGEntity) (damage int, dmgType DamageType) {
	return it.damageRollFunc
}

func (it *BasicRPGItemType) MetaTypes() []RPGItemMetaType {
	return it.metaTypes
}

func ItemTypeBow() RPGItemType {
	return &BasicRPGItemType{
		damageRollFunc: func(victim, attacker RPGEntity) (damage int, dmgType DamageType) {
			// TODO: Ranged
			return RollD8(1), DamageType_Physical_Piercing
		},
		metaTypes: []RPGItemMetaType{MetaItemType_Physical_Weapon, MetaItemType_Weapon},
		BasicItemType: item.CreateBasicItemType(
			1000,
			"Bow",
			"To shoot arrows with",
			')',
			" |)",
			1,
			item.EquippedSlotDominantHand,
			tcell.StyleDefault.Foreground(tcell.ColorBrown),
		),
	}
}

func ItemTypeLongsword() RPGItemType {
	return &BasicRPGItemType{
		damageRollFunc: func(victim, attacker RPGEntity) (damage int, dmgType DamageType) {
			return RollD8(1), DamageType_Physical_Slashing
		},
		metaTypes: []RPGItemMetaType{MetaItemType_Physical_Weapon, MetaItemType_Weapon},
		BasicItemType: item.CreateBasicItemType(
			1001,
			"Longsword",
			"You know nothing.",
			'/',
			"╪══",
			1,
			item.EquippedSlotDominantHand,
			tcell.StyleDefault.Foreground(tcell.ColorSilver),
		),
	}
}

func ItemTypeClub() RPGItemType {
	return &BasicRPGItemType{
		damageRollFunc: func(victim, attacker RPGEntity) (damage int, dmgType DamageType) {
			return RollD8(1), DamageType_Physical_Bludgeoning
		},
		metaTypes: []RPGItemMetaType{MetaItemType_Physical_Weapon, MetaItemType_Weapon},
		BasicItemType: item.CreateBasicItemType(
			1002,
			"Club",
			"Bonk",
			'!',
			"-══",
			1,
			item.EquippedSlotDominantHand,
			tcell.StyleDefault.Foreground(tcell.ColorSaddleBrown),
		),
	}
}

func ItemTypeDagger() RPGItemType {
	return &BasicRPGItemType{
		damageRollFunc: func(victim, attacker RPGEntity) (damage int, dmgType DamageType) {
			return RollD6(1), DamageType_Physical_Piercing
		},
		metaTypes: []RPGItemMetaType{MetaItemType_Physical_Weapon, MetaItemType_Weapon},
		BasicItemType: item.CreateBasicItemType(
			1003,
			"Dagger",
			"Stabby, stabby",
			'-',
			" +─",
			1,
			item.EquippedSlotDominantHand,
			tcell.StyleDefault.Foreground(tcell.ColorSilver),
		),
	}
}

func ItemTypeHandaxe() RPGItemType {
	return &BasicRPGItemType{
		damageRollFunc: func(victim, attacker RPGEntity) (damage int, dmgType DamageType) {
			return RollD6(1), DamageType_Physical_Slashing
		},
		metaTypes: []RPGItemMetaType{MetaItemType_Physical_Weapon, MetaItemType_Weapon},
		BasicItemType: item.CreateBasicItemType(
			1004,
			"Handaxe",
			"Choppy, choppy",
			'¶',
			" ─╗",
			1,
			item.EquippedSlotDominantHand,
			tcell.StyleDefault.Foreground(tcell.ColorSilver),
		),
	}
}

func ItemTypeJavelin() RPGItemType {
	return &BasicRPGItemType{
		damageRollFunc: func(victim, attacker RPGEntity) (damage int, dmgType DamageType) {
			// TODO: Ranged
			return RollD6(1), DamageType_Physical_Piercing
		},
		metaTypes: []RPGItemMetaType{MetaItemType_Physical_Weapon, MetaItemType_Weapon},
		BasicItemType: item.CreateBasicItemType(
			1005,
			"Javelin",
			"Ranged pokey, pokey",
			'Î',
			" ─>",
			20,
			item.EquippedSlotDominantHand,
			tcell.StyleDefault.Foreground(tcell.ColorSilver),
		),
	}
}

func ItemTypeLightHammer() RPGItemType {
	return &BasicRPGItemType{
		damageRollFunc: func(victim, attacker RPGEntity) (damage int, dmgType DamageType) {
			return RollD6(1), DamageType_Physical_Bludgeoning
		},
		metaTypes: []RPGItemMetaType{MetaItemType_Physical_Weapon, MetaItemType_Weapon},
		BasicItemType: item.CreateBasicItemType(
			1006,
			"Handaxe",
			"Choppy, choppy",
			'¶',
			" ─╗",
			1,
			item.EquippedSlotDominantHand,
			tcell.StyleDefault.Foreground(tcell.ColorSilver),
		),
	}
}

func ItemTypeMace() RPGItemType {
	return &BasicRPGItemType{
		damageRollFunc: func(victim, attacker RPGEntity) (damage int, dmgType DamageType) {
			return RollD6(1), DamageType_Physical_Bludgeoning
		},
		metaTypes: []RPGItemMetaType{MetaItemType_Physical_Weapon, MetaItemType_Weapon},
		BasicItemType: item.CreateBasicItemType(
			1007,
			"Mace",
			"Smashey, smashey",
			'i',
			" ─¤",
			1,
			item.EquippedSlotDominantHand,
			tcell.StyleDefault.Foreground(tcell.ColorSilver),
		),
	}
}

func ItemTypeQuarterstaff() RPGItemType {

	return &BasicRPGItemType{
		damageRollFunc: func(victim, attacker RPGEntity) (damage int, dmgType DamageType) {
			return RollD6(1), DamageType_Physical_Bludgeoning
		},
		metaTypes: []RPGItemMetaType{MetaItemType_Physical_Weapon, MetaItemType_Weapon},
		BasicItemType: item.CreateBasicItemType(
			1008,
			"Quarterstaff",
			"Whacky, whacky",
			'|',
			"───",
			1,
			item.EquippedSlotDominantHand,
			tcell.StyleDefault.Foreground(tcell.ColorSaddleBrown),
		),
	}
}

func ItemTypeSickle() RPGItemType {
	return &BasicRPGItemType{
		damageRollFunc: func(victim, attacker RPGEntity) (damage int, dmgType DamageType) {
			return RollD6(1), DamageType_Physical_Slashing
		},
		metaTypes: []RPGItemMetaType{MetaItemType_Physical_Weapon, MetaItemType_Weapon},
		BasicItemType: item.CreateBasicItemType(
			1009,
			"Sickle",
			"Slicey, slicey?",
			'?',
			" ─U",
			1,
			item.EquippedSlotDominantHand,
			tcell.StyleDefault.Foreground(tcell.ColorSilver),
		),
	}
}

func ItemTypeSpear() RPGItemType {
	return &BasicRPGItemType{
		damageRollFunc: func(victim, attacker RPGEntity) (damage int, dmgType DamageType) {
			return RollD8(1), DamageType_Physical_Piercing
		},
		metaTypes: []RPGItemMetaType{MetaItemType_Physical_Weapon, MetaItemType_Weapon},
		BasicItemType: item.CreateBasicItemType(
			1010,
			"Spear",
			"Pokey, pokey",
			'Î',
			"──>",
			1,
			item.EquippedSlotDominantHand,
			tcell.StyleDefault.Foreground(tcell.ColorSilver),
		),
	}
}

type BasicRPGItem struct {
	modifiers []StatModifier

	item.BasicItem
}

func (i *BasicRPGItem) Modifiers() []StatModifier {
	return i.modifiers
}

func CreateRPGItem(name string, style tcell.Style, itemType RPGItemType, modifiers []StatModifier) RPGItem {
	return &BasicRPGItem{
		modifiers: modifiers,
		BasicItem: item.CreateBasicItemWithName(
			name,
			style,
			itemType,
			1,
		),
	}
}
