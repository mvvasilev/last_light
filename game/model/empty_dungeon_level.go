package model

import "mvvasilev/last_light/util"

type EmptyDungeonLevel struct {
	level *BasicMap
}

func CreateEmptyDungeonLevel(width, height int) *EmptyDungeonLevel {
	m := new(EmptyDungeonLevel)

	tiles := make([][]Tile, height)

	for h := range height {
		tiles[h] = make([]Tile, width)
	}

	m.level = CreateBasicMap(tiles)

	return m
}

func (edl *EmptyDungeonLevel) Size() util.Size {
	return edl.level.Size()
}

func (edl *EmptyDungeonLevel) SetTileAt(x int, y int, t Tile) {
	edl.level.SetTileAt(x, y, t)
}

func (edl *EmptyDungeonLevel) TileAt(x int, y int) Tile {
	return edl.level.TileAt(x, y)
}

func (edl *EmptyDungeonLevel) Tick() {

}
