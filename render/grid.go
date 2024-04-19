package render

import (
	"mvvasilev/last_light/util"

	"github.com/gdamore/tcell/v2"
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

func (g grid) UniqueId() uuid.UUID {
	return g.id
}

func (g grid) Draw(s tcell.Screen) {

}
