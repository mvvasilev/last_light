package model

import "mvvasilev/last_light/util"

type Inventory struct {
	contents []*Item
	shape    util.Size
}

func CreateInventory(shape util.Size) *Inventory {
	inv := new(Inventory)

	inv.contents = make([]*Item, 0, shape.Height()*shape.Width())
	inv.shape = shape

	return inv
}

func (i *Inventory) Items() (items []*Item) {
	return i.contents
}

func (i *Inventory) Shape() util.Size {
	return i.shape
}

func (i *Inventory) Push(item Item) (success bool) {
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

func (i *Inventory) Drop(x, y int) {
	index := y*i.shape.Width() + x

	if index > len(i.contents)-1 {
		return
	}

	i.contents[index] = nil
}

func (i *Inventory) ItemAt(x, y int) (item *Item) {
	index := y*i.shape.Width() + x

	if index > len(i.contents)-1 {
		return nil
	}

	return i.contents[index]
}
