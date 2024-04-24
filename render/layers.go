package render

import (
	"slices"

	"github.com/gdamore/tcell/v2/views"
	"github.com/google/uuid"
)

type layer struct {
	zIndex   uint8
	contents []Drawable
}

func makeLayer(zIndex uint8) *layer {
	l := new(layer)

	l.zIndex = zIndex
	l.contents = make([]Drawable, 0, 1)

	return l
}

func (l *layer) push(drawable Drawable) {
	l.contents = append(l.contents, drawable)
}

func (l *layer) remove(uuid uuid.UUID) {
	l.contents = slices.DeleteFunc(l.contents, func(d Drawable) bool {
		return d.UniqueId() == uuid
	})
}

func (l *layer) draw(s views.View) {
	for _, d := range l.contents {
		d.Draw(s)
	}
}

type UnorderedDrawContainer struct {
	id       uuid.UUID
	contents []Drawable
}

func CreateUnorderedDrawContainer(contents []Drawable) UnorderedDrawContainer {
	return UnorderedDrawContainer{
		id:       uuid.New(),
		contents: contents,
	}
}

type layeredDrawContainer struct {
	id     uuid.UUID
	layers []*layer
}

func CreateLayeredDrawContainer() *layeredDrawContainer {
	container := new(layeredDrawContainer)

	container.layers = make([]*layer, 0, 32)

	return container
}

func (ldc *layeredDrawContainer) Insert(zLevel uint8, drawable Drawable) {
	// if no layers exist, just insert a new one
	if len(ldc.layers) == 0 {
		l := makeLayer(zLevel)
		l.push(drawable)
		ldc.layers = append(ldc.layers, l)

		return
	}

	// find a layer with this z-index
	i := slices.IndexFunc(ldc.layers, func(l *layer) bool {
		return l.zIndex == zLevel
	})

	// z index already exists
	if i > 0 {
		l := ldc.layers[i]
		l.push(drawable)

		return
	}

	// no such layer exists, create it

	l := makeLayer(zLevel)
	l.push(drawable)

	ldc.layers = append(ldc.layers, l)

	// order layers ascending
	slices.SortFunc(ldc.layers, func(l1 *layer, l2 *layer) int {
		return int(l1.zIndex) - int(l2.zIndex)
	})
}

func (ldc *layeredDrawContainer) Remove(id uuid.UUID) {
	for _, l := range ldc.layers {
		l.remove(id)
	}
}

func (ldc *layeredDrawContainer) Clear() {
	ldc.layers = make([]*layer, 0, 32)
}

func (ldc *layeredDrawContainer) UniqueId() uuid.UUID {
	return ldc.id
}

func (ldc *layeredDrawContainer) Draw(s views.View) {
	for _, d := range ldc.layers {
		d.draw(s)
	}
}
