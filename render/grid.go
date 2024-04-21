package render

import (
	"mvvasilev/last_light/util"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	"github.com/google/uuid"
)

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

	internalVerticalBorder   rune
	internalHorizontalBorder rune

	nwCorner rune
	swCorner rune
	seCorner rune
	neCorner rune

	verticalDownwardsTJunction rune
	verticalUpwardsTJunction   rune
	horizontalLeftTJunction    rune
	horizontalRightTJunction   rune
	crossJunction              rune

	fillRune rune
}

func CreateSimpleGrid(
	x, y uint16,
	cellWidth, cellHeight uint16,
	numCellsHorizontal, numCellsVertical uint16,
	borderRune, fillRune rune,
	style tcell.Style,
) grid {
	return CreateGrid(
		x, y, cellWidth, cellHeight, numCellsHorizontal, numCellsVertical,
		borderRune, borderRune, borderRune, borderRune,
		borderRune, fillRune, borderRune, borderRune,
		borderRune, borderRune, borderRune, borderRune,
		borderRune, borderRune, borderRune, borderRune,
		style,
	)
}

// '┌', '─', '┬', '┐',
// '│', '#', '│', '│',
// '├', '─', '┼', '┤',
// '└', '─', '┴', '┘',
func CreateGrid(
	x uint16,
	y uint16,
	cellWidth uint16,
	cellHeight uint16,
	numCellsHorizontal uint16,
	numCellsVertical uint16,
	nwCorner, northBorder, verticalDownwardsTJunction, neCorner,
	westBorder, fillRune, internalVerticalBorder, eastBorder,
	horizontalRightTJunction, internalHorizontalBorder, crossJunction, horizontalLeftTJunction,
	swCorner, southBorder, verticalUpwardsTJunction, seCorner rune,
	style tcell.Style,
) grid {
	return grid{
		id:                         uuid.New(),
		internalCellSize:           util.SizeOf(cellWidth, cellHeight),
		numCellsHorizontal:         numCellsHorizontal,
		numCellsVertical:           numCellsVertical,
		position:                   util.PositionAt(x, y),
		style:                      style,
		northBorder:                northBorder,
		eastBorder:                 eastBorder,
		southBorder:                southBorder,
		westBorder:                 westBorder,
		internalVerticalBorder:     internalVerticalBorder,
		internalHorizontalBorder:   internalHorizontalBorder,
		nwCorner:                   nwCorner,
		seCorner:                   seCorner,
		swCorner:                   swCorner,
		neCorner:                   neCorner,
		verticalDownwardsTJunction: verticalDownwardsTJunction,
		verticalUpwardsTJunction:   verticalUpwardsTJunction,
		horizontalRightTJunction:   horizontalRightTJunction,
		horizontalLeftTJunction:    horizontalLeftTJunction,
		fillRune:                   fillRune,

		crossJunction: crossJunction,
	}
}

func (g grid) UniqueId() uuid.UUID {
	return g.id
}

// C###T###T###C
// #   #   #   #
// #   #   #   #
// #   #   #   #
// T###X###X###T
// #   #   #   #
// #   #   #   #
// #   #   #   #
// T###X###X###T
// #   #   #   #
// #   #   #   #
// #   #   #   #
// C###T###T###C
func (g grid) drawBorders(v views.View) {
	width := 2 + (g.internalCellSize.Width() * int(g.numCellsHorizontal)) + (int(g.numCellsHorizontal) - 1)
	height := 2 + (g.internalCellSize.Height() * int(g.numCellsVertical)) + (int(g.numCellsVertical) - 1)
	x := g.position.X()
	y := g.position.Y()

	v.SetContent(x, y, g.nwCorner, nil, g.style)
	v.SetContent(x+width-1, y, g.neCorner, nil, g.style)
	v.SetContent(x, y+height-1, g.swCorner, nil, g.style)
	v.SetContent(x+width-1, y+height-1, g.seCorner, nil, g.style)

	for w := range width - 2 {
		v.SetContent(1+w, y, g.northBorder, nil, g.style)
		v.SetContent(1+w, y+height-1, g.southBorder, nil, g.style)
	}

	for h := range height - 2 {
		v.SetContent(x, 1+h, g.westBorder, nil, g.style)
		v.SetContent(x+width-1, 1+h, g.eastBorder, nil, g.style)
	}
}

func (g grid) drawFill(v views.View) {

}

func (g grid) Draw(v views.View) {
	g.drawBorders(v)
	g.drawFill(v)
}
