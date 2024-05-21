package world

import (
	"math/rand"
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/item"
	"mvvasilev/last_light/game/model"
	"mvvasilev/last_light/game/rpg"
	"slices"

	"github.com/gdamore/tcell/v2"
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
	groundLevel interface {
		Map
		WithPlayerSpawnPoint
		WithNextLevelStaircasePosition
		WithPreviousLevelStaircasePosition
	}
	entityLevel *EntityMap
	itemLevel   Map

	multilevel Map
}

func CreateDungeonLevel(width, height int, dungeonType DungeonType) *DungeonLevel {

	genTable := rpg.CreateLootTable()

	genTable.Add(10, func() item.Item {
		return item.CreateBasicItem(item.ItemTypeFish(), 1)
	})

	genTable.Add(1, func() item.Item {
		itemTypes := []rpg.RPGItemType{
			rpg.ItemTypeBow(),
			rpg.ItemTypeLongsword(),
			rpg.ItemTypeClub(),
			rpg.ItemTypeDagger(),
			rpg.ItemTypeHandaxe(),
			rpg.ItemTypeJavelin(),
			rpg.ItemTypeLightHammer(),
			rpg.ItemTypeMace(),
			rpg.ItemTypeQuarterstaff(),
			rpg.ItemTypeSickle(),
			rpg.ItemTypeSpear(),
		}

		itemType := itemTypes[rand.Intn(len(itemTypes))]

		rarities := []rpg.ItemRarity{
			rpg.ItemRarity_Common,
			rpg.ItemRarity_Uncommon,
			rpg.ItemRarity_Rare,
			rpg.ItemRarity_Epic,
			rpg.ItemRarity_Legendary,
		}

		return rpg.GenerateItemOfTypeAndRarity(itemType, rarities[rand.Intn(len(rarities))])
	})

	var groundLevel interface {
		Map
		WithRooms
		WithPlayerSpawnPoint
		WithNextLevelStaircasePosition
		WithPreviousLevelStaircasePosition
	}

	switch dungeonType {
	case DungeonTypeBSP:
		groundLevel = CreateBSPDungeonMap(width, height, 4)
	default:
		groundLevel = CreateBSPDungeonMap(width, height, 4)
	}

	items := SpawnItems(groundLevel.Rooms(), 0.02, genTable, []engine.Position{
		groundLevel.NextLevelStaircasePosition(),
		groundLevel.PlayerSpawnPoint(),
		groundLevel.PreviousLevelStaircasePosition(),
	})

	itemLevel := CreateEmptyDungeonLevel(width, height)

	for _, it := range items {
		if !groundLevel.TileAt(it.Position().XY()).Passable() {
			continue
		}

		itemLevel.SetTileAt(it.Position().X(), it.Position().Y(), it)
	}

	d := &DungeonLevel{
		groundLevel: groundLevel,
		entityLevel: CreateEntityMap(width, height),
		itemLevel:   itemLevel,
	}

	d.multilevel = CreateMultilevelMap(
		d.groundLevel,
		d.itemLevel,
		d.entityLevel,
	)

	return d
}

func SpawnItems(spawnableAreas []engine.BoundingBox, maxItemRatio float32, genTable *rpg.LootTable, forbiddenPositions []engine.Position) []Tile {
	rooms := spawnableAreas

	itemTiles := make([]Tile, 0, 10)

	for _, r := range rooms {
		maxItems := int(maxItemRatio * float32(r.Size().Area()))

		if maxItems < 1 {
			continue
		}

		numItems := rand.Intn(maxItems)

		for range numItems {
			itemType := genTable.Generate()

			if itemType == nil {
				continue
			}

			pos := engine.PositionAt(
				engine.RandInt(r.Position().X()+1, r.Position().X()+r.Size().Width()-1),
				engine.RandInt(r.Position().Y()+1, r.Position().Y()+r.Size().Height()-1),
			)

			if slices.Contains(forbiddenPositions, pos) {
				continue
			}

			itemTiles = append(itemTiles, CreateItemTile(pos, itemType))
		}
	}

	return itemTiles
}

func (d *DungeonLevel) PlayerSpawnPoint() engine.Position {
	return d.groundLevel.PlayerSpawnPoint()
}

func (d *DungeonLevel) NextLevelStaircase() engine.Position {
	return d.groundLevel.NextLevelStaircasePosition()
}

func (d *DungeonLevel) PreviousLevelStaircase() engine.Position {
	return d.groundLevel.PreviousLevelStaircasePosition()
}

func (d *DungeonLevel) DropEntity(uuid uuid.UUID) {
	d.entityLevel.DropEntity(uuid)
}

func (d *DungeonLevel) AddEntity(entity model.MovableEntity, presentation rune, style tcell.Style) {
	d.entityLevel.AddEntity(entity, presentation, style)
}

func (d *DungeonLevel) MoveEntity(uuid uuid.UUID, dx, dy int) {
	d.entityLevel.MoveEntity(uuid, dx, dy)
}

func (d *DungeonLevel) MoveEntityTo(uuid uuid.UUID, x, y int) {
	d.entityLevel.MoveEntityTo(uuid, x, y)
}

func (d *DungeonLevel) RemoveItemAt(x, y int) item.Item {
	if !d.groundLevel.Size().Contains(x, y) {
		return nil
	}

	tile := d.itemLevel.TileAt(x, y)
	itemTile, ok := tile.(*ItemTile)

	if !ok {
		return nil
	}

	d.itemLevel.SetTileAt(x, y, nil)

	return itemTile.Item()
}

func (d *DungeonLevel) SetItemAt(x, y int, it item.Item) (success bool) {
	if !d.TileAt(x, y).Passable() {
		return false
	}

	d.itemLevel.SetTileAt(x, y, CreateItemTile(engine.PositionAt(x, y), it))

	return true
}

func (d *DungeonLevel) TileAt(x, y int) Tile {
	return d.multilevel.TileAt(x, y)
}

func (d *DungeonLevel) IsTilePassable(x, y int) bool {
	if !d.groundLevel.Size().Contains(x, y) {
		return false
	}

	return d.TileAt(x, y).Passable()
}

func (d *DungeonLevel) Flatten() Map {
	return d.multilevel
}
