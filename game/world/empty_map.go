package world

import "mvvasilev/last_light/util"

type EmptyDungeonMap struct {
	level *BasicMap
}

func (edl *EmptyDungeonMap) Size() util.Size {
	return edl.level.Size()
}

func (edl *EmptyDungeonMap) SetTileAt(x int, y int, t Tile) {
	edl.level.SetTileAt(x, y, t)
}

func (edl *EmptyDungeonMap) TileAt(x int, y int) Tile {
	return edl.level.TileAt(x, y)
}

func (edl *EmptyDungeonMap) Tick(dt int64) {

}
