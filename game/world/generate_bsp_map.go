package world

import (
	"math/rand"
	"mvvasilev/last_light/engine"
)

type splitDirection bool

const (
	splitDirectionVertical   splitDirection = true
	splitDirectionHorizontal splitDirection = false
)

type bspNode struct {
	origin engine.Position
	size   engine.Size

	room    engine.BoundingBox
	hasRoom bool

	left  *bspNode
	right *bspNode

	splitDir splitDirection
}

func CreateBSPDungeonMap(width, height int, numSplits int) *BSPDungeonMap {
	root := new(bspNode)

	root.origin = engine.PositionAt(0, 0)
	root.size = engine.SizeOf(width, height)

	split(root, numSplits)

	tiles := make([][]Tile, height)

	for h := range height {
		tiles[h] = make([]Tile, width)
	}

	rooms := make([]engine.BoundingBox, 0, 2^numSplits)

	iterateBspLeaves(root, func(leaf *bspNode) {
		x := engine.RandInt(leaf.origin.X(), leaf.origin.X()+leaf.size.Width()/4)
		y := engine.RandInt(leaf.origin.Y(), leaf.origin.Y()+leaf.size.Height()/4)
		w := engine.RandInt(3, leaf.size.Width()-1)
		h := engine.RandInt(3, leaf.size.Height()-1)

		if x+w >= width {
			w = w - (x + w - width) - 1
		}

		if y+h >= height {
			h = h - (y + h - height) - 1
		}

		room := engine.BoundingBox{
			Positioned: engine.WithPosition(engine.PositionAt(x, y)),
			Sized:      engine.WithSize(engine.SizeOf(w, h)),
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
			engine.PositionAt(
				roomLeft.Position().X()+roomLeft.Size().Width()/2,
				roomLeft.Position().Y()+roomLeft.Size().Height()/2,
			),
			engine.PositionAt(
				roomRight.Position().X()+roomRight.Size().Width()/2,
				roomRight.Position().Y()+roomRight.Size().Height()/2,
			),
			parent.splitDir,
		)
	})

	bsp := new(BSPDungeonMap)

	spawnRoom := findRoom(root.left)

	bsp.rooms = rooms
	bsp.level = CreateBasicMap(tiles)
	bsp.playerSpawnPoint = engine.PositionAt(
		spawnRoom.Position().X()+spawnRoom.Size().Width()/2,
		spawnRoom.Position().Y()+spawnRoom.Size().Height()/2,
	)

	return bsp
}

func findRoom(parent *bspNode) engine.BoundingBox {
	if parent.hasRoom {
		return parent.room
	}

	if rand.Float32() > 0.5 {
		return findRoom(parent.left)
	} else {
		return findRoom(parent.right)
	}
}

func zCorridor(tiles [][]Tile, from engine.Position, to engine.Position, direction splitDirection) {
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
		leftSplitWidth := engine.RandInt(int(float32(parent.size.Width())*0.45), int(float32(parent.size.Width())*0.65))

		parent.splitDir = splitDirectionVertical

		parent.left = new(bspNode)
		parent.left.origin = parent.origin
		parent.left.size = engine.SizeOf(leftSplitWidth, parent.size.Height())

		parent.right = new(bspNode)
		parent.right.origin = parent.origin.WithOffset(leftSplitWidth, 0)
		parent.right.size = engine.SizeOf(parent.size.Width()-leftSplitWidth, parent.size.Height())
	} else { // split horizontally
		// New splits will be between 45% and 65% of the parent's height
		leftSplitHeight := engine.RandInt(int(float32(parent.size.Height())*0.45), int(float32(parent.size.Height())*0.65))

		parent.splitDir = splitDirectionHorizontal

		parent.left = new(bspNode)
		parent.left.origin = parent.origin
		parent.left.size = engine.SizeOf(parent.size.Width(), leftSplitHeight)

		parent.right = new(bspNode)
		parent.right.origin = parent.origin.WithOffset(0, leftSplitHeight)
		parent.right.size = engine.SizeOf(parent.size.Width(), parent.size.Height()-leftSplitHeight)
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

func makeRoom(tiles [][]Tile, room engine.BoundingBox) {
	width := room.Size().Width()
	height := room.Size().Height()
	x := room.Position().X()
	y := room.Position().Y()

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
