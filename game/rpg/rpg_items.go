package rpg

import (
	"mvvasilev/last_light/game/item"

	"github.com/gdamore/tcell/v2"
)

type RPGItemType interface {
	RollDamage() func(victim, attacker RPGEntity) (damage int, dmgType DamageType)

	item.ItemType
}

type RPGItem interface {
	Modifiers() []StatModifier

	item.Item
}

type BasicRPGItemType struct {
	damageRollFunc func(victim, attacker RPGEntity) (damage int, dmgType DamageType)

	*item.BasicItemType
}

func (it *BasicRPGItemType) RollDamage() func(victim, attacker RPGEntity) (damage int, dmgType DamageType) {
	return it.damageRollFunc
}

func ItemTypeBow() RPGItemType {
	return &BasicRPGItemType{
		damageRollFunc: func(victim, attacker RPGEntity) (damage int, dmgType DamageType) {
			// TODO: Ranged
			return RollD8(1), DamageType_Physical_Piercing
		},
		BasicItemType: item.CreateBasicItemType(
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
		BasicItemType: item.CreateBasicItemType(
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
		BasicItemType: item.CreateBasicItemType(
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
		BasicItemType: item.CreateBasicItemType(
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
		BasicItemType: item.CreateBasicItemType(
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
		BasicItemType: item.CreateBasicItemType(
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
		BasicItemType: item.CreateBasicItemType(
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
		BasicItemType: item.CreateBasicItemType(
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
		BasicItemType: item.CreateBasicItemType(
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
		BasicItemType: item.CreateBasicItemType(
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
		BasicItemType: item.CreateBasicItemType(
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
