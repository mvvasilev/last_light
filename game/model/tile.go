package model

import "mvvasilev/last_light/util"

type Material uint

const (
	MaterialGround Material = iota
	MaterialRock
	MaterialGrass
	MaterialVoid
)

type TileType struct {
	Material     Material
	Passable     bool
	Presentation rune
}

func Ground() TileType {
	return TileType{
		Material:     MaterialGround,
		Passable:     true,
		Presentation: '.',
	}
}

func Rock() TileType {
	return TileType{
		Material:     MaterialRock,
		Passable:     false,
		Presentation: '█',
	}
}

func Grass() TileType {
	return TileType{
		Material:     MaterialGrass,
		Passable:     true,
		Presentation: ',',
	}
}

func Void() TileType {
	return TileType{
		Material:     MaterialVoid,
		Passable:     false,
		Presentation: ' ',
	}
}

type Tile interface {
	Position() util.Position
	Presentation() rune
	Passable() bool
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

func (st *StaticTile) Presentation() rune {
	return st.t.Presentation
}

func (st *StaticTile) Passable() bool {
	return st.t.Passable
}

func (st *StaticTile) Type() TileType {
	return st.t
}
