package model

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

func (mm *MultilevelMap) TileAt(x int, y int) Tile {
	for i := len(mm.layers) - 1; i >= 0; i-- {
		tile := mm.layers[i].TileAt(x, y)

		if tile != nil {
			return tile
		}
	}

	return nil
}

func (mm *MultilevelMap) Tick() {
	for _, l := range mm.layers {
		l.Tick()
	}
}
