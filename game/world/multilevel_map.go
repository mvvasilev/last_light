package world

import "mvvasilev/last_light/util"

type MultilevelMap struct {
	layers []Map
}

func CreateMultilevelMap(maps ...Map) *MultilevelMap {
	m := new(MultilevelMap)

	m.layers = maps

	return m
}

func (mm *MultilevelMap) Size() util.Size {
	if len(mm.layers) == 0 {
		return util.SizeOf(0, 0)
	}

	return mm.layers[0].Size()
}

func (mm *MultilevelMap) SetTileAt(x, y int, t Tile) {
	mm.layers[0].SetTileAt(x, y, t)
}

func (mm *MultilevelMap) UnsetTileAtHeight(x, y, height int) {
	if len(mm.layers) < height {
		return
	}

	mm.layers[height].SetTileAt(x, y, nil)
}

func (mm *MultilevelMap) SetTileAtHeight(x, y, height int, t Tile) {
	if len(mm.layers) < height {
		return
	}

	mm.layers[height].SetTileAt(x, y, t)
}

func (mm *MultilevelMap) CollectTilesAt(x, y int, filter func(t Tile) bool) []Tile {
	tiles := make([]Tile, len(mm.layers))

	if x < 0 || y < 0 {
		return tiles
	}

	if x >= mm.Size().Width() || y >= mm.Size().Height() {
		return tiles
	}

	for i := len(mm.layers) - 1; i >= 0; i-- {
		tile := mm.layers[i].TileAt(x, y)

		if tile != nil && !tile.Transparent() && filter(tile) {
			tiles = append(tiles, tile)
		}

	}

	return tiles
}

func (mm *MultilevelMap) TileAt(x int, y int) Tile {
	if x < 0 || y < 0 {
		return CreateStaticTile(x, y, TileTypeVoid())
	}

	if x >= mm.Size().Width() || y >= mm.Size().Height() {
		return CreateStaticTile(x, y, TileTypeVoid())
	}

	for i := len(mm.layers) - 1; i >= 0; i-- {
		tile := mm.layers[i].TileAt(x, y)

		if tile != nil && !tile.Transparent() {
			return tile
		}

	}

	return CreateStaticTile(x, y, TileTypeVoid())
}

func (mm *MultilevelMap) TileAtHeight(x, y, height int) Tile {
	if x < 0 || y < 0 {
		return CreateStaticTile(x, y, TileTypeVoid())
	}

	if x >= mm.Size().Width() || y >= mm.Size().Height() {
		return CreateStaticTile(x, y, TileTypeVoid())
	}

	if height > len(mm.layers)-1 {
		return CreateStaticTile(x, y, TileTypeVoid())
	}

	return mm.layers[height].TileAt(x, y)
}

func (mm *MultilevelMap) Tick(dt int64) {
	for _, l := range mm.layers {
		l.Tick(dt)
	}
}
