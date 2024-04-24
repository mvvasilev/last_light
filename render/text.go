package render

import (
	"mvvasilev/last_light/util"
	"strings"
	"unicode/utf8"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	"github.com/google/uuid"
)

type Text struct {
	id       uuid.UUID
	content  []string
	position util.Position
	size     util.Size
	style    tcell.Style
}

func CreateText(
	x, y uint16,
	width, height uint16,
	content string,
	style tcell.Style,
) *Text {
	text := new(Text)

	text.id = uuid.New()
	text.content = strings.Split(content, " ")
	text.style = style
	text.size = util.SizeOf(width, height)
	text.position = util.PositionAt(x, y)

	return text
}

func (t *Text) UniqueId() uuid.UUID {
	return t.id
}

func (t *Text) Position() util.Position {
	return t.position
}

func (t *Text) Content() string {
	return strings.Join(t.content, " ")
}

func (t *Text) Size() util.Size {
	return t.size
}

func (t *Text) SetStyle(style tcell.Style) {
	t.style = style
}

func (t *Text) Style() tcell.Style {
	return t.style
}

func (t *Text) Draw(s views.View) {
	width := t.size.Width()
	height := t.size.Height()
	x := t.position.X()
	y := t.position.Y()

	currentHPos := 0
	currentVPos := 0

	drawText := func(text string) {
		for i, r := range text {
			s.SetContent(x+currentHPos+i, y+currentVPos, r, nil, t.style)
		}
	}

	for _, s := range t.content {
		runeCount := utf8.RuneCountInString(s)

		if currentVPos > height {
			break
		}

		// The current word cannot fit within the remaining space on the line
		if runeCount > (width - currentHPos) {
			currentVPos += 1 // next line
			currentHPos = 0  // reset to start of line

			drawText(s + " ")
			currentHPos += runeCount + 1

			continue
		}

		// The current word fits exactly within the remaining space on the line
		if runeCount == (width - currentHPos) {
			drawText(s)

			currentVPos += 1 // next line
			currentHPos = 0  // reset to start of line

			continue
		}

		// The current word fits within the remaining space, and there's more space left over
		drawText(s + " ")
		currentHPos += runeCount + 1 // add +1 to account for space after word
	}
}
