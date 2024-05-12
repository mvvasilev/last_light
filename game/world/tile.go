package world

import (
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/model"

	"github.com/gdamore/tcell/v2"
)

type Material uint

const (
	MaterialGround Material = iota
	MaterialRock
	MaterialWall
	MaterialGrass
	MaterialVoid
	MaterialClosedDoor
	MaterialOpenDoor
	MaterialStaircaseDown
	MaterialStaircaseUp
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

func TileTypeClosedDoor() TileType {
	return TileType{
		Material:     MaterialClosedDoor,
		Passable:     false,
		Transparent:  false,
		Presentation: '[',
		Style:        tcell.StyleDefault.Foreground(tcell.ColorLightSteelBlue).Background(tcell.ColorSaddleBrown),
	}
}

func TileTypeOpenDoor() TileType {
	return TileType{
		Material:     MaterialClosedDoor,
		Passable:     false,
		Transparent:  false,
		Presentation: '_',
		Style:        tcell.StyleDefault.Foreground(tcell.ColorLightSteelBlue),
	}
}

func TileTypeStaircaseDown() TileType {
	return TileType{
		Material:     MaterialStaircaseDown,
		Passable:     true,
		Transparent:  false,
		Presentation: '≡',
		Style:        tcell.StyleDefault.Foreground(tcell.ColorDarkSlateGray).Attributes(tcell.AttrBold),
	}
}

func TileTypeStaircaseUp() TileType {
	return TileType{
		Material:     MaterialStaircaseUp,
		Passable:     true,
		Transparent:  false,
		Presentation: '^',
		Style:        tcell.StyleDefault.Foreground(tcell.ColorDarkSlateGray).Attributes(tcell.AttrBold),
	}
}

type Tile interface {
	Position() engine.Position
	Presentation() (rune, tcell.Style)
	Passable() bool
	Transparent() bool
}

type StaticTile struct {
	position engine.Position
	t        TileType
}

func CreateStaticTile(x, y int, t TileType) Tile {
	st := new(StaticTile)

	st.position = engine.PositionAt(x, y)
	st.t = t

	return st
}

func (st *StaticTile) Position() engine.Position {
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
	position engine.Position
	itemType *model.ItemType
	quantity int
}

func CreateItemTile(position engine.Position, itemType *model.ItemType, quantity int) *ItemTile {
	it := new(ItemTile)

	it.position = position
	it.itemType = itemType
	it.quantity = quantity

	return it
}

func (it *ItemTile) Type() *model.ItemType {
	return it.itemType
}

func (it *ItemTile) Quantity() int {
	return it.quantity
}

func (it *ItemTile) Position() engine.Position {
	return it.position
}

func (it *ItemTile) Presentation() (rune, tcell.Style) {
	return it.itemType.TileIcon(), it.itemType.Style()
}

func (it *ItemTile) Passable() bool {
	return true
}

func (it *ItemTile) Transparent() bool {
	return false
}

type EntityTile interface {
	Entity() model.MovableEntity
	Tile
}

type BasicEntityTile struct {
	entity model.MovableEntity

	presentation rune
	style        tcell.Style
}

func CreateBasicEntityTile(entity model.MovableEntity, presentation rune, style tcell.Style) *BasicEntityTile {
	return &BasicEntityTile{
		entity:       entity,
		presentation: presentation,
		style:        style,
	}
}

func (bet *BasicEntityTile) Entity() model.MovableEntity {
	return bet.entity
}

func (bet *BasicEntityTile) Position() engine.Position {
	return bet.entity.Position()
}

func (bet *BasicEntityTile) Presentation() (rune, tcell.Style) {
	return bet.presentation, bet.style
}

func (bet *BasicEntityTile) Passable() bool {
	return false
}

func (bet *BasicEntityTile) Transparent() bool {
	return false
}
