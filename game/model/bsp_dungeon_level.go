package model

import (
	"math/rand"
	"mvvasilev/last_light/util"
)

type splitDirection bool

const (
	splitDirectionVertical   splitDirection = true
	splitDirectionHorizontal splitDirection = false
)

type bspNode struct {
	origin util.Position
	size   util.Size

	room    Room
	hasRoom bool

	left  *bspNode
	right *bspNode

	splitDir splitDirection
}

type Room struct {
	position util.Position
	size     util.Size
}

func (r Room) Size() util.Size {
	return r.size
}

func (r Room) Position() util.Position {
	return r.position
}

type BSPDungeonLevel struct {
	level *BasicMap

	playerSpawnPoint util.Position
	rooms            []Room
}

func CreateBSPDungeonLevel(width, height int, numSplits int) *BSPDungeonLevel {
	root := new(bspNode)

	root.origin = util.PositionAt(0, 0)
	root.size = util.SizeOf(width, height)

	split(root, numSplits)

	tiles := make([][]Tile, height)

	for h := range height {
		tiles[h] = make([]Tile, width)
	}

	rooms := make([]Room, 0, 2^numSplits)

	iterateBspLeaves(root, func(leaf *bspNode) {
		x := util.RandInt(leaf.origin.X(), leaf.origin.X()+leaf.size.Width()/4)
		y := util.RandInt(leaf.origin.Y(), leaf.origin.Y()+leaf.size.Height()/4)
		w := util.RandInt(3, leaf.size.Width()-1)
		h := util.RandInt(3, leaf.size.Height()-1)

		if x+w >= width {
			w = w - (x + w - width) - 1
		}

		if y+h >= height {
			h = h - (y + h - height) - 1
		}

		room := Room{
			position: util.PositionAt(x, y),
			size:     util.SizeOf(w, h),
		}

		rooms = append(rooms, room)

		makeRoom(tiles, room)

		leaf.room = room
		leaf.hasRoom = true
	})

	iterateBspParents(root, func(parent *bspNode) {
		roomLeft := findRoom(parent.left)
		roomRight := findRoom(parent.right)

		zCorridor(
			tiles,
			util.PositionAt(
				roomLeft.position.X()+roomLeft.size.Width()/2,
				roomLeft.position.Y()+roomLeft.size.Height()/2,
			),
			util.PositionAt(
				roomRight.position.X()+roomRight.size.Width()/2,
				roomRight.position.Y()+roomRight.size.Height()/2,
			),
			parent.splitDir,
		)
	})

	bsp := new(BSPDungeonLevel)

	spawnRoom := findRoom(root.left)

	bsp.rooms = rooms
	bsp.level = CreateBasicMap(tiles)
	bsp.playerSpawnPoint = util.PositionAt(
		spawnRoom.position.X()+spawnRoom.size.Width()/2,
		spawnRoom.position.Y()+spawnRoom.size.Height()/2,
	)

	return bsp
}

func (bsp *BSPDungeonLevel) PlayerSpawnPoint() util.Position {
	return bsp.playerSpawnPoint
}

func findRoom(parent *bspNode) Room {
	if parent.hasRoom {
		return parent.room
	}

	if rand.Float32() > 0.5 {
		return findRoom(parent.left)
	} else {
		return findRoom(parent.right)
	}
}

func zCorridor(tiles [][]Tile, from util.Position, to util.Position, direction splitDirection) {
	switch direction {
	case splitDirectionHorizontal:
		xMidPoint := (from.X() + to.X()) / 2
		horizontalTunnel(tiles, from.X(), xMidPoint, from.Y())
		horizontalTunnel(tiles, to.X(), xMidPoint, to.Y())
		verticalTunnel(tiles, from.Y(), to.Y(), xMidPoint)
	case splitDirectionVertical:
		yMidPoint := (from.Y() + to.Y()) / 2
		verticalTunnel(tiles, from.Y(), yMidPoint, from.X())
		verticalTunnel(tiles, to.Y(), yMidPoint, to.X())
		horizontalTunnel(tiles, from.X(), to.X(), yMidPoint)
	}
}

func iterateBspParents(parent *bspNode, iter func(parent *bspNode)) {
	if parent.left != nil && parent.right != nil {
		iter(parent)
	}

	if parent.left != nil {
		iterateBspParents(parent.left, iter)
	}

	if parent.right != nil {
		iterateBspParents(parent.right, iter)
	}
}

func iterateBspLeaves(parent *bspNode, iter func(leaf *bspNode)) {
	if parent.left == nil && parent.right == nil {
		iter(parent)
		return
	}

	if parent.left != nil {
		iterateBspLeaves(parent.left, iter)
	}

	if parent.right != nil {
		iterateBspLeaves(parent.right, iter)
	}
}

