package render

import (
	"slices"

	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
	"github.com/tidwall/btree"
)

type layeredDrawContainer struct {
	id     uuid.UUID
	layers *btree.Map[uint8, []Drawable]
}

func CreateLayeredDrawContainer() *layeredDrawContainer {
	container := new(layeredDrawContainer)

	container.layers = btree.NewMap[uint8, []Drawable](2)

	return container
}

func (ldc *layeredDrawContainer) Insert(zLevel uint8, drawable Drawable) {
	arr, found := ldc.layers.Get(zLevel)

	if !found {
		arr = make([]Drawable, 1, 2)
	}

	arr = append(arr, drawable)

	ldc.layers.Set(zLevel, arr)
}

func (ldc *layeredDrawContainer) Remove(id uuid.UUID) {
	ldc.layers.ScanMut(func(key uint8, value []Drawable) bool {
		newSlices := slices.DeleteFunc(value, func(v Drawable) bool { return v.UniqueId() == id })

		ldc.layers.Set(key, newSlices)

		if len(newSlices) != len(value) {
			return false // the slice has been modified, we have found the drawable. Return false to stop iteration.
		} else {
			return true // we haven't found it yet, keep going
		}
	})
}

func (ldc *layeredDrawContainer) Clear() {
	ldc.layers = btree.NewMap[uint8, []Drawable](2)
}

func (ldc *layeredDrawContainer) UniqueId() uuid.UUID {
	return ldc.id
}

func (ldc *layeredDrawContainer) Draw(s tcell.Screen) {
	ldc.layers.Ascend(0, func(key uint8, value []Drawable) bool {
		for _, d := range value {
			d.Draw(s)
		}

		return true
	})
}
