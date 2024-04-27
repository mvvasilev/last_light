package ui

import (
	"mvvasilev/last_light/render"
	"mvvasilev/last_light/util"
	"strings"
	"unicode/utf8"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	"github.com/google/uuid"
)

type UISimpleButton struct {
	id                 uuid.UUID
	isHighlighted      bool
	text               *render.Text
	selectHandler      func()
	unhighlightedStyle tcell.Style
	highlightedStyle   tcell.Style
}

func CreateSimpleButton(x, y int, text string, unhighlightedStyle, highlightedStyle tcell.Style, onSelect func()) *UISimpleButton {
	sb := new(UISimpleButton)

	sb.id = uuid.New()
	sb.text = render.CreateText(x, y, int(utf8.RuneCountInString(text)), 1, text, unhighlightedStyle)
	sb.isHighlighted = false
	sb.selectHandler = onSelect
	sb.highlightedStyle = highlightedStyle
	sb.unhighlightedStyle = unhighlightedStyle

	return sb
}

func (sb *UISimpleButton) Select() {
	sb.selectHandler()
}

func (sb *UISimpleButton) OnSelect(f func()) {
	sb.selectHandler = f
}

func (sb *UISimpleButton) IsHighlighted() bool {
	return sb.isHighlighted
}

func (sb *UISimpleButton) Highlight() {
	sb.isHighlighted = true

	newContent := "[ " + sb.text.Content() + " ]"

	sb.text = render.CreateText(
		int(sb.Position().X()-2), int(sb.Position().Y()),
		int(utf8.RuneCountInString(newContent)), 1,
		newContent,
		sb.highlightedStyle,
	)
}

func (sb *UISimpleButton) Unhighlight() {
	sb.isHighlighted = false

	content := strings.Trim(sb.text.Content(), " ]")
	content = strings.Trim(content, "[ ")
	contentLen := utf8.RuneCountInString(content)

	sb.text = render.CreateText(
		int(sb.Position().X()+2), int(sb.Position().Y()),
		int(contentLen), 1,
		content,
		sb.unhighlightedStyle,
	)
}

func (sb *UISimpleButton) SetHighlighted(highlighted bool) {
	sb.isHighlighted = highlighted
}

func (sb *UISimpleButton) UniqueId() uuid.UUID {
	return sb.id
}

func (sb *UISimpleButton) MoveTo(x int, y int) {
	sb.text = render.CreateText(x, y, int(utf8.RuneCountInString(sb.text.Content())), 1, sb.text.Content(), sb.highlightedStyle)
}

func (sb *UISimpleButton) Position() util.Position {
	return sb.text.Position()
}

func (sb *UISimpleButton) Size() util.Size {
	return sb.text.Size()
}

func (sb *UISimpleButton) Draw(v views.View) {
	sb.text.Draw(v)
}

func (sb *UISimpleButton) Input(e *tcell.EventKey) {

}
