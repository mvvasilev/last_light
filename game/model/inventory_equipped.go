package model

import (
	"mvvasilev/last_light/engine"
)

type EquippedSlot int

const (
	EquippedSlotNone EquippedSlot = 0

	EquippedSlotOffhand      EquippedSlot = 1
	EquippedSlotDominantHand EquippedSlot = 2
	EquippedSlotHead         EquippedSlot = 3
	EquippedSlotChestplate   EquippedSlot = 4
	EquippedSlotLeggings     EquippedSlot = 5
	EquippedSlotShoes        EquippedSlot = 6
)

type EquippedInventory struct {
	offHand      Item
	dominantHand Item

	head       Item
	chestplate Item
	leggings   Item
	shoes      Item

	*BasicInventory
}

func CreateEquippedInventory() *EquippedInventory {
	return &EquippedInventory{
		BasicInventory: CreateInventory(engine.SizeOf(8, 4)),
	}
}

func BuildInventory(manips ...func(*EquippedInventory)) *EquippedInventory {
	ei := CreateEquippedInventory()

	for _, m := range manips {
		m(ei)
	}

	return ei
}

func Inv_WithOffHand(item Item) func(*EquippedInventory) {
	return func(ei *EquippedInventory) {
		ei.offHand = item
	}
}

func Inv_WithDominantHand(item Item) func(*EquippedInventory) {
	return func(ei *EquippedInventory) {
		ei.dominantHand = item
	}
}

func Inv_WithHead(item Item) func(*EquippedInventory) {
	return func(ei *EquippedInventory) {
		ei.head = item
	}
}

func Inv_WithChest(item Item) func(*EquippedInventory) {
	return func(ei *EquippedInventory) {
		ei.chestplate = item
	}
}

func Inv_WithLegs(item Item) func(*EquippedInventory) {
	return func(ei *EquippedInventory) {
		ei.leggings = item
	}
}

func Inv_WithShoes(item Item) func(*EquippedInventory) {
	return func(ei *EquippedInventory) {
		ei.shoes = item
	}
}

func Inv_WithContents(items ...Item) func(*EquippedInventory) {
	return func(ei *EquippedInventory) {
		for _, i := range items {
			ei.Push(i)
		}
	}
}

func (ei *EquippedInventory) AtSlot(slot EquippedSlot) Item {
	switch slot {
	case EquippedSlotOffhand:
		return ei.offHand
	case EquippedSlotDominantHand:
		return ei.dominantHand
	case EquippedSlotHead:
		return ei.head
	case EquippedSlotChestplate:
		return ei.chestplate
	case EquippedSlotLeggings:
		return ei.leggings
	case EquippedSlotShoes:
		return ei.shoes
	default:
		return nil
	}
}

func (ei *EquippedInventory) Equip(item Item, slot EquippedSlot) Item {
	switch slot {
	case EquippedSlotOffhand:
		ei.offHand = item
	case EquippedSlotDominantHand:
		ei.dominantHand = item
	case EquippedSlotHead:
		ei.head = item
	case EquippedSlotChestplate:
		ei.chestplate = item
	case EquippedSlotLeggings:
		ei.leggings = item
	case EquippedSlotShoes:
		ei.shoes = item
	}

	return item
}
