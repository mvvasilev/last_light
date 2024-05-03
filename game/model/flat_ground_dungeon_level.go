package model

import (
	"math/rand"
	"mvvasilev/last_light/util"

	"github.com/gdamore/tcell/v2"
)

type FlatGroundDungeonLevel struct {
	tiles [][]Tile
}

func CreateFlatGroundDungeonLevel(width, height int) *FlatGroundDungeonLevel {
	level := new(FlatGroundDungeonLevel)

	level.tiles = make([][]Tile, height)

	for h := range height {
		level.tiles[h] = make([]Tile, width)

		for w := range width {
			if w == 0 || h == 0 || w >= width-1 || h >= height-1 {
				level.tiles[h][w] = CreateStaticTile(w, h, TileTypeRock())
				continue
			}

			level.tiles[h][w] = genRandomGroundTile(w, h)
		}
	}

	return level
}

func genRandomGroundTile(width, height int) Tile {
	switch rand.Intn(2) {
	case 0:
		return CreateStaticTile(width, height, TileTypeGround())
	case 1:
		return CreateStaticTile(width, height, TileTypeGrass())
	default:
		return CreateStaticTile(width, height, TileTypeGround())
	}
}

func (edl *FlatGroundDungeonLevel) Size() util.Size {
	return util.SizeOf(len(edl.tiles[0]), len(edl.tiles))
}

func (edl *FlatGroundDungeonLevel) SetTileAt(x int, y int, t Tile) {
	if len(edl.tiles) <= y || len(edl.tiles[0]) <= x {
		return
	}

	edl.tiles[y][x] = t
}

func (edl *FlatGroundDungeonLevel) TileAt(x int, y int) Tile {
	if y < 0 || y >= len(edl.tiles) {
		return nil
	}

	if x < 0 || x >= len(edl.tiles[y]) {
		return nil
	}

	return edl.tiles[y][x]
}

func (edl *FlatGroundDungeonLevel) Input(e *tcell.EventKey) {

}

func (edl *FlatGroundDungeonLevel) Tick() {
}
