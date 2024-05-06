package model

import "mvvasilev/last_light/engine"

type Inventory interface {
	Items() []*Item
	Shape() engine.Size
	Push(item Item) bool
	Drop(x, y int) *Item
	ItemAt(x, y int) *Item
}

type BasicInventory struct {
	contents []*Item
	shape    engine.Size
}

func CreateInventory(shape engine.Size) *BasicInventory {
	inv := new(BasicInventory)

	inv.contents = make([]*Item, 0, shape.Height()*shape.Width())
	inv.shape = shape

	return inv
}

func (i *BasicInventory) Items() (items []*Item) {
	return i.contents
}

func (i *BasicInventory) Shape() engine.Size {
	return i.shape
}

func (i *BasicInventory) Push(item Item) (success bool) {
	if len(i.contents) == i.shape.Area() {
		return false
	}

	itemType := item.Type()

	// Try to first find a matching item with capacity
	for index, existingItem := range i.contents {
		if existingItem != nil && existingItem.itemType == itemType {
			if existingItem.Quantity()+1 > existingItem.Type().MaxStack() {
				continue
			}

			it := CreateItem(itemType, existingItem.Quantity()+1)
			i.contents[index] = &it

			return true
		}
	}

	// Next, try to find an intermediate empty slot to fit this item into
	for index, existingItem := range i.contents {
		if existingItem == nil {
			i.contents[index] = &item

			return true
		}
	}

	// Finally, just append the new item at the end
	i.contents = append(i.contents, &item)

	return true
}

func (i *BasicInventory) Drop(x, y int) *Item {
	index := y*i.shape.Width() + x

	if index > len(i.contents)-1 {
		return nil
	}

	item := i.contents[index]

	i.contents[index] = nil

	return item
}

func (i *BasicInventory) ItemAt(x, y int) (item *Item) {
	index := y*i.shape.Width() + x

	if index > len(i.contents)-1 {
		return nil
	}

	return i.contents[index]
}
