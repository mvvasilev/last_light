package render

import (
	"mvvasilev/last_light/util"
	"strings"
	"unicode/utf8"

	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
)

type Drawable interface {
	UniqueId() uuid.UUID
	Draw(s tcell.Screen)
}

type rectangle struct {
	id uuid.UUID

	size     util.Size
	position util.Position
	style    tcell.Style

	northBorder rune
	westBorder  rune
	eastBorder  rune
	southBorder rune

	nwCorner rune
	swCorner rune
	seCorner rune
	neCorner rune

	fillRune rune
}

func CreateSimpleRectangle(x uint16, y uint16, width uint16, height uint16, borderRune rune, fillRune rune, style tcell.Style) rectangle {
	return CreateRectangle(
		x, y, width, height,
		borderRune, borderRune, borderRune,
		borderRune, fillRune, borderRune,
		borderRune, borderRune, borderRune,
		style,
	)
}

// CreateRectangle(
//
//	x, y, width, height,
//	'┌', '─', '┐',
//	'│', ' ', '│',
//	'└', '─', '┘',
//	style
//
// )
func CreateRectangle(
	x uint16,
	y uint16,
	width uint16,
	height uint16,
	nwCorner, northBorder, neCorner,
	westBorder, fillRune, eastBorder,
	swCorner, southBorder, seCorner rune,
	style tcell.Style,
) rectangle {
	return rectangle{
		id:          uuid.New(),
		size:        util.SizeOf(width, height),
		position:    util.PositionAt(x, y),
		style:       style,
		northBorder: northBorder,
		eastBorder:  eastBorder,
		southBorder: southBorder,
		westBorder:  westBorder,
		nwCorner:    nwCorner,
		seCorner:    seCorner,
		swCorner:    swCorner,
		neCorner:    neCorner,
		fillRune:    fillRune,
	}
}

func (rect rectangle) UniqueId() uuid.UUID {
	return rect.id
}

func (rect rectangle) Draw(s tcell.Screen) {
	width := rect.size.Width()
	height := rect.size.Height()
	x := rect.position.X()
	y := rect.position.Y()

	for h := range height {
		for w := range width {

			// nw corner
			if w == 0 && h == 0 {
				s.SetContent(x+w, y+h, rect.nwCorner, nil, rect.style)
				continue
			}

			// ne corner
			if w == (width-1) && h == 0 {
				s.SetContent(x+w, y+h, rect.neCorner, nil, rect.style)
				continue
			}

			// sw corner
			if w == 0 && h == (height-1) {
				s.SetContent(x+w, y+h, rect.swCorner, nil, rect.style)
				continue
			}

			// se corner
			if w == (width-1) && h == (height-1) {
				s.SetContent(x+w, y+h, rect.seCorner, nil, rect.style)
				continue
			}

			// north border
			if h == 0 && (w != 0 && w != (width-1)) {
				s.SetContent(x+w, y+h, rect.northBorder, nil, rect.style)
				continue
			}

			// south border
			if h == (height-1) && (w != 0 && w != (width-1)) {
				s.SetContent(x+w, y+h, rect.southBorder, nil, rect.style)
				continue
			}

			// west border
			if w == 0 && (h != 0 && h != (height-1)) {
				s.SetContent(x+w, y+h, rect.westBorder, nil, rect.style)
				continue
			}

			// east border
			if w == (width-1) && (h != 0 && h != (height-1)) {
				s.SetContent(x+w, y+h, rect.eastBorder, nil, rect.style)
				continue
			}

			s.SetContent(x+w, y+h, rect.fillRune, nil, rect.style)
		}
	}

}

type text struct {
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
) text {
	return text{
		id:       uuid.New(),
		content:  strings.Split(content, " "),
		style:    style,
		size:     util.SizeOf(width, height),
		position: util.PositionAt(x, y),
	}
}

func (t text) UniqueId() uuid.UUID {
	return t.id
}

func (t text) Draw(s tcell.Screen) {
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

type grid struct {
	id uuid.UUID

	internalCellSize   util.Size
	numCellsHorizontal uint16
	numCellsVertical   uint16
	position           util.Position
	style              tcell.Style

	northBorder rune
	westBorder  rune
	eastBorder  rune
	southBorder rune

	nwCorner rune
	swCorner rune
	seCorner rune
	neCorner rune

	verticalTJunction   rune
	horizontalTJunction rune
	crossJunction       rune

	fillRune rune
}

func CreateGrid(
	x uint16,
	y uint16,
	cellWidth uint16,
	cellHeight uint16,
	numCellsHorizontal uint16,
	numCellsVertical uint16,
	nwCorner, northBorder, neCorner,
	westBorder, fillRune, eastBorder,
	swCorner, southBorder, seCorner,
	verticalTJunction, horizontalTJunction,
	crossJunction rune,
	style tcell.Style,
) grid {
	return grid{
		id:                  uuid.New(),
		internalCellSize:    util.SizeOf(cellWidth, cellHeight),
		numCellsHorizontal:  numCellsHorizontal,
		numCellsVertical:    numCellsVertical,
		position:            util.PositionAt(x, y),
		style:               style,
		northBorder:         northBorder,
		eastBorder:          eastBorder,
		southBorder:         southBorder,
		westBorder:          westBorder,
		nwCorner:            nwCorner,
		seCorner:            seCorner,
		swCorner:            swCorner,
		neCorner:            neCorner,
		fillRune:            fillRune,
		verticalTJunction:   verticalTJunction,
		horizontalTJunction: horizontalTJunction,
		crossJunction:       crossJunction,
	}
}

func (g grid) Draw(s tcell.Screen) {

}
