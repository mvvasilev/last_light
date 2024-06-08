package model

import (
	"slices"

	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
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

type Tile_ItemComponent struct {
	Item Item
}

type Tile_EntityComponent struct {
	Entities []Entity
}

type Tile interface {
	DefaultPresentation() (rune, tcell.Style)
	Material() Material
	Passable() bool
	Opaque() bool
	Transparent() bool

	Item() *Tile_ItemComponent
	RemoveItem()
	WithItem(item Item)

	Entities() *Tile_EntityComponent
	RemoveEntity(uuid uuid.UUID)
	AddEntity(entity Entity)
}

type BaseTile struct {
	defaultSymbol rune
	defaultStyle  tcell.Style

	material                      Material
	passable, opaque, transparent bool

	item     *Tile_ItemComponent
	entities *Tile_EntityComponent
}

func CreateTileFromPrototype(prototype Tile, components ...func(*BaseTile)) Tile {
	defaultSymbol, defaultStyle := prototype.DefaultPresentation()

	return CreateTile(
		defaultSymbol,
		defaultStyle,
		prototype.Material(),
		prototype.Passable(),
		prototype.Opaque(),
		prototype.Transparent(),
		components...,
	)
}

func CreateTile(defaultSymbol rune, defaultStyle tcell.Style, material Material, passable, opaque, transparent bool, components ...func(*BaseTile)) Tile {
	t := &BaseTile{
		defaultSymbol: defaultSymbol,
		defaultStyle:  defaultStyle,
		material:      material,
		passable:      passable,
		opaque:        opaque,
		transparent:   transparent,
	}

	for _, c := range components {
		c(t)
	}

	return t
}

func (t *BaseTile) DefaultPresentation() (rune, tcell.Style) {
	return t.defaultSymbol, t.defaultStyle
}

func (t *BaseTile) Material() Material {
	return t.material
}

func (t *BaseTile) Passable() bool {
	return t.passable
}

func (t *BaseTile) Opaque() bool {
	return t.opaque
}

func (t *BaseTile) Transparent() bool {
	return t.transparent
}

func (t *BaseTile) Item() *Tile_ItemComponent {
	return t.item
}

func (t *BaseTile) RemoveItem() {
	t.item = nil
}

func (t *BaseTile) WithItem(item Item) {
	t.item = &Tile_ItemComponent{
		Item: item,
	}
}

func (t *BaseTile) Entities() *Tile_EntityComponent {
	return t.entities
}

func (t *BaseTile) RemoveEntity(uuid uuid.UUID) {
	if t.entities == nil {
		return
	}

	t.entities.Entities = slices.DeleteFunc(t.entities.Entities, func(e Entity) bool { return e.UniqueId() == uuid })
}

func (t *BaseTile) AddEntity(entity Entity) {
	if t.entities == nil {
		t.entities = &Tile_EntityComponent{
			Entities: []Entity{
				entity,
			},
		}

		return
	}

	t.entities.Entities = append(t.entities.Entities, entity)
}

func Tile_WithEntity(entity Entity) func(*BaseTile) {
	return func(bt *BaseTile) {
		bt.entities = &Tile_EntityComponent{
			Entities: []Entity{
				entity,
			},
		}
	}
}

func Tile_WithEntities(entities []Entity) func(*BaseTile) {
	return func(bt *BaseTile) {
		bt.entities = &Tile_EntityComponent{
			Entities: entities,
		}
	}
}

func Tile_WithItem(item Item) func(*BaseTile) {
	return func(bt *BaseTile) {
		bt.item = &Tile_ItemComponent{
			Item: item,
		}
	}
}

func Tile_Void() Tile {
	return CreateTile(
		' ',
		tcell.StyleDefault,
		MaterialVoid,
		false, true, true,
	)
}

func Tile_Ground() Tile {
	return CreateTile(
		'.',
		tcell.StyleDefault,
		MaterialGround,
		true, false, false,
	)
}

func Tile_Rock() Tile {
	return CreateTile(
		'█',
		tcell.StyleDefault,
		MaterialRock,
		false, true, false,
	)
}

func Tile_Wall() Tile {
	return CreateTile(
		'#',
		tcell.StyleDefault.Background(tcell.ColorGray),
		MaterialWall,
		false, true, false,
	)
}

// func TileTypeClosedDoor() TileType {
// 	return TileType{
// 		Material:     MaterialClosedDoor,
// 		Passable:     false,
// 		Transparent:  false,
// 		Presentation: '[',
// 		Opaque:       true,
// 		Style:        tcell.StyleDefault.Foreground(tcell.ColorLightSteelBlue).Background(tcell.ColorSaddleBrown),
// 	}
// }

// func TileTypeOpenDoor() TileType {
// 	return TileType{
// 		Material:     MaterialClosedDoor,
// 		Passable:     false,
// 		Transparent:  false,
// 		Presentation: '_',
// 		Opaque:       false,
// 		Style:        tcell.StyleDefault.Foreground(tcell.ColorLightSteelBlue),
// 	}
// }

func Tile_StaircaseDown() Tile {
	return CreateTile(
		'≡',
		tcell.StyleDefault.Foreground(tcell.ColorDarkSlateGray).Attributes(tcell.AttrBold),
		MaterialStaircaseDown,
		true, false, false,
	)
}

func Tile_StaircaseUp() Tile {
	return CreateTile(
		'^',
		tcell.StyleDefault.Foreground(tcell.ColorDarkSlateGray).Attributes(tcell.AttrBold),
		MaterialStaircaseDown,
		true, false, false,
	)
}

// type Tile interface {
// 	Position() engine.Position
// 	Presentation() (rune, tcell.Style)
// 	Passable() bool
// 	Transparent() bool
// 	Opaque() bool
// 	Type() TileType
// }

// type StaticTile struct {
// 	position engine.Position
// 	t        TileType

// 	style tcell.Style
// }

// func CreateStaticTile(x, y int, t TileType) Tile {
// 	st := new(StaticTile)

// 	st.position = engine.PositionAt(x, y)
// 	st.t = t
// 	st.style = t.Style

// 	return st
// }

// func CreateStaticTileWithStyleOverride(x, y int, t TileType, style tcell.Style) Tile {
// 	return &StaticTile{
// 		position: engine.PositionAt(x, y),
// 		t:        t,
// 		style:    style,
// 	}
// }

// func (st *StaticTile) Position() engine.Position {
// 	return st.position
// }

// func (st *StaticTile) Presentation() (rune, tcell.Style) {
// 	return st.t.Presentation, st.style
// }

// func (st *StaticTile) Passable() bool {
// 	return st.t.Passable
// }

// func (st *StaticTile) Transparent() bool {
// 	return st.t.Transparent
// }

// func (st *StaticTile) Opaque() bool {
// 	return st.t.Opaque
// }

// func (st *StaticTile) Type() TileType {
// 	return st.t
// }

// type ItemTile struct {
// 	position engine.Position
// 	item     item.Item
// }

// func CreateItemTile(position engine.Position, item item.Item) *ItemTile {
// 	it := new(ItemTile)

// 	it.position = position
// 	it.item = item

// 	return it
// }

// func (it *ItemTile) Item() item.Item {
// 	return it.item
// }

// func (it *ItemTile) Position() engine.Position {
// 	return it.position
// }

// func (it *ItemTile) Presentation() (rune, tcell.Style) {
// 	return it.item.Type().TileIcon(), it.item.Type().Style()
// }

// func (it *ItemTile) Passable() bool {
// 	return true
// }

// func (it *ItemTile) Transparent() bool {
// 	return false
// }

// func (it *ItemTile) Opaque() bool {
// 	return false
// }

// func (it *ItemTile) Type() TileType {
// 	return TileType{}
// }

// type EntityTile interface {
// 	Entity() npc.MovableEntity
// 	Tile
// }

// type BasicEntityTile struct {
// 	entity npc.MovableEntity
// }

// func CreateBasicEntityTile(entity npc.MovableEntity) *BasicEntityTile {
// 	return &BasicEntityTile{
// 		entity: entity,
// 	}
// }

// func (bet *BasicEntityTile) Entity() npc.MovableEntity {
// 	return bet.entity
// }

// func (bet *BasicEntityTile) Position() engine.Position {
// 	return bet.entity.Position()
// }

// func (bet *BasicEntityTile) Presentation() (rune, tcell.Style) {
// 	return bet.entity.Presentation()
// }

// func (bet *BasicEntityTile) Passable() bool {
// 	return false
// }

// func (bet *BasicEntityTile) Transparent() bool {
// 	return false
// }

// func (bet *BasicEntityTile) Opaque() bool {
// 	return false
// }

// func (bet *BasicEntityTile) Type() TileType {
// 	return TileType{}
// }
