package world

import "mvvasilev/last_light/game/model"

type dungeonLevel struct {
	groundLevel Map
	entityLevel *EntityMap
	itemLevel   *Map
}

type Dungeon struct {
	player *model.Player

	levels []*dungeonLevel
}
