package render

import (
	"mvvasilev/last_light/util"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	"github.com/google/uuid"
)

type Rectangle struct {
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

func CreateBorderlessRectangle(x, y int, width, height int, fillRune rune, style tcell.Style) Rectangle {
	return CreateRectangle(
		x, y, width, height,
		0, 0, 0,
		0, fillRune, 0,
		0, 0, 0,
		true, true, style,
	)
}

func CreateSimpleEmptyRectangle(x, y int, width, height int, borderRune rune, style tcell.Style) Rectangle {
	return CreateRectangle(
		x, y, width, height,
		borderRune, borderRune, borderRune,
		borderRune, 0, borderRune,
		borderRune, borderRune, borderRune,
		false, false, style,
	)
}

func CreateSimpleRectangle(x int, y int, width int, height int, borderRune rune, fillRune rune, style tcell.Style) Rectangle {
	return CreateRectangle(
		x, y, width, height,
		borderRune, borderRune, borderRune,
		borderRune, fillRune, borderRune,
		borderRune, borderRune, borderRune,
		false, true, style,
	)
}

func CreateRectangleV2(
	x, y int, width, height int,
	upper, middle, lower string,
	isBorderless, isFilled bool,
	style tcell.Style,
) Rectangle {
	upperRunes := []rune(upper)
	middleRunes := []rune(middle)
	lowerRunes := []rune(lower)

	return CreateRectangle(
		x, y, width, height,
		upperRunes[0], upperRunes[1], upperRunes[2],
		middleRunes[0], middleRunes[1], middleRunes[2],
		lowerRunes[0], lowerRunes[1], lowerRunes[2],
		isBorderless, isFilled, style,
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
	x int,
	y int,
	width int,
	height int,
	nwCorner, northBorder, neCorner,
	westBorder, fillRune, eastBorder,
	swCorner, southBorder, seCorner rune,
	isBorderless, isFilled bool,
	style tcell.Style,
) Rectangle {
	return Rectangle{
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

func (rect Rectangle) UniqueId() uuid.UUID {
	return rect.id
}

func (rect Rectangle) Position() util.Position {
	return rect.position
}

func (rect Rectangle) drawBorders(v views.View) {
	width := rect.size.Width()
	height := rect.size.Height()
	x := rect.position.X()
	y := rect.position.Y()

	v.SetContent(x, y, rect.nwCorner, nil, rect.style)
	v.SetContent(x+width-1, y, rect.neCorner, nil, rect.style)
	v.SetContent(x, y+height-1, rect.swCorner, nil, rect.style)
	v.SetContent(x+width-1, y+height-1, rect.seCorner, nil, rect.style)

	for w := 1; w < width-1; w++ {
		v.SetContent(x+w, y, rect.northBorder, nil, rect.style)
		v.SetContent(x+w, y+height-1, rect.southBorder, nil, rect.style)
	}

	for h := 1; h < height-1; h++ {
		v.SetContent(x, y+h, rect.westBorder, nil, rect.style)
		v.SetContent(x+width-1, y+h, rect.eastBorder, nil, rect.style)
	}
}

func (rect Rectangle) drawFill(v views.View) {
	width := rect.size.Width()
	height := rect.size.Height()
	x := rect.position.X()
	y := rect.position.Y()

	for w := 1; w < width-1; w++ {
		for h := 1; h < height-1; h++ {
			v.SetContent(x+w, y+h, rect.fillRune, nil, rect.style)
		}
	}
}

func (rect Rectangle) Draw(v views.View) {
	if !rect.isBorderless {
		rect.drawBorders(v)
	}

	if rect.isFilled {
		rect.drawFill(v)
	}
}
