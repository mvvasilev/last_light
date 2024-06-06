package model

import (
	"mvvasilev/last_light/engine"

	"github.com/gdamore/tcell/v2"
)

type ItemMetaType int

const (
	MetaItemType_Physical_Weapon ItemMetaType = iota
	MetaItemType_Magic_Weapon
	MetaItemType_Weapon
	MetaItemType_Physical_Armour
	MetaItemType_Magic_Armour
	MetaItemType_Armour
	MetaItemType_Consumable
	MetaItemType_Potion
)

type Item_V2 interface {
	TileIcon() rune
	Icon() string
	Style() tcell.Style
	Type() ItemType

	Quantifiable() *Item_QuantifiableComponent
	Usable() *Item_UsableComponent
	Equippable() *Item_EquippableComponent
	Named() *Item_NamedComponent
	Described() *Item_DescribedComponent
	Damaging() *Item_DamagingComponent
	StatModifier() *Item_StatModifierComponent
	MetaTypes() *Item_MetaTypesComponent
}

type Item_QuantifiableComponent struct {
	MaxQuantity     int
	CurrentQuantity int
}

type Item_UsableComponent struct {
	IsUsableBy func(Entity_V2) bool
	Use        func(*engine.GameEventLog, *Dungeon, Entity_V2)
}

type Item_EquippableComponent struct {
	Slot EquippedSlot
}

type Item_NamedComponent struct {
	Name  string
	Style tcell.Style
}

type Item_DescribedComponent struct {
	Description string
	Style       tcell.Style
}

type Item_DamagingComponent struct {
	DamageRoll func() (damage int, dmgType DamageType)
}

type Item_StatModifierComponent struct {
	StatModifiers []StatModifier
}

type Item_MetaTypesComponent struct {
	Types []ItemMetaType
}

type BaseItem_V2 struct {
	tileIcon rune
	icon     string
	style    tcell.Style
	itemType ItemType

	quantifiable *Item_QuantifiableComponent
	usable       *Item_UsableComponent
	equippable   *Item_EquippableComponent
	named        *Item_NamedComponent
	described    *Item_DescribedComponent
	damaging     *Item_DamagingComponent
	statModifier *Item_StatModifierComponent
	metaTypes    *Item_MetaTypesComponent
}

func (i *BaseItem_V2) TileIcon() rune {
	return i.tileIcon
}

func (i *BaseItem_V2) Icon() string {
	return i.icon
}

func (i *BaseItem_V2) Style() tcell.Style {
	return i.style
}

func (i *BaseItem_V2) Type() ItemType {
	return i.itemType
}

func (i *BaseItem_V2) Quantifiable() *Item_QuantifiableComponent {
	return i.quantifiable
}

func (i *BaseItem_V2) Usable() *Item_UsableComponent {
	return i.usable
}

func (i *BaseItem_V2) Equippable() *Item_EquippableComponent {
	return i.equippable
}

func (i *BaseItem_V2) Named() *Item_NamedComponent {
	return i.named
}

func (i *BaseItem_V2) Described() *Item_DescribedComponent {
	return i.described
}

func (i *BaseItem_V2) Damaging() *Item_DamagingComponent {
	return i.damaging
}

func (i *BaseItem_V2) StatModifier() *Item_StatModifierComponent {
	return i.statModifier
}

func (i *BaseItem_V2) MetaTypes() *Item_MetaTypesComponent {
	return i.metaTypes
}

func createBaseItem(itemType ItemType, tileIcon rune, icon string, style tcell.Style, components ...func(*BaseItem_V2)) *BaseItem_V2 {
	i := &BaseItem_V2{
		itemType: itemType,
		tileIcon: tileIcon,
		icon:     icon,
		style:    style,
	}

	for _, comp := range components {
		comp(i)
	}

	return i
}

func item_WithQuantity(quantity, maxQuantity int) func(*BaseItem_V2) {
	return func(bi *BaseItem_V2) {
		bi.quantifiable = &Item_QuantifiableComponent{
			CurrentQuantity: quantity,
			MaxQuantity:     maxQuantity,
		}
	}
}

