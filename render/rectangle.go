package render

import (
	"mvvasilev/last_light/util"

	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
)

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
