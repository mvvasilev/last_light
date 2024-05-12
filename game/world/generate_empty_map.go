package world

func CreateEmptyDungeonLevel(width, height int) *EmptyDungeonMap {
	m := new(EmptyDungeonMap)

	tiles := make([][]Tile, height)

	for h := range height {
		tiles[h] = make([]Tile, width)
	}

	m.level = CreateBasicMap(tiles)

	//m.level.SetTileAt(width/2, height/2, CreateStaticTile(width/2, height/2, TileTypeStaircaseDown()))
	//m.level.SetTileAt(width/3, height/3, CreateStaticTile(width/3, height/3, TileTypeStaircaseUp()))

	return m
}
