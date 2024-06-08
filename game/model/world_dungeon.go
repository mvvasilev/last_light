package model

import (
	"math/rand"
	"mvvasilev/last_light/engine"
	"slices"

	"github.com/google/uuid"
)

type DungeonType int

const (
	DungeonTypeBSP DungeonType = iota
	DungeonTypeCaverns
	DungeonTypeMine
	DungeonTypeUndercity
)

func randomDungeonType() DungeonType {
	return DungeonType(rand.Intn(4))
}

type Dungeon struct {
	levels []*DungeonLevel

	current int
}

func CreateDungeon(width, height int, depth int) *Dungeon {
	levels := make([]*DungeonLevel, 0, depth)

	for range depth {
		levels = append(levels, CreateDungeonLevel(width, height, randomDungeonType()))
	}

	return &Dungeon{
		levels:  levels,
		current: 0,
	}
}

func (d *Dungeon) CurrentLevel() *DungeonLevel {
	return d.levels[d.current]
}

func (d *Dungeon) MoveToNextLevel() (moved bool) {
	if !d.HasNextLevel() {
		return false
	}

	d.current++

	return true
}

func (d *Dungeon) MoveToPreviousLevel() (moved bool) {
	if !d.HasPreviousLevel() {
		return false
	}

	d.current--

	return true
}

func (d *Dungeon) NextLevel() *DungeonLevel {
	if !d.HasNextLevel() {
		return nil
	}

	return d.levels[d.current+1]
}

func (d *Dungeon) PreviousLevel() *DungeonLevel {
	if !d.HasPreviousLevel() {
		return nil
	}

	return d.levels[d.current-1]
}

func (d *Dungeon) HasPreviousLevel() bool {
	return d.current-1 >= 0
}

func (d *Dungeon) HasNextLevel() bool {
	return d.current+1 < len(d.levels)
}

type DungeonLevel struct {
	ground             Map
	entitiesByPosition map[engine.Position][]Entity
	entities           map[uuid.UUID]Entity
}

func CreateDungeonLevel(width, height int, dungeonType DungeonType) (dLevel *DungeonLevel) {

	genTable := CreateLootTable()

	genTable.Add(1, func() Item {
		return Item_HealthPotion()
	})

	itemPool := []Item{
		Item_Bow(),
		Item_Longsword(),
		Item_Club(),
		Item_Dagger(),
		Item_Handaxe(),
		Item_Javelin(),
		Item_LightHammer(),
		Item_Mace(),
		Item_Quarterstaff(),
		Item_Sickle(),
		Item_Spear(),
	}

	genTable.Add(1, func() Item {
		item := itemPool[rand.Intn(len(itemPool))]

		rarities := []ItemRarity{
			ItemRarity_Common,
			ItemRarity_Uncommon,
			ItemRarity_Rare,
			ItemRarity_Epic,
			ItemRarity_Legendary,
		}

		return GenerateItemOfTypeAndRarity(item, rarities[rand.Intn(len(rarities))])
	})

	var groundLevel Map

	switch dungeonType {
	case DungeonTypeBSP:
		groundLevel = CreateBSPDungeonMap(width, height, 4)
	default:
		groundLevel = CreateBSPDungeonMap(width, height, 4)
	}

	dLevel = &DungeonLevel{
		ground:             groundLevel,
		entities:           map[uuid.UUID]Entity{},
		entitiesByPosition: map[engine.Position][]Entity{},
	}

	if groundLevel.Rooms() == nil {
		return dLevel
	}

	forbiddenItemPositions := make([]engine.Position, 0)

	if groundLevel.NextLevelStaircase() != nil {
		forbiddenItemPositions = append(forbiddenItemPositions, groundLevel.NextLevelStaircase().Position)
	}

	if groundLevel.PreviousLevelStaircase() != nil {
		forbiddenItemPositions = append(forbiddenItemPositions, groundLevel.PreviousLevelStaircase().Position)
	}

	if groundLevel.PlayerSpawnPoint() != nil {
		forbiddenItemPositions = append(forbiddenItemPositions, groundLevel.PreviousLevelStaircase().Position)
	}

	items := SpawnItems(groundLevel.Rooms().Rooms, 0.01, genTable, forbiddenItemPositions)

	for pos, it := range items {
		tile := Map_TileAt(groundLevel, pos.X(), pos.Y())

		if !tile.Passable() {
			continue
		}

		Map_SetTileAt(
			groundLevel,
			pos.X(),
			pos.Y(),
			CreateTileFromPrototype(tile, Tile_WithItem(it)),
		)
	}

	return dLevel
}

