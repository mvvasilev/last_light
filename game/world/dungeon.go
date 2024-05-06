package world

type dungeonLevel struct {
	groundLevel *Map
	entityLevel *EntityMap
	itemLevel   *Map
}

type Dungeon struct {
	levels []*dungeonLevel
}
