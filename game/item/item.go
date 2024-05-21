package item

import (
	"github.com/gdamore/tcell/v2"
)

type Item interface {
	Name() (string, tcell.Style)
	Description() string
	Type() ItemType
	Quantity() int
}

type BasicItem struct {
	name        string
	nameStyle   tcell.Style
	description string
	itemType    ItemType
	quantity    int
}

func EmptyItem() BasicItem {
	return BasicItem{
		nameStyle: tcell.StyleDefault,
		itemType: &BasicItemType{
			name:        "",
			description: "",
			tileIcon:    ' ',
			itemIcon:    "   ",
			style:       tcell.StyleDefault,
			maxStack:    0,
		},
	}
}

func CreateBasicItem(itemType ItemType, quantity int) BasicItem {
	return BasicItem{
		itemType: itemType,
		quantity: quantity,
	}
}

func CreateBasicItemWithName(name string, style tcell.Style, itemType ItemType, quantity int) BasicItem {
	return BasicItem{
		name:      name,
		nameStyle: style,
		itemType:  itemType,
		quantity:  quantity,
	}
}

func (i BasicItem) WithName(name string, style tcell.Style) BasicItem {
	i.name = name
	i.nameStyle = style

	return i
}

func (i BasicItem) Name() (string, tcell.Style) {
	if i.name == "" {
		return i.itemType.Name(), i.nameStyle
	}

	return i.name, i.nameStyle
}

func (i BasicItem) Description() string {
	if i.description == "" {
		return i.itemType.Description()
	}

	return i.description
}

func (i BasicItem) WithDescription(description string) BasicItem {
	i.description = description

	return i
}

func (i BasicItem) Type() ItemType {
	return i.itemType
}

func (i BasicItem) Quantity() int {
	return i.quantity
}
