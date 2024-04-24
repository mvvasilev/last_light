package ui

import (
	"mvvasilev/last_light/render"
	"mvvasilev/last_light/util"
	"unicode/utf8"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	"github.com/google/uuid"
)

type UILabel struct {
	id   uuid.UUID
	text *render.Text
}

func CreateUILabel(x, y uint16, width, height uint16, content string, style tcell.Style) *UILabel {
	label := new(UILabel)

	label.id = uuid.New()
	label.text = render.CreateText(x, y, width, height, content, style)

	return label
}

func CreateSingleLineUILabel(x, y uint16, content string, style tcell.Style) *UILabel {
	label := new(UILabel)

	label.id = uuid.New()
	label.text = render.CreateText(x, y, uint16(utf8.RuneCountInString(content)), 1, content, style)

	return label
}

func (t *UILabel) UniqueId() uuid.UUID {
	return t.id
}

func (t *UILabel) MoveTo(x uint16, y uint16) {
	t.text = render.CreateText(x, y, uint16(t.text.Size().Width()), uint16(t.Size().Height()), t.text.Content(), t.text.Style())
}

func (t *UILabel) Position() util.Position {
	return t.text.Position()
}

func (t *UILabel) Size() util.Size {
	return t.text.Size()
}

func (t *UILabel) Draw(v views.View) {
	t.text.Draw(v)
}

func (t *UILabel) Input(e *tcell.EventKey) {}
