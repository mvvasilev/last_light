package model

import (
	"mvvasilev/last_light/util"
)

type Map interface {
	Size() util.Size
	SetTileAt(x, y int, t Tile)
	TileAt(x, y int) Tile
	Tick()
}
