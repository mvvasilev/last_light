package ui

import (
	"mvvasilev/last_light/util"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	"github.com/google/uuid"
)

type UIContainerLayout int

const (
	// These change the provided ui positions
	UpperLeft UIContainerLayout = iota
	MiddleLeft
	LowerLeft
	UpperRight
	MiddleRight
	LowerRight
	UpperCenter
	MiddleCenter
	LowerCenter
	// This uses the positions as provided in the ui elements
	Manual
)

type UIContainer struct {
	id       uuid.UUID
	layout   UIContainerLayout
	position util.Position
	size     util.Size
	elements []UIElement
}

func CreateUIContainer(x, y uint16, width, height uint16, layout UIContainerLayout) *UIContainer {
	container := new(UIContainer)

	container.id = uuid.New()
	container.layout = layout
	container.position = util.PositionAt(x, y)
	container.size = util.SizeOf(width, height)
	container.elements = make([]UIElement, 0)

	return container
}

func (uic *UIContainer) Push(element UIElement) {
	uic.elements = append(uic.elements, element)
}

func (uic *UIContainer) Clear() {
	uic.elements = make([]UIElement, 0)
}

func (uic *UIContainer) UniqueId() uuid.UUID {
	return uic.id
}

func (uic *UIContainer) MoveTo(x, y uint16) {
	uic.position = util.PositionAt(x, y)
}

func (uic *UIContainer) Position() util.Position {
	return uic.position
}

func (uic *UIContainer) Size() util.Size {
	return uic.size
}

func (uic *UIContainer) Draw(v views.View) {
	for _, e := range uic.elements {
		e.Draw(v)
	}
}

func (uic *UIContainer) Input(ev *tcell.EventKey) {
	for _, e := range uic.elements {
		e.Input(ev)
	}
}
