package model

import "mvvasilev/last_light/util"

type EmptyDungeonLevel struct {
	tiles [][]Tile
}

func CreateEmptyDungeonLevel(width, height int) *EmptyDungeonLevel {
	m := new(EmptyDungeonLevel)

	m.tiles = make([][]Tile, height)

	for h := range height {
		m.tiles[h] = make([]Tile, width)
	}

	return m
}

func (edl *EmptyDungeonLevel) Size() util.Size {
	return util.SizeOf(len(edl.tiles[0]), len(edl.tiles))
}

func (edl *EmptyDungeonLevel) SetTileAt(x int, y int, t Tile) {
	if len(edl.tiles) <= y || len(edl.tiles[0]) <= x {
		return
	}

	edl.tiles[y][x] = t
}

func (edl *EmptyDungeonLevel) TileAt(x int, y int) Tile {
	if y < 0 || y >= len(edl.tiles) {
		return nil
	}

	if x < 0 || x >= len(edl.tiles[y]) {
		return nil
	}

	return edl.tiles[y][x]
}

func (edl *EmptyDungeonLevel) Tick() {

}
