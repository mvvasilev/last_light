package render

import (
	"mvvasilev/last_light/util"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
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

	isBorderless bool
	isFilled     bool

	fillRune rune
}

func CreateBorderlessRectangle(x, y uint16, width, height uint16, fillRune rune, style tcell.Style) rectangle {
	return CreateRectangle(
		x, y, width, height,
		0, 0, 0,
		0, fillRune, 0,
		0, 0, 0,
		true, true, style,
	)
}

func CreateSimpleEmptyRectangle(x, y uint16, width, height uint16, borderRune rune, style tcell.Style) rectangle {
	return CreateRectangle(
		x, y, width, height,
		borderRune, borderRune, borderRune,
		borderRune, 0, borderRune,
		borderRune, borderRune, borderRune,
		false, false, style,
	)
}

func CreateSimpleRectangle(x uint16, y uint16, width uint16, height uint16, borderRune rune, fillRune rune, style tcell.Style) rectangle {
	return CreateRectangle(
		x, y, width, height,
		borderRune, borderRune, borderRune,
		borderRune, fillRune, borderRune,
		borderRune, borderRune, borderRune,
		false, true, style,
	)
}

// CreateRectangle(
//
//		x, y, width, height,
//		'┌', '─', '┐',
//		'│', ' ', '│',
//		'└', '─', '┘',
//	 false, true,
//		style
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
	isBorderless, isFilled bool,
	style tcell.Style,
) rectangle {
	return rectangle{
		id:           uuid.New(),
		size:         util.SizeOf(width, height),
		position:     util.PositionAt(x, y),
		style:        style,
		northBorder:  northBorder,
		eastBorder:   eastBorder,
		southBorder:  southBorder,
		westBorder:   westBorder,
		nwCorner:     nwCorner,
		seCorner:     seCorner,
		swCorner:     swCorner,
		neCorner:     neCorner,
		isBorderless: isBorderless,
		isFilled:     isFilled,
		fillRune:     fillRune,
	}
}

func (rect rectangle) UniqueId() uuid.UUID {
	return rect.id
}

func (rect rectangle) drawBorders(v views.View) {
	width := rect.size.Width()
	height := rect.size.Height()
	x := rect.position.X()
	y := rect.position.Y()

	v.SetContent(x, y, rect.nwCorner, nil, rect.style)
	v.SetContent(x+width-1, y, rect.neCorner, nil, rect.style)
	v.SetContent(x, y+height-1, rect.swCorner, nil, rect.style)
	v.SetContent(x+width-1, y+height-1, rect.seCorner, nil, rect.style)

	for w := range width - 2 {
		v.SetContent(1+w, y, rect.northBorder, nil, rect.style)
		v.SetContent(1+w, y+height-1, rect.southBorder, nil, rect.style)
	}

	for h := range height - 2 {
		v.SetContent(x, 1+h, rect.westBorder, nil, rect.style)
		v.SetContent(x+width-1, 1+h, rect.eastBorder, nil, rect.style)
	}
}

func (rect rectangle) drawFill(v views.View) {
	for w := range rect.size.Width() - 2 {
		for h := range rect.size.Height() - 2 {
			v.SetContent(1+w, 1+h, rect.fillRune, nil, rect.style)
		}
	}
}

func (rect rectangle) Draw(v views.View) {
	if !rect.isBorderless {
		rect.drawBorders(v)
	}

	if rect.isFilled {
		rect.drawFill(v)
	}
}