func item_WithUsable(usabilityCheck func(Entity_V2) bool, useFunc func(*engine.GameEventLog, *Dungeon, Entity_V2)) func(*BaseItem_V2) {
	return func(bi *BaseItem_V2) {
		bi.usable = &Item_UsableComponent{
			IsUsableBy: usabilityCheck,
			Use:        useFunc,
		}
	}
}

func item_WithEquippable(slot EquippedSlot) func(*BaseItem_V2) {
	return func(bi *BaseItem_V2) {
		bi.equippable = &Item_EquippableComponent{
			Slot: slot,
		}
	}
}

func item_WithDamaging(damageFunc func() (damage int, dmgType DamageType)) func(*BaseItem_V2) {
	return func(bi *BaseItem_V2) {
		bi.damaging = &Item_DamagingComponent{
			DamageRoll: damageFunc,
		}
	}
}

func item_WithName(name string, style tcell.Style) func(*BaseItem_V2) {
	return func(bi *BaseItem_V2) {
		bi.named = &Item_NamedComponent{
			Name:  name,
			Style: style,
		}
	}
}

func item_WithDescription(description string, style tcell.Style) func(*BaseItem_V2) {
	return func(bi *BaseItem_V2) {
		bi.described = &Item_DescribedComponent{
			Description: description,
			Style:       style,
		}
	}
}

func item_WithStatModifiers(statModifiers []StatModifier) func(*BaseItem_V2) {
	return func(bi *BaseItem_V2) {
		bi.statModifier = &Item_StatModifierComponent{
			StatModifiers: statModifiers,
		}
	}
}

func item_WithMetaTypes(metaTypes []ItemMetaType) func(*BaseItem_V2) {
	return func(bi *BaseItem_V2) {
		bi.metaTypes = &Item_MetaTypesComponent{
			Types: metaTypes,
		}
	}
}

// type Item interface {
// 	Name() (string, tcell.Style)
// 	Description() string
// 	Type() ItemType
// 	Quantity() int
// 	SetQuantity(quant int) Item
// }

// type BasicItem struct {
// 	name        string
// 	nameStyle   tcell.Style
// 	description string
// 	itemType    ItemType
// 	quantity    int
// }

// func EmptyItem() BasicItem {
// 	return BasicItem{
// 		nameStyle: tcell.StyleDefault,
// 		itemType: &BasicItemType{
// 			name:        "",
// 			description: "",
// 			tileIcon:    ' ',
// 			itemIcon:    "   ",
// 			style:       tcell.StyleDefault,
// 			maxStack:    0,
// 		},
// 	}
// }

// func CreateBasicItem(itemType ItemType, quantity int) BasicItem {
// 	return BasicItem{
// 		itemType: itemType,
// 		quantity: quantity,
// 	}
// }

// func CreateBasicItemWithName(name string, style tcell.Style, itemType ItemType, quantity int) BasicItem {
// 	return BasicItem{
// 		name:      name,
// 		nameStyle: style,
// 		itemType:  itemType,
// 		quantity:  quantity,
// 	}
// }

// func (i BasicItem) WithName(name string, style tcell.Style) BasicItem {
// 	i.name = name
// 	i.nameStyle = style

// 	return i
// }

// func (i BasicItem) Name() (string, tcell.Style) {
// 	if i.name == "" {
// 		return i.itemType.Name(), i.nameStyle
// 	}

// 	return i.name, i.nameStyle
// }

// func (i BasicItem) Description() string {
// 	if i.description == "" {
// 		return i.itemType.Description()
// 	}

// 	return i.description
// }

// func (i BasicItem) WithDescription(description string) BasicItem {
// 	i.description = description

// 	return i
// }

// func (i BasicItem) Type() ItemType {
// 	return i.itemType
// }

// func (i BasicItem) Quantity() int {
// 	return i.quantity
// }

// func (i BasicItem) SetQuantity(amount int) Item {
// 	i.quantity = i.quantity - amount

// 	return i
// }
