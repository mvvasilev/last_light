package model

import (
	"github.com/gdamore/tcell/v2"
)

type Item struct {
	name        string
	description string
	itemType    *ItemType
	quantity    int
}

func EmptyItem() Item {
	return Item{
		itemType: &ItemType{
			name:        "",
			description: "",
			tileIcon:    ' ',
			itemIcon:    "   ",
			style:       tcell.StyleDefault,
			maxStack:    0,
		},
	}
}

func CreateItem(itemType *ItemType, quantity int) Item {
	return Item{
		itemType: itemType,
		quantity: quantity,
	}
}

func (i Item) WithName(name string) Item {
	i.name = name

	return i
}

func (i Item) Name() string {
	if i.name == "" {
		return i.itemType.name
	}

	return i.name
}

func (i Item) Description() string {
	if i.description == "" {
		return i.itemType.description
	}

	return i.description
}

func (i Item) WithDescription(description string) Item {
	i.description = description

	return i
}

func (i Item) Type() *ItemType {
	return i.itemType
}

func (i Item) Quantity() int {
	return i.quantity
}
