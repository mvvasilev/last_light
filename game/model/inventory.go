package model

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
		if existingItem == nil {
			continue
		}

		itemsAreSame := existingItem.Type() == itemType
		bothItemsAreQuantifiable := existingItem.Quantifiable() != nil && i.Quantifiable() != nil

		if itemsAreSame && bothItemsAreQuantifiable {
			existingCurrent := existingItem.Quantifiable().CurrentQuantity
			incomingCurrent := i.Quantifiable().CurrentQuantity
			existingMax := existingItem.Quantifiable().MaxQuantity

			// Cannot add even a single more item to this stack, skip it
			if existingItem.Quantifiable().CurrentQuantity+1 > existingItem.Quantifiable().MaxQuantity {
				continue
			}

			total := existingCurrent + incomingCurrent
			leftOver := engine.AbsInt(existingMax - total)

			// Existing item is filled, and remained is turned into new stack
			if leftOver > 0 {
				// If we have don't have enough free slots, just say we can't push it
				if index+1 >= inv.shape.Area() {
					return false
				}

				existingItem.Quantifiable().CurrentQuantity = existingMax

				i.Quantifiable().CurrentQuantity = leftOver

				inv.contents[index+1] = i

				return true
			}

			// Otherwise, just set the existing item quantity to the total
			existingItem.Quantifiable().CurrentQuantity = total

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

func (i *BasicInventory) ReduceQuantityAt(x, y int, amount int) {
	it := i.ItemAt(x, y)

	if it == nil {
		return
	}

	quantityData := it.Quantifiable()

	if quantityData != nil {
		if quantityData.CurrentQuantity-amount <= 0 {
			i.Drop(x, y)
		} else {
			quantityData.CurrentQuantity = quantityData.CurrentQuantity - amount
		}
	} else {
		i.Drop(x, y)
	}
}

func (i *BasicInventory) ItemAt(x, y int) (item Item) {
	index := y*i.shape.Width() + x

	if index > len(i.contents)-1 {
		return nil
	}

	return i.contents[index]
}

func (i *BasicInventory) Find(filter func(i Item) bool) Item {
	for _, c := range i.contents {
		if filter(c) {
			return c
		}
	}

	return nil
}
