package world

import (
	"mvvasilev/last_light/engine"
)

type Map interface {
	Size() engine.Size
	SetTileAt(x, y int, t Tile)
	TileAt(x, y int) Tile
	Tick(dt int64)
}

type WithPlayerSpawnPoint interface {
	PlayerSpawnPoint() engine.Position
	Map
}

type WithRooms interface {
	Rooms() []engine.BoundingBox
	Map
}

type BasicMap struct {
	tiles [][]Tile
}

func CreateBasicMap(tiles [][]Tile) *BasicMap {
	bm := new(BasicMap)

	bm.tiles = tiles

	return bm
}

func (bm *BasicMap) Tick() {
}

func (bm *BasicMap) Size() engine.Size {
	return engine.SizeOf(len(bm.tiles[0]), len(bm.tiles))
}

func (bm *BasicMap) SetTileAt(x int, y int, t Tile) {
	if len(bm.tiles) <= y || len(bm.tiles[0]) <= x {
		return
	}

	if x < 0 || y < 0 {
		return
	}

	bm.tiles[y][x] = t
}

func (bm *BasicMap) TileAt(x int, y int) Tile {
	if x < 0 || y < 0 {
		return CreateStaticTile(x, y, TileTypeVoid())
	}

	if x >= bm.Size().Width() || y >= bm.Size().Height() {
		return CreateStaticTile(x, y, TileTypeVoid())
	}

	tile := bm.tiles[y][x]

	if tile == nil {
		return CreateStaticTile(x, y, TileTypeVoid())
	}

	return tile
}
