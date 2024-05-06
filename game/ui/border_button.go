package ui

import (
	"mvvasilev/last_light/engine"
	engine1 "mvvasilev/last_light/engine"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	"github.com/google/uuid"
)

type UIBorderedButton struct {
	id uuid.UUID

	text   engine.Text
	border engine.Rectangle

	isSelected bool

	unselectedStyle tcell.Style
	selectedStyle   tcell.Style
}

func (b *UIBorderedButton) IsSelected() bool {
	return b.isSelected
}

func (b *UIBorderedButton) Select() {
	b.isSelected = true
}

func (b *UIBorderedButton) Deselect() {
	b.isSelected = false
}

func (b *UIBorderedButton) SetSelected(selected bool) {
	b.isSelected = selected
}

func (b *UIBorderedButton) UniqueId() uuid.UUID {
	return b.id
}

func (b *UIBorderedButton) MoveTo(x int, y int) {
	panic("not implemented") // TODO: Implement
}

func (b *UIBorderedButton) Position() engine1.Position {
	panic("not implemented") // TODO: Implement
}

func (b *UIBorderedButton) Size() engine1.Size {
	panic("not implemented") // TODO: Implement
}

func (b *UIBorderedButton) Draw(v views.View) {
	panic("not implemented") // TODO: Implement
}

func (b *UIBorderedButton) Input(e *tcell.EventKey) {
	panic("not implemented") // TODO: Implement
}
