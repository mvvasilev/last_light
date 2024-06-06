package ui

import (
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/systems"
	"unicode/utf8"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	"github.com/google/uuid"
)

type UILabel struct {
	id   uuid.UUID
	text *engine.Text
}

func CreateUILabel(x, y int, width, height int, content string, style tcell.Style) *UILabel {
	label := new(UILabel)

	label.id = uuid.New()
	label.text = engine.CreateText(x, y, width, height, content, style)

	return label
}

func CreateSingleLineUILabel(x, y int, content string, style tcell.Style) *UILabel {
	label := new(UILabel)

	label.id = uuid.New()
	label.text = engine.CreateText(x, y, int(utf8.RuneCountInString(content)), 1, content, style)

	return label
}

func (t *UILabel) UniqueId() uuid.UUID {
	return t.id
}

func (t *UILabel) MoveTo(x int, y int) {
	t.text = engine.CreateText(x, y, int(t.text.Size().Width()), int(t.Size().Height()), t.text.Content(), t.text.Style())
}

func (t *UILabel) Position() engine.Position {
	return t.text.Position()
}

func (t *UILabel) SetContent(content string) {
	t.text = engine.CreateText(t.text.Position().X(), t.text.Position().Y(), int(t.text.Size().Width()), int(t.Size().Height()), content, t.text.Style())
}

func (t *UILabel) Size() engine.Size {
	return t.text.Size()
}

func (t *UILabel) Draw(v views.View) {
	t.text.Draw(v)
}

func (t *UILabel) Input(inputAction systems.InputAction) {}
