package world

import (
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/item"
	"mvvasilev/last_light/game/npc"

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
	Opaque       bool
	Style        tcell.Style
}

func TileTypeGround() TileType {
	return TileType{
		Material:     MaterialGround,
		Passable:     true,
		Presentation: '.',
		Transparent:  false,
		Opaque:       false,
		Style:        tcell.StyleDefault,
	}
}

func TileTypeRock() TileType {
	return TileType{
		Material:     MaterialRock,
		Passable:     false,
		Presentation: '█',
		Transparent:  false,
		Opaque:       true,
		Style:        tcell.StyleDefault,
	}
}

func TileTypeGrass() TileType {
	return TileType{
		Material:     MaterialGrass,
		Passable:     true,
		Presentation: ',',
		Transparent:  false,
		Opaque:       false,
		Style:        tcell.StyleDefault,
	}
}

func TileTypeVoid() TileType {
	return TileType{
		Material:     MaterialVoid,
		Passable:     false,
		Presentation: ' ',
		Transparent:  true,
		Opaque:       true,
		Style:        tcell.StyleDefault,
	}
}

func TileTypeWall() TileType {
	return TileType{
		Material:     MaterialWall,
		Passable:     false,
		Presentation: '#',
		Transparent:  false,
		Opaque:       true,
		Style:        tcell.StyleDefault.Background(tcell.ColorGray),
	}
}

func TileTypeClosedDoor() TileType {
	return TileType{
		Material:     MaterialClosedDoor,
		Passable:     false,
		Transparent:  false,
		Presentation: '[',
		Opaque:       true,
		Style:        tcell.StyleDefault.Foreground(tcell.ColorLightSteelBlue).Background(tcell.ColorSaddleBrown),
	}
}

func TileTypeOpenDoor() TileType {
	return TileType{
		Material:     MaterialClosedDoor,
		Passable:     false,
		Transparent:  false,
		Presentation: '_',
		Opaque:       false,
		Style:        tcell.StyleDefault.Foreground(tcell.ColorLightSteelBlue),
	}
}

func TileTypeStaircaseDown() TileType {
	return TileType{
		Material:     MaterialStaircaseDown,
		Passable:     true,
		Transparent:  false,
		Presentation: '≡',
		Opaque:       false,
		Style:        tcell.StyleDefault.Foreground(tcell.ColorDarkSlateGray).Attributes(tcell.AttrBold),
	}
}

func TileTypeStaircaseUp() TileType {
	return TileType{
		Material:     MaterialStaircaseUp,
		Passable:     true,
		Transparent:  false,
		Presentation: '^',
		Opaque:       false,
		Style:        tcell.StyleDefault.Foreground(tcell.ColorDarkSlateGray).Attributes(tcell.AttrBold),
	}
}

type Tile interface {
	Position() engine.Position
	Presentation() (rune, tcell.Style)
	Passable() bool
	Transparent() bool
	Opaque() bool
	Type() TileType
}

type StaticTile struct {
	position engine.Position
	t        TileType

	style tcell.Style
}

func CreateStaticTile(x, y int, t TileType) Tile {
	st := new(StaticTile)

	st.position = engine.PositionAt(x, y)
	st.t = t
	st.style = t.Style

	return st
}

func CreateStaticTileWithStyleOverride(x, y int, t TileType, style tcell.Style) Tile {
	return &StaticTile{
		position: engine.PositionAt(x, y),
		t:        t,
		style:    style,
	}
}

func (st *StaticTile) Position() engine.Position {
	return st.position
}

func (st *StaticTile) Presentation() (rune, tcell.Style) {
	return st.t.Presentation, st.style
}

func (st *StaticTile) Passable() bool {
	return st.t.Passable
}

func (st *StaticTile) Transparent() bool {
	return st.t.Transparent
}

func (st *StaticTile) Opaque() bool {
	return st.t.Opaque
}

func (st *StaticTile) Type() TileType {
	return st.t
}

type ItemTile struct {
	position engine.Position
	item     item.Item
}

func CreateItemTile(position engine.Position, item item.Item) *ItemTile {
	it := new(ItemTile)

	it.position = position
	it.item = item

	return it
}

func (it *ItemTile) Item() item.Item {
	return it.item
}

func (it *ItemTile) Position() engine.Position {
	return it.position
}

func (it *ItemTile) Presentation() (rune, tcell.Style) {
	return it.item.Type().TileIcon(), it.item.Type().Style()
}

func (it *ItemTile) Passable() bool {
	return true
}

func (it *ItemTile) Transparent() bool {
	return false
}

func (it *ItemTile) Opaque() bool {
	return false
}

func (it *ItemTile) Type() TileType {
	return TileType{}
}

type EntityTile interface {
	Entity() npc.MovableEntity
	Tile
}

type BasicEntityTile struct {
	entity npc.MovableEntity

	presentation rune
	style        tcell.Style
}

func CreateBasicEntityTile(entity npc.MovableEntity, presentation rune, style tcell.Style) *BasicEntityTile {
	return &BasicEntityTile{
		entity:       entity,
		presentation: presentation,
		style:        style,
	}
}

func (bet *BasicEntityTile) Entity() npc.MovableEntity {
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

func (bet *BasicEntityTile) Opaque() bool {
	return false
}

func (bet *BasicEntityTile) Type() TileType {
	return TileType{}
}
