package model

// import (
// 	"maps"
// 	"mvvasilev/last_light/engine"
// 	"mvvasilev/last_light/game/npc"

// 	"github.com/google/uuid"
// )

// type EntityMap struct {
// 	entities map[int]EntityTile

// 	engine.Sized
// }

// func CreateEntityMap(width, height int) *EntityMap {
// 	return &EntityMap{
// 		entities: make(map[int]EntityTile, 0),
// 		Sized:    engine.WithSize(engine.SizeOf(width, height)),
// 	}
// }

// func (em *EntityMap) SetTileAt(x int, y int, t Tile) Tile {
// 	return nil
// 	// if !em.FitsWithin(x, y) {
// 	// 	return
// 	// }

// 	// index := em.Size().AsArrayIndex(x, y)

// 	// TODO? May not be necessary
// }

// func (em *EntityMap) FindEntityByUuid(uuid uuid.UUID) (key int, entity EntityTile) {
// 	for i, e := range em.entities {
// 		if e.Entity().UniqueId() == uuid {
// 			return i, e
// 		}
// 	}

// 	return -1, nil
// }

// func (em *EntityMap) AddEntity(entity Entity_V2) {
// 	if entity.Positioned() == nil {
// 		return
// 	}

// 	if !em.FitsWithin(entity.Positioned().Position.XY()) {
// 		return
// 	}

// 	key := em.Size().AsArrayIndex(entity.Positioned().Position.XY())
// 	et := CreateBasicEntityTile(entity)

// 	em.entities[key] = et
// }

// func (em *EntityMap) DropEntity(uuid uuid.UUID) {
// 	maps.DeleteFunc(em.entities, func(i int, et EntityTile) bool {
// 		return et.Entity().UniqueId() == uuid
// 	})
// }

// func (em *EntityMap) MoveEntity(uuid uuid.UUID, dx, dy int) {
// 	oldKey, e := em.FindEntityByUuid(uuid)

// 	if e == nil {
// 		return
// 	}

// 	if !em.FitsWithin(e.Entity().Positioned().Position.WithOffset(dx, dy).XY()) {
// 		return
// 	}

// 	delete(em.entities, oldKey)

// 	newPos := e.Entity().Position().WithOffset(dx, dy)
// 	e.Entity().MoveTo(newPos)

// 	newKey := em.Size().AsArrayIndex(e.Entity().Position().XY())

// 	em.entities[newKey] = e
// }

// func (em *EntityMap) MoveEntityTo(uuid uuid.UUID, x, y int) {
// 	oldKey, e := em.FindEntityByUuid(uuid)

// 	if e == nil {
// 		return
// 	}

// 	if !em.FitsWithin(x, y) {
// 		return
// 	}

// 	delete(em.entities, oldKey)

// 	e.Entity().MoveTo(engine.PositionAt(x, y))

// 	newKey := em.Size().AsArrayIndex(e.Entity().Position().XY())

// 	em.entities[newKey] = e
// }

// func (em *EntityMap) TileAt(x int, y int) Tile {
// 	if !em.FitsWithin(x, y) {
// 		return CreateStaticTile(x, y, TileTypeVoid())
// 	}

// 	key := em.Size().AsArrayIndex(x, y)

// 	return em.entities[key]
// }

// func (em *EntityMap) EntityAt(x, y int) (ent npc.MovableEntity) {
// 	tile := em.TileAt(x, y)

// 	if tile == nil {
// 		return nil
// 	}

// 	return tile.(EntityTile).Entity()
// }

// func (em *EntityMap) IsInBounds(x, y int) bool {
// 	return em.FitsWithin(x, y)
// }

// func (em *EntityMap) MarkExplored(x, y int) {

// }

// func (em *EntityMap) ExploredTileAt(x, y int) Tile {
// 	return CreateStaticTile(x, y, TileTypeVoid())
// }

// func (em *EntityMap) Tick(dt int64) {
// }
