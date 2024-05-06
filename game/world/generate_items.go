package world

import (
	"math/rand"
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/model"
)

func SpawnItems(spawnableAreas []engine.BoundingBox, maxItemRatio float32, genTable map[float32]*model.ItemType) []Tile {
	rooms := spawnableAreas

	itemTiles := make([]Tile, 0, 10)

	for _, r := range rooms {
		maxItems := int(maxItemRatio * float32(r.Size().Area()))

		if maxItems < 1 {
			continue
		}

		numItems := rand.Intn(maxItems)

		for range numItems {
			itemType := GenerateItemType(genTable)

			if itemType == nil {
				continue
			}

			pos := engine.PositionAt(
				engine.RandInt(r.Position().X()+1, r.Position().X()+r.Size().Width()-1),
				engine.RandInt(r.Position().Y()+1, r.Position().Y()+r.Size().Height()-1),
			)

			itemTiles = append(itemTiles, CreateItemTile(pos, itemType, 1))
		}
	}

	return itemTiles
}

func GenerateItemType(genTable map[float32]*model.ItemType) *model.ItemType {
	num := rand.Float32()

	for k, v := range genTable {
		if num > k {
			return v
		}
	}

	return nil
}
