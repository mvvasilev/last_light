package model

import (
	"math/rand"

	"github.com/gdamore/tcell/v2"
)

type ItemType struct {
	name        string
	description string
	tileIcon    rune
	itemIcon    string
	maxStack    int

	style tcell.Style
}

func (it *ItemType) Name() string {
	return it.name
}

func (it *ItemType) Description() string {
	return it.description
}

func (it *ItemType) TileIcon() rune {
	return it.tileIcon
}

func (it *ItemType) Icon() string {
	return it.itemIcon
}

func (it *ItemType) Style() tcell.Style {
	return it.style
}

func (it *ItemType) MaxStack() int {
	return it.maxStack
}

func ItemTypeFish() *ItemType {
	return &ItemType{
		name:        "Fish",
		description: "What's a fish doing down here?",
		tileIcon:    '>',
		itemIcon:    "»o>",
		style:       tcell.StyleDefault.Foreground(tcell.ColorDarkCyan),
		maxStack:    16,
	}
}

func ItemTypeGold() *ItemType {
	return &ItemType{
		name:        "Gold",
		description: "Not all those who wander are lost",
		tileIcon:    '¤',
		itemIcon:    " ¤ ",
		style:       tcell.StyleDefault.Foreground(tcell.ColorGoldenrod),
		maxStack:    255,
	}
}

func ItemTypeArrow() *ItemType {
	return &ItemType{
		name:        "Arrow",
		description: "Ammunition for a bow",
		tileIcon:    '-',
		itemIcon:    "»->",
		style:       tcell.StyleDefault.Foreground(tcell.ColorGoldenrod),
		maxStack:    32,
	}
}

func ItemTypeBow() *ItemType {
	return &ItemType{
		name:        "Bow",
		description: "To shoot arrows with",
		tileIcon:    ')',
		itemIcon:    " |)",
		style:       tcell.StyleDefault.Foreground(tcell.ColorBrown),
		maxStack:    1,
	}
}

func ItemTypeLongsword() *ItemType {
	return &ItemType{
		name:        "Longsword",
		description: "You know nothing.",
		tileIcon:    '/',
		itemIcon:    "╪══",
		style:       tcell.StyleDefault.Foreground(tcell.ColorSilver),
		maxStack:    1,
	}
}

func ItemTypeKey() *ItemType {
	return &ItemType{
		name:        "Key",
		description: "Indispensable for unlocking things",
		tileIcon:    '¬',
		itemIcon:    " o╖",
		style:       tcell.StyleDefault.Foreground(tcell.ColorDarkGoldenrod),
		maxStack:    1,
	}
}

type ItemTypeGenTable struct {
}

func GenerateItemType(genTable map[float32]*ItemType) *ItemType {
	num := rand.Float32()

	for k, v := range genTable {
		if num > k {
			return v
		}
	}

	return nil
}

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
