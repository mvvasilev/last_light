package world

import (
	"mvvasilev/last_light/util"
)

type BSPDungeonMap struct {
	level *BasicMap

	playerSpawnPoint util.Position
	rooms            []util.Room
}

func (bsp *BSPDungeonMap) PlayerSpawnPoint() util.Position {
	return bsp.playerSpawnPoint
}

func (bsp *BSPDungeonMap) Size() util.Size {
	return bsp.level.Size()
}

func (bsp *BSPDungeonMap) SetTileAt(x int, y int, t Tile) {
	bsp.level.SetTileAt(x, y, t)
}

func (bsp *BSPDungeonMap) TileAt(x int, y int) Tile {
	return bsp.level.TileAt(x, y)
}

func (bsp *BSPDungeonMap) Tick(dt int64) {
}

func (bsp *BSPDungeonMap) Rooms() []util.Room {
	return bsp.rooms
}
