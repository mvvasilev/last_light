package model

import "mvvasilev/last_light/util"

type EquippedSlot int

const (
	EquippedSlotOffhand EquippedSlot = iota
	EquippedSlotDominantHand
	EquippedSlotHead
	EquippedSlotChestplate
	EquippedSlotLeggings
	EquippedSlotShoes
)

type EquippedInventory struct {
	offHand      *Item
	dominantHand *Item

	head       *Item
	chestplate *Item
	leggings   *Item
	shoes      *Item

	*BasicInventory
}

func CreatePlayerInventory() *EquippedInventory {
	return &EquippedInventory{
		BasicInventory: CreateInventory(util.SizeOf(8, 4)),
	}
}

func (ei *EquippedInventory) AtSlot(slot EquippedSlot) *Item {
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

func (ei *EquippedInventory) Equip(item Item, slot EquippedSlot) *Item {
	ref := &item

	switch slot {
	case EquippedSlotOffhand:
		ei.offHand = ref
	case EquippedSlotDominantHand:
		ei.dominantHand = ref
	case EquippedSlotHead:
		ei.head = ref
	case EquippedSlotChestplate:
		ei.chestplate = ref
	case EquippedSlotLeggings:
		ei.leggings = ref
	case EquippedSlotShoes:
		ei.shoes = ref
	}

	return ref
}
