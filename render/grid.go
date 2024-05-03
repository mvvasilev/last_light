package render

import (
	"mvvasilev/last_light/util"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	"github.com/google/uuid"
)

type Grid struct {
	id uuid.UUID

	internalCellSize   util.Size
	numCellsHorizontal int
	numCellsVertical   int
	position           util.Position
	style              tcell.Style
	highlightStyle     tcell.Style

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

	isHighlighted   bool
	highlightedGrid util.Position

	fillRune rune
}

func CreateSimpleGrid(
	x, y int,
	cellWidth, cellHeight int,
	numCellsHorizontal, numCellsVertical int,
	borderRune, fillRune rune,
	style tcell.Style, highlightStyle tcell.Style,
) *Grid {
	return CreateGrid(
		x, y, cellWidth, cellHeight, numCellsHorizontal, numCellsVertical,
		borderRune, borderRune, borderRune, borderRune,
		borderRune, fillRune, borderRune, borderRune,
		borderRune, borderRune, borderRune, borderRune,
		borderRune, borderRune, borderRune, borderRune,
		style, highlightStyle,
	)
}

// '┌', '─', '┬', '┐',
// '│', '#', '│', '│',
// '├', '─', '┼', '┤',
// '└', '─', '┴', '┘',
func CreateGrid(
	x int,
	y int,
	cellWidth int,
	cellHeight int,
	numCellsHorizontal int,
	numCellsVertical int,
	nwCorner, northBorder, verticalDownwardsTJunction, neCorner,
	westBorder, fillRune, internalVerticalBorder, eastBorder,
	horizontalRightTJunction, internalHorizontalBorder, crossJunction, horizontalLeftTJunction,
	swCorner, southBorder, verticalUpwardsTJunction, seCorner rune,
	style tcell.Style, highlightStyle tcell.Style,
) *Grid {
	return &Grid{
		id:                         uuid.New(),
		internalCellSize:           util.SizeOf(cellWidth, cellHeight),
		numCellsHorizontal:         numCellsHorizontal,
		numCellsVertical:           numCellsVertical,
		isHighlighted:              false,
		position:                   util.PositionAt(x, y),
		style:                      style,
		highlightStyle:             highlightStyle,
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

func (g *Grid) UniqueId() uuid.UUID {
	return g.id
}

func (g *Grid) Highlight(highlightedGrid util.Position) {
	g.isHighlighted = true
	g.highlightedGrid = highlightedGrid
}

func (g *Grid) Unhighlight() {
	g.isHighlighted = false
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
func (g *Grid) drawBorders(v views.View) {
	iCellSizeWidth := g.internalCellSize.Width()
	iCellSizeHeight := g.internalCellSize.Height()
	width := 1 + (iCellSizeWidth * int(g.numCellsHorizontal)) + (int(g.numCellsHorizontal))
	height := 1 + (iCellSizeHeight * int(g.numCellsVertical)) + (int(g.numCellsVertical))
	x := g.position.X()
	y := g.position.Y()

	style := g.style

	for w := 0; w < width; w++ {
		for iw := 1; iw < g.numCellsVertical; iw++ {
			v.SetContent(x+w, y+(iw*iCellSizeHeight+iw), g.internalHorizontalBorder, nil, style)
		}

		if w%(iCellSizeWidth+1) == 0 {
			v.SetContent(x+w, y, g.verticalDownwardsTJunction, nil, style)
			v.SetContent(x+w, y+height-1, g.verticalUpwardsTJunction, nil, style)
			continue
		}

		v.SetContent(x+w, y, g.northBorder, nil, style)
		v.SetContent(x+w, y+height-1, g.southBorder, nil, style)
	}

	for h := 0; h < height; h++ {
		if h == 0 {
			v.SetContent(x, y, g.nwCorner, nil, style)
			v.SetContent(x, y+height-1, g.swCorner, nil, style)
			continue
		}

		if h == height-1 {
			v.SetContent(x+width-1, y, g.neCorner, nil, style)
			v.SetContent(x+width-1, y+height-1, g.seCorner, nil, style)
			continue
		}

		for ih := 1; ih < g.numCellsHorizontal; ih++ {
			if h%(iCellSizeHeight+1) == 0 {
				v.SetContent(x+(ih*iCellSizeWidth+ih), y+h, g.crossJunction, nil, style)
				continue
			}

			v.SetContent(x+(ih*iCellSizeWidth+ih), y+h, g.internalVerticalBorder, nil, style)
		}

		if h%(iCellSizeHeight+1) == 0 {
			v.SetContent(x, y+h, g.horizontalRightTJunction, nil, style)
			v.SetContent(x+width-1, y+h, g.horizontalLeftTJunction, nil, style)
			continue
		}

		v.SetContent(x, y+h, g.westBorder, nil, style)
		v.SetContent(x+width-1, y+h, g.eastBorder, nil, style)
	}
}

func (g *Grid) drawFill(v views.View) {

}

func (g *Grid) Draw(v views.View) {
	g.drawBorders(v)
	g.drawFill(v)
}