func SpawnItems(spawnableAreas []engine.BoundingBox, maxItemRatio float32, genTable *LootTable, forbiddenPositions []engine.Position) map[engine.Position]Item {
	rooms := spawnableAreas

	itemLocations := make(map[engine.Position]Item, 0)

	for _, r := range rooms {
		maxItems := int(maxItemRatio * float32(r.Size().Area()))

		if maxItems < 1 {
			continue
		}

		numItems := rand.Intn(maxItems)

		for range numItems {
			item := genTable.Generate()

			if item == nil {
				continue
			}

			pos := engine.PositionAt(
				engine.RandInt(r.Position().X()+1, r.Position().X()+r.Size().Width()-1),
				engine.RandInt(r.Position().Y()+1, r.Position().Y()+r.Size().Height()-1),
			)

			if slices.Contains(forbiddenPositions, pos) {
				continue
			}

			itemLocations[pos] = item
		}
	}

	return itemLocations
}

func (d *DungeonLevel) Ground() Map {
	return d.ground
}

func (d *DungeonLevel) DropEntity(uuid uuid.UUID) {
	ent := d.entities[uuid]

	if ent != nil {
		delete(d.entitiesByPosition, ent.Positioned().Position)
	}

	delete(d.entities, uuid)
}

func (d *DungeonLevel) AddEntity(entity Entity) {
	d.entities[entity.UniqueId()] = entity

	if entity.Positioned() != nil {
		if d.entitiesByPosition[entity.Positioned().Position] == nil {
			d.entitiesByPosition[entity.Positioned().Position] = []Entity{entity}
		} else {
			d.entitiesByPosition[entity.Positioned().Position] = append(d.entitiesByPosition[entity.Positioned().Position], entity)
		}
	}
}

func (d *DungeonLevel) MoveEntityTo(uuid uuid.UUID, x, y int) {
	ent := d.entities[uuid]

	if ent == nil || ent.Positioned() == nil {
		return
	}

	d.RemoveEntityAt(ent.Positioned().Position.XY())

	ent.Positioned().Position = engine.PositionAt(x, y)

	if d.entitiesByPosition[ent.Positioned().Position] == nil {
		d.entitiesByPosition[ent.Positioned().Position] = []Entity{ent}
	} else {
		d.entitiesByPosition[ent.Positioned().Position] = append(d.entitiesByPosition[ent.Positioned().Position], ent)
	}
}

func (d *DungeonLevel) RemoveEntityAt(x, y int) {
	delete(d.entitiesByPosition, engine.PositionAt(x, y))
}

func (d *DungeonLevel) RemoveItemAt(x, y int) (item Item) {
	if !Map_IsInBounds(d.ground, x, y) {
		return nil
	}

	tile := Map_TileAt(d.ground, x, y)

	if tile.Item() == nil {
		return nil
	}

	item = tile.Item().Item

	tile.RemoveItem()

	return
}

func (d *DungeonLevel) SetItemAt(x, y int, it Item) (success bool) {
	if !d.TileAt(x, y).Passable() {
		return false
	}

	tile := d.TileAt(x, y)

	tile.WithItem(it)

	return true
}

func (d *DungeonLevel) TileAt(x, y int) Tile {
	entity := d.entitiesByPosition[engine.PositionAt(x, y)]
	tile := Map_TileAt(d.ground, x, y)

	if entity != nil {
		return CreateTileFromPrototype(tile, Tile_WithEntities(entity))
	}

	return tile
}

func (d *DungeonLevel) IsTilePassable(x, y int) bool {
	if !Map_IsInBounds(d.ground, x, y) {
		return false
	}

	tile := d.TileAt(x, y)

	if tile.Entities() != nil {
		return false
	}

	return tile.Passable()
}

func (d *DungeonLevel) EntitiesAt(x, y int) (e []Entity) {
	return d.entitiesByPosition[engine.PositionAt(x, y)]
}

func (d *DungeonLevel) IsGroundTileOpaque(x, y int) bool {
	if !Map_IsInBounds(d.ground, x, y) {
		return false
	}

	return Map_TileAt(d.ground, x, y).Opaque()
}
