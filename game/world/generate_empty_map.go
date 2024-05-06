package world

func CreateEmptyDungeonLevel(width, height int) *EmptyDungeonMap {
	m := new(EmptyDungeonMap)

	tiles := make([][]Tile, height)

	for h := range height {
		tiles[h] = make([]Tile, width)
	}

	m.level = CreateBasicMap(tiles)

	return m
}
