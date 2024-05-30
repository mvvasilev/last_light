package engine

import (
	"strings"
	"unicode/utf8"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	"github.com/google/uuid"
)

type Text struct {
	id       uuid.UUID
	content  []string
	position Position
	size     Size
	style    tcell.Style
}

func CreateText(
	x, y int,
	width, height int,
	content string,
	style tcell.Style,
) *Text {
	text := new(Text)

	text.id = uuid.New()
	text.content = strings.Split(content, " ")
	text.style = style
	text.size = SizeOf(width, height)
	text.position = PositionAt(x, y)

	return text
}

func (t *Text) UniqueId() uuid.UUID {
	return t.id
}

func (t *Text) Position() Position {
	return t.position
}

func (t *Text) Content() string {
	return strings.Join(t.content, " ")
}

func (t *Text) Size() Size {
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
		lastPos := 0

		for _, r := range text {
			s.SetContent(x+currentHPos+lastPos, y+currentVPos, r, nil, t.style)
			lastPos++
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

func DrawText(x, y int, content string, style tcell.Style, s views.View) {
	currentHPos := 0
	currentVPos := 0

	lastPos := 0

	for _, r := range content {
		s.SetContent(x+currentHPos+lastPos, y+currentVPos, r, nil, style)
		lastPos++
	}
}
