package world

import (
	"mvvasilev/last_light/engine"

	"github.com/gdamore/tcell/v2"
)

type Map interface {
	Size() engine.Size
	SetTileAt(x, y int, t Tile) Tile
	TileAt(x, y int) Tile
	IsInBounds(x, y int) bool
	ExploredTileAt(x, y int) Tile
	MarkExplored(x, y int)
	Tick(dt int64)
}

type WithPlayerSpawnPoint interface {
	PlayerSpawnPoint() engine.Position
}

type WithRooms interface {
	Rooms() []engine.BoundingBox
}

type WithNextLevelStaircasePosition interface {
	NextLevelStaircasePosition() engine.Position
}

type WithPreviousLevelStaircasePosition interface {
	PreviousLevelStaircasePosition() engine.Position
}

type BasicMap struct {
	tiles         [][]Tile
	exploredTiles map[engine.Position]Tile

	exploredStyle tcell.Style
}

func CreateBasicMap(tiles [][]Tile, exploredStyle tcell.Style) *BasicMap {
	bm := new(BasicMap)

	bm.tiles = tiles
	bm.exploredTiles = make(map[engine.Position]Tile, 0)
	bm.exploredStyle = exploredStyle

	return bm
}

func (bm *BasicMap) Tick(dt int64) {
}

func (bm *BasicMap) Size() engine.Size {
	return engine.SizeOf(len(bm.tiles[0]), len(bm.tiles))
}

func (bm *BasicMap) SetTileAt(x int, y int, t Tile) Tile {
	if !bm.IsInBounds(x, y) {
		return CreateStaticTile(x, y, TileTypeVoid())
	}

	bm.tiles[y][x] = t

	return bm.tiles[y][x]
}

func (bm *BasicMap) TileAt(x int, y int) Tile {
	if !bm.IsInBounds(x, y) {
		return CreateStaticTile(x, y, TileTypeVoid())
	}

	tile := bm.tiles[y][x]

	if tile == nil {
		return CreateStaticTile(x, y, TileTypeVoid())
	}

	return tile
}

func (bm *BasicMap) IsInBounds(x, y int) bool {
	if x < 0 || y < 0 {
		return false
	}

	if x >= bm.Size().Width() || y >= bm.Size().Height() {
		return false
	}

	return true
}

func (bm *BasicMap) ExploredTileAt(x, y int) Tile {
	return bm.exploredTiles[engine.PositionAt(x, y)]
}

func (bm *BasicMap) MarkExplored(x, y int) {
	if !bm.IsInBounds(x, y) {
		return
	}

	tile := bm.TileAt(x, y)

	bm.exploredTiles[engine.PositionAt(x, y)] = CreateStaticTileWithStyleOverride(tile.Position().X(), tile.Position().Y(), tile.Type(), bm.exploredStyle)
}
