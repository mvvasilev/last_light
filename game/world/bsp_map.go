package world

import (
	"mvvasilev/last_light/engine"
)

type BSPDungeonMap struct {
	level *BasicMap

	playerSpawnPoint   engine.Position
	nextLevelStaircase engine.Position
	rooms              []engine.BoundingBox
}

func (bsp *BSPDungeonMap) PlayerSpawnPoint() engine.Position {
	return bsp.playerSpawnPoint
}

func (bsp *BSPDungeonMap) NextLevelStaircasePosition() engine.Position {
	return bsp.nextLevelStaircase
}

func (bsp *BSPDungeonMap) Size() engine.Size {
	return bsp.level.Size()
}

func (bsp *BSPDungeonMap) SetTileAt(x int, y int, t Tile) Tile {
	return bsp.level.SetTileAt(x, y, t)
}

func (bsp *BSPDungeonMap) TileAt(x int, y int) Tile {
	return bsp.level.TileAt(x, y)
}

func (bsp *BSPDungeonMap) Tick(dt int64) {
}

func (bsp *BSPDungeonMap) Rooms() []engine.BoundingBox {
	return bsp.rooms
}

func (bsp *BSPDungeonMap) PreviousLevelStaircasePosition() engine.Position {
	return bsp.playerSpawnPoint
}
