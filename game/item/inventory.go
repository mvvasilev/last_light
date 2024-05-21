package item

import (
	"mvvasilev/last_light/engine"
)

type Inventory interface {
	Items() []Item
	Shape() engine.Size
	Push(item Item) bool
	Drop(x, y int) Item
	ItemAt(x, y int) Item
}

type BasicInventory struct {
	contents []Item
	shape    engine.Size
}

func CreateInventory(shape engine.Size) *BasicInventory {
	inv := new(BasicInventory)

	inv.contents = make([]Item, 0, shape.Height()*shape.Width())
	inv.shape = shape

	return inv
}

func (i *BasicInventory) Items() (items []Item) {
	return i.contents
}

func (i *BasicInventory) Shape() engine.Size {
	return i.shape
}

func (inv *BasicInventory) Push(i Item) (success bool) {
	if len(inv.contents) == inv.shape.Area() {
		return false
	}

	itemType := i.Type()

	// Try to first find a matching item with capacity
	for index, existingItem := range inv.contents {
		if existingItem != nil && existingItem.Type() == itemType {
			if existingItem.Quantity()+1 > existingItem.Type().MaxStack() {
				continue
			}

			it := CreateBasicItem(itemType, existingItem.Quantity()+1)
			inv.contents[index] = &it

			return true
		}
	}

	// Next, try to find an intermediate empty slot to fit this item into
	for index, existingItem := range inv.contents {
		if existingItem == nil {
			inv.contents[index] = i

			return true
		}
	}

	// Finally, just append the new item at the end
	inv.contents = append(inv.contents, i)

	return true
}

func (i *BasicInventory) Drop(x, y int) Item {
	index := y*i.shape.Width() + x

	if index > len(i.contents)-1 {
		return nil
	}

	item := i.contents[index]

	i.contents[index] = nil

	return item
}

func (i *BasicInventory) ItemAt(x, y int) (item Item) {
	index := y*i.shape.Width() + x

	if index > len(i.contents)-1 {
		return nil
	}

	return i.contents[index]
}
