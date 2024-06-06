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

type Item interface {
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
	IsUsableBy func(Entity) bool
	Use        func(*engine.GameEventLog, *Dungeon, Entity)
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

type BaseItem struct {
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

func (i *BaseItem) TileIcon() rune {
	return i.tileIcon
}

func (i *BaseItem) Icon() string {
	return i.icon
}

func (i *BaseItem) Style() tcell.Style {
	return i.style
}

func (i *BaseItem) Type() ItemType {
	return i.itemType
}

func (i *BaseItem) Quantifiable() *Item_QuantifiableComponent {
	return i.quantifiable
}

func (i *BaseItem) Usable() *Item_UsableComponent {
	return i.usable
}

func (i *BaseItem) Equippable() *Item_EquippableComponent {
	return i.equippable
}

func (i *BaseItem) Named() *Item_NamedComponent {
	return i.named
}

func (i *BaseItem) Described() *Item_DescribedComponent {
	return i.described
}

func (i *BaseItem) Damaging() *Item_DamagingComponent {
	return i.damaging
}

func (i *BaseItem) StatModifier() *Item_StatModifierComponent {
	return i.statModifier
}

func (i *BaseItem) MetaTypes() *Item_MetaTypesComponent {
	return i.metaTypes
}

func createBaseItem(itemType ItemType, tileIcon rune, icon string, style tcell.Style, components ...func(*BaseItem)) *BaseItem {
	i := &BaseItem{
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

func item_WithQuantity(quantity, maxQuantity int) func(*BaseItem) {
	return func(bi *BaseItem) {
		bi.quantifiable = &Item_QuantifiableComponent{
			CurrentQuantity: quantity,
			MaxQuantity:     maxQuantity,
		}
	}
}

func item_WithUsable(usabilityCheck func(Entity) bool, useFunc func(*engine.GameEventLog, *Dungeon, Entity)) func(*BaseItem) {
	return func(bi *BaseItem) {
		bi.usable = &Item_UsableComponent{
			IsUsableBy: usabilityCheck,
			Use:        useFunc,
		}
	}
}

func item_WithEquippable(slot EquippedSlot) func(*BaseItem) {
	return func(bi *BaseItem) {
		bi.equippable = &Item_EquippableComponent{
			Slot: slot,
		}
	}
}

func item_WithDamaging(damageFunc func() (damage int, dmgType DamageType)) func(*BaseItem) {
	return func(bi *BaseItem) {
		bi.damaging = &Item_DamagingComponent{
			DamageRoll: damageFunc,
		}
	}
}

func item_WithName(name string, style tcell.Style) func(*BaseItem) {
	return func(bi *BaseItem) {
		bi.named = &Item_NamedComponent{
			Name:  name,
			Style: style,
		}
	}
}

func item_WithDescription(description string, style tcell.Style) func(*BaseItem) {
	return func(bi *BaseItem) {
		bi.described = &Item_DescribedComponent{
			Description: description,
			Style:       style,
		}
	}
}

func item_WithStatModifiers(statModifiers []StatModifier) func(*BaseItem) {
	return func(bi *BaseItem) {
		bi.statModifier = &Item_StatModifierComponent{
			StatModifiers: statModifiers,
		}
	}
}

func item_WithMetaTypes(metaTypes []ItemMetaType) func(*BaseItem) {
	return func(bi *BaseItem) {
		bi.metaTypes = &Item_MetaTypesComponent{
			Types: metaTypes,
		}
	}
}
