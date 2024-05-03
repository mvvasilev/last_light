package model

import (
	"mvvasilev/last_light/util"

	"github.com/gdamore/tcell/v2"
)

type Material uint

const (
	MaterialGround Material = iota
	MaterialRock
	MaterialWall
	MaterialGrass
	MaterialVoid
)

type TileType struct {
	Material     Material
	Passable     bool
	Presentation rune
	Transparent  bool
	Style        tcell.Style
}

func TileTypeGround() TileType {
	return TileType{
		Material:     MaterialGround,
		Passable:     true,
		Presentation: '.',
		Transparent:  false,
		Style:        tcell.StyleDefault,
	}
}

func TileTypeRock() TileType {
	return TileType{
		Material:     MaterialRock,
		Passable:     false,
		Presentation: '█',
		Transparent:  false,
		Style:        tcell.StyleDefault,
	}
}

func TileTypeGrass() TileType {
	return TileType{
		Material:     MaterialGrass,
		Passable:     true,
		Presentation: ',',
		Transparent:  false,
		Style:        tcell.StyleDefault,
	}
}

func TileTypeVoid() TileType {
	return TileType{
		Material:     MaterialVoid,
		Passable:     false,
		Presentation: ' ',
		Transparent:  true,
		Style:        tcell.StyleDefault,
	}
}

func TileTypeWall() TileType {
	return TileType{
		Material:     MaterialWall,
		Passable:     false,
		Presentation: '#',
		Transparent:  false,
		Style:        tcell.StyleDefault.Background(tcell.ColorGray),
	}
}

type Tile interface {
	Position() util.Position
	Presentation() (rune, tcell.Style)
	Passable() bool
	Transparent() bool
}

type StaticTile struct {
	position util.Position
	t        TileType
}

func CreateStaticTile(x, y int, t TileType) Tile {
	st := new(StaticTile)

	st.position = util.PositionAt(x, y)
	st.t = t

	return st
}

func (st *StaticTile) Position() util.Position {
	return st.position
}

func (st *StaticTile) Presentation() (rune, tcell.Style) {
	return st.t.Presentation, st.t.Style
}

func (st *StaticTile) Passable() bool {
	return st.t.Passable
}

func (st *StaticTile) Transparent() bool {
	return st.t.Transparent
}

func (st *StaticTile) Type() TileType {
	return st.t
}

type ItemTile struct {
	position util.Position
	itemType *ItemType
	quantity int
}

func CreateItemTile(position util.Position, itemType *ItemType, quantity int) *ItemTile {
	it := new(ItemTile)

	it.position = position
	it.itemType = itemType
	it.quantity = quantity

	return it
}

func (it *ItemTile) Type() *ItemType {
	return it.itemType
}

func (it *ItemTile) Quantity() int {
	return it.quantity
}

func (it *ItemTile) Position() util.Position {
	return it.position
}

func (it *ItemTile) Presentation() (rune, tcell.Style) {
	return it.itemType.tileIcon, it.itemType.style
}

func (it *ItemTile) Passable() bool {
	return true
}

func (it *ItemTile) Transparent() bool {
	return false
}