func split(parent *bspNode, numSplits int) {
	if numSplits <= 0 {
		return
	}

	// split vertically
	if parent.size.Width() > parent.size.Height() {
		// New splits will be between 45% and 65% of the parent's width
		leftSplitWidth := util.RandInt(int(float32(parent.size.Width())*0.45), int(float32(parent.size.Width())*0.65))

		parent.splitDir = splitDirectionVertical

		parent.left = new(bspNode)
		parent.left.origin = parent.origin
		parent.left.size = util.SizeOf(leftSplitWidth, parent.size.Height())

		parent.right = new(bspNode)
		parent.right.origin = parent.origin.WithOffset(leftSplitWidth, 0)
		parent.right.size = util.SizeOf(parent.size.Width()-leftSplitWidth, parent.size.Height())
	} else { // split horizontally
		// New splits will be between 45% and 65% of the parent's height
		leftSplitHeight := util.RandInt(int(float32(parent.size.Height())*0.45), int(float32(parent.size.Height())*0.65))

		parent.splitDir = splitDirectionHorizontal

		parent.left = new(bspNode)
		parent.left.origin = parent.origin
		parent.left.size = util.SizeOf(parent.size.Width(), leftSplitHeight)

		parent.right = new(bspNode)
		parent.right.origin = parent.origin.WithOffset(0, leftSplitHeight)
		parent.right.size = util.SizeOf(parent.size.Width(), parent.size.Height()-leftSplitHeight)
	}

	split(parent.left, numSplits-1)
	split(parent.right, numSplits-1)
}

func horizontalTunnel(tiles [][]Tile, x1, x2, y int) {
	if x1 > x2 {
		tx := x2
		x2 = x1
		x1 = tx
	}

	placeWallAtIfNotPassable(tiles, x1, y-1)
	placeWallAtIfNotPassable(tiles, x1, y)
	placeWallAtIfNotPassable(tiles, x1, y+1)

	for x := x1; x <= x2; x++ {
		if tiles[y][x] != nil && tiles[y][x].Passable() {
			continue
		}

		tiles[y][x] = CreateStaticTile(x, y, TileTypeGround())

		placeWallAtIfNotPassable(tiles, x, y-1)
		placeWallAtIfNotPassable(tiles, x, y+1)
	}

	placeWallAtIfNotPassable(tiles, x2, y-1)
	placeWallAtIfNotPassable(tiles, x2, y)
	placeWallAtIfNotPassable(tiles, x2, y+1)
}

func verticalTunnel(tiles [][]Tile, y1, y2, x int) {
	if y1 > y2 {
		ty := y2
		y2 = y1
		y1 = ty
	}

	placeWallAtIfNotPassable(tiles, x-1, y1)
	placeWallAtIfNotPassable(tiles, x, y1)
	placeWallAtIfNotPassable(tiles, x+1, y1)

	for y := y1; y <= y2; y++ {
		if tiles[y][x] != nil && tiles[y][x].Passable() {
			continue
		}

		tiles[y][x] = CreateStaticTile(x, y, TileTypeGround())

		placeWallAtIfNotPassable(tiles, x-1, y)
		placeWallAtIfNotPassable(tiles, x+1, y)
	}

	placeWallAtIfNotPassable(tiles, x-1, y2)
	placeWallAtIfNotPassable(tiles, x, y2)
	placeWallAtIfNotPassable(tiles, x+1, y2)
}

func placeWallAtIfNotPassable(tiles [][]Tile, x, y int) {
	if tiles[y][x] != nil && tiles[y][x].Passable() {
		return
	}

	tiles[y][x] = CreateStaticTile(x, y, TileTypeWall())
}

func makeRoom(tiles [][]Tile, room Room) {
	width := room.size.Width()
	height := room.size.Height()
	x := room.position.X()
	y := room.position.Y()

	for w := x; w < x+width+1; w++ {
		tiles[y][w] = CreateStaticTile(w, y, TileTypeWall())
		tiles[y+height][w] = CreateStaticTile(w, y+height, TileTypeWall())
	}

	for h := y; h < y+height+1; h++ {
		tiles[h][x] = CreateStaticTile(x, h, TileTypeWall())
		tiles[h][x+width] = CreateStaticTile(x+width, h, TileTypeWall())
	}

	for h := y + 1; h < y+height; h++ {
		for w := x + 1; w < x+width; w++ {
			tiles[h][w] = CreateStaticTile(w, h, TileTypeGround())
		}
	}
}

func (bsp *BSPDungeonLevel) Size() util.Size {
	return bsp.level.Size()
}

func (bsp *BSPDungeonLevel) SetTileAt(x int, y int, t Tile) {
	bsp.level.SetTileAt(x, y, t)
}

func (bsp *BSPDungeonLevel) TileAt(x int, y int) Tile {
	return bsp.level.TileAt(x, y)
}

func (bsp *BSPDungeonLevel) Tick() {
}

func (bsp *BSPDungeonLevel) Rooms() []Room {
	return bsp.rooms
}
