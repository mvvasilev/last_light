package model

import (
	"mvvasilev/last_light/engine"

	"github.com/gdamore/tcell/v2"
)

type Map interface {
	Size() engine.Size
	Tiles() [][]Tile
	ExploredTiles() map[engine.Position]Tile
	ExploredTileStyle() tcell.Style
	DefaultTile() Tile

	PlayerSpawnPoint() *Map_PlayerSpawnPointComponent
	Rooms() *Map_RoomsComponent
	NextLevelStaircase() *Map_NextLevelStaircaseComponent
	PreviousLevelStaircase() *Map_PreviousLevelStaircaseComponent
}

type Map_PlayerSpawnPointComponent struct {
	Position engine.Position
}

type Map_RoomsComponent struct {
	Rooms []engine.BoundingBox
}

type Map_NextLevelStaircaseComponent struct {
	Position engine.Position
}

type Map_PreviousLevelStaircaseComponent struct {
	Position engine.Position
}

type BaseMap struct {
	size          engine.Size
	tiles         [][]Tile
	exploredTiles map[engine.Position]Tile
	exploredStyle tcell.Style
	defaultTile   Tile

	playerSpawnPos *Map_PlayerSpawnPointComponent
	rooms          *Map_RoomsComponent
	nextLevel      *Map_NextLevelStaircaseComponent
	prevLevel      *Map_PreviousLevelStaircaseComponent
}

func CreateMap(size engine.Size, tiles [][]Tile, exploredStyle tcell.Style, defaultTile Tile, components ...func(*BaseMap)) Map {
	m := &BaseMap{
		size:          size,
		tiles:         tiles,
		exploredTiles: make(map[engine.Position]Tile, 0),
		exploredStyle: exploredStyle,
		defaultTile:   defaultTile,
	}

	for _, c := range components {
		c(m)
	}

	return m
}

func (m *BaseMap) Size() engine.Size {
	return m.size
}

func (m *BaseMap) Tiles() [][]Tile {
	return m.tiles
}

func (m *BaseMap) ExploredTiles() map[engine.Position]Tile {
	return m.exploredTiles
}

func (m *BaseMap) ExploredTileStyle() tcell.Style {
	return m.exploredStyle
}

func (m *BaseMap) DefaultTile() Tile {
	return m.defaultTile
}

func (m *BaseMap) PlayerSpawnPoint() *Map_PlayerSpawnPointComponent {
	return m.playerSpawnPos
}

func (m *BaseMap) Rooms() *Map_RoomsComponent {
	return m.rooms
}

func (m *BaseMap) NextLevelStaircase() *Map_NextLevelStaircaseComponent {
	return m.nextLevel
}

func (m *BaseMap) PreviousLevelStaircase() *Map_PreviousLevelStaircaseComponent {
	return m.prevLevel
}

func Map_WithRooms(rooms []engine.BoundingBox) func(*BaseMap) {
	return func(bm *BaseMap) {
		bm.rooms = &Map_RoomsComponent{
			Rooms: rooms,
		}
	}
}

func Map_WithPlayerSpawnPoint(pos engine.Position) func(*BaseMap) {
	return func(bm *BaseMap) {
		bm.playerSpawnPos = &Map_PlayerSpawnPointComponent{
			Position: pos,
		}
	}
}

func Map_WithNextLevelStaircase(pos engine.Position) func(*BaseMap) {
	return func(bm *BaseMap) {
		bm.nextLevel = &Map_NextLevelStaircaseComponent{
			Position: pos,
		}
	}
}

func Map_WithPreviousLevelStaircase(pos engine.Position) func(*BaseMap) {
	return func(bm *BaseMap) {
		bm.prevLevel = &Map_PreviousLevelStaircaseComponent{
			Position: pos,
		}
	}
}

func Map_SetTileAt(bm Map, x int, y int, t Tile) Tile {
	if !Map_IsInBounds(bm, x, y) {
		return bm.DefaultTile()
	}

	bm.Tiles()[y][x] = t

	return bm.Tiles()[y][x]
}

func Map_TileAt(bm Map, x int, y int) Tile {
	if !Map_IsInBounds(bm, x, y) {
		return bm.DefaultTile()
	}

	tile := bm.Tiles()[y][x]

	if tile == nil {
		return bm.DefaultTile()
	}

	return tile
}

func Map_IsInBounds(bm Map, x, y int) bool {
	if x < 0 || y < 0 {
		return false
	}

	if x >= bm.Size().Width() || y >= bm.Size().Height() {
		return false
	}

	return true
}

func Map_ExploredTileAt(bm Map, x, y int) Tile {
	return bm.ExploredTiles()[engine.PositionAt(x, y)]
}

func Map_MarkExplored(bm Map, x, y int) {
	if !Map_IsInBounds(bm, x, y) {
		return
	}

	tile := Map_TileAt(bm, x, y)

	symbol, _ := tile.DefaultPresentation()

	bm.ExploredTiles()[engine.PositionAt(x, y)] = &BaseTile{
		defaultSymbol: symbol,
		defaultStyle:  bm.ExploredTileStyle(),
	}
}
