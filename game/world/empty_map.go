package world

import "mvvasilev/last_light/engine"

type EmptyDungeonMap struct {
	level *BasicMap
}

func (edl *EmptyDungeonMap) Size() engine.Size {
	return edl.level.Size()
}

func (edl *EmptyDungeonMap) SetTileAt(x int, y int, t Tile) Tile {
	return edl.level.SetTileAt(x, y, t)
}

func (edl *EmptyDungeonMap) TileAt(x int, y int) Tile {
	return edl.level.TileAt(x, y)
}

func (edl *EmptyDungeonMap) Tick(dt int64) {

}

func (edl *EmptyDungeonMap) Rooms() []engine.BoundingBox {
	rooms := make([]engine.BoundingBox, 1)

	rooms = append(rooms, engine.BoundingBox{
		Sized:      engine.WithSize(edl.Size()),
		Positioned: engine.WithPosition(engine.PositionAt(0, 0)),
	})

	return rooms
}

func (edl *EmptyDungeonMap) PlayerSpawnPoint() engine.Position {
	return engine.PositionAt(edl.Size().Width()/2, edl.Size().Height()/2)
}

func (edl *EmptyDungeonMap) NextLevelStaircasePosition() engine.Position {
	return engine.PositionAt(edl.Size().Width()/3, edl.Size().Height()/3)
}

func (bsp *EmptyDungeonMap) PreviousLevelStaircasePosition() engine.Position {
	return bsp.PlayerSpawnPoint()
}
