package item

import (
	"github.com/gdamore/tcell/v2"
)

type ItemType interface {
	Name() string
	Description() string
	TileIcon() rune
	Icon() string
	Style() tcell.Style
	MaxStack() int
	EquippableSlot() EquippedSlot
}

type BasicItemType struct {
	name           string
	description    string
	tileIcon       rune
	itemIcon       string
	maxStack       int
	equippableSlot EquippedSlot

	style tcell.Style
}

func CreateBasicItemType(
	name, description string,
	tileIcon rune,
	icon string,
	maxStack int,
	equippableSlot EquippedSlot,
	style tcell.Style,
) *BasicItemType {
	return &BasicItemType{
		name:           name,
		description:    description,
		tileIcon:       tileIcon,
		itemIcon:       icon,
		style:          style,
		maxStack:       maxStack,
		equippableSlot: equippableSlot,
	}
}

func (it *BasicItemType) Name() string {
	return it.name
}

func (it *BasicItemType) Description() string {
	return it.description
}

func (it *BasicItemType) TileIcon() rune {
	return it.tileIcon
}

func (it *BasicItemType) Icon() string {
	return it.itemIcon
}

func (it *BasicItemType) Style() tcell.Style {
	return it.style
}

func (it *BasicItemType) MaxStack() int {
	return it.maxStack
}

func (it *BasicItemType) EquippableSlot() EquippedSlot {
	return it.equippableSlot
}

func ItemTypeFish() ItemType {
	return &BasicItemType{
		name:           "Fish",
		description:    "What's a fish doing down here?",
		tileIcon:       '>',
		itemIcon:       "»o>",
		style:          tcell.StyleDefault.Foreground(tcell.ColorDarkCyan),
		equippableSlot: EquippedSlotNone,
		maxStack:       16,
	}
}

func ItemTypeGold() ItemType {
	return &BasicItemType{
		name:           "Gold",
		description:    "Not all those who wander are lost",
		tileIcon:       '¤',
		itemIcon:       " ¤ ",
		equippableSlot: EquippedSlotNone,
		style:          tcell.StyleDefault.Foreground(tcell.ColorGoldenrod),
		maxStack:       255,
	}
}

func ItemTypeArrow() ItemType {
	return &BasicItemType{
		name:           "Arrow",
		description:    "Ammunition for a bow",
		tileIcon:       '-',
		itemIcon:       "»->",
		equippableSlot: EquippedSlotNone,
		style:          tcell.StyleDefault.Foreground(tcell.ColorGoldenrod),
		maxStack:       32,
	}
}

func ItemTypeKey() ItemType {
	return &BasicItemType{
		name:           "Key",
		description:    "Indispensable for unlocking things",
		tileIcon:       '¬',
		itemIcon:       " o╖",
		equippableSlot: EquippedSlotNone,
		style:          tcell.StyleDefault.Foreground(tcell.ColorDarkGoldenrod),
		maxStack:       1,
	}
}
