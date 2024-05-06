package state

import (
	"math/rand"
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/model"
	"mvvasilev/last_light/game/world"
	"mvvasilev/last_light/util"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
)

type PlayingState struct {
	player    *model.Player
	entityMap *world.EntityMap
	level     *world.MultilevelMap

	viewport *engine.Viewport

	movePlayerDirection model.Direction
	pauseGame           bool
	openInventory       bool
	pickUpUnderPlayer   bool
}

func BeginPlayingState() *PlayingState {
	s := new(PlayingState)

	mapSize := util.SizeOf(128, 128)

	dungeonLevel := world.CreateBSPDungeonMap(mapSize.Width(), mapSize.Height(), 4)

	itemTiles := spawnItems(dungeonLevel)

	itemLevel := world.CreateEmptyDungeonLevel(mapSize.Width(), mapSize.Height())

	for _, it := range itemTiles {
		itemLevel.SetTileAt(it.Position().X(), it.Position().Y(), it)
	}

	s.player = model.CreatePlayer(dungeonLevel.PlayerSpawnPoint().XY())

	s.entityMap = world.CreateEntityMap(mapSize.WH())

	s.level = world.CreateMultilevelMap(
		dungeonLevel,
		itemLevel,
		s.entityMap,
	)

	s.entityMap.AddEntity(s.player, '@', tcell.StyleDefault)

	s.viewport = engine.CreateViewport(
		util.PositionAt(0, 0),
		dungeonLevel.PlayerSpawnPoint(),
		util.SizeOf(80, 24),
		tcell.StyleDefault,
	)

	return s
}

func spawnItems(level *world.BSPDungeonMap) []world.Tile {
	rooms := level.Rooms()

	genTable := make(map[float32]*model.ItemType)

	genTable[0.2] = model.ItemTypeFish()
	genTable[0.05] = model.ItemTypeBow()
	genTable[0.051] = model.ItemTypeLongsword()
	genTable[0.052] = model.ItemTypeKey()

	itemTiles := make([]world.Tile, 0, 10)

	for _, r := range rooms {
		maxItems := int(0.10 * float64(r.Size().Area()))

		if maxItems < 1 {
			continue
		}

		numItems := rand.Intn(maxItems)

		for range numItems {
			itemType := model.GenerateItemType(genTable)

			if itemType == nil {
				continue
			}

			pos := util.PositionAt(
				util.RandInt(r.Position().X()+1, r.Position().X()+r.Size().Width()-1),
				util.RandInt(r.Position().Y()+1, r.Position().Y()+r.Size().Height()-1),
			)

			itemTiles = append(itemTiles, world.CreateItemTile(
				pos, itemType, 1,
			))
		}
	}

	return itemTiles
}

func (ps *PlayingState) Pause() {
	ps.pauseGame = true
}

func (ps *PlayingState) Unpause() {
	ps.pauseGame = false
}

func (ps *PlayingState) SetPaused(paused bool) {
	ps.pauseGame = paused
}

func (ps *PlayingState) MovePlayer() {
	if ps.movePlayerDirection == model.DirectionNone {
		return
	}

	newPlayerPos := ps.player.Position().WithOffset(model.MovementDirectionOffset(ps.movePlayerDirection))
	tileAtMovePos := ps.level.TileAt(newPlayerPos.XY())

	if tileAtMovePos.Passable() {
		dx, dy := model.MovementDirectionOffset(ps.movePlayerDirection)
		ps.entityMap.MoveEntity(ps.player.UniqueId(), dx, dy)
		ps.viewport.SetCenter(ps.player.Position())
	}

	ps.movePlayerDirection = model.DirectionNone
}

func (ps *PlayingState) PickUpItemUnderPlayer() {
	pos := ps.player.Position()
	tile := ps.level.TileAtHeight(pos.X(), pos.Y(), 1)

	itemTile, ok := tile.(*world.ItemTile)

	if !ok {
		return
	}

	item := model.CreateItem(itemTile.Type(), itemTile.Quantity())

	success := ps.player.Inventory().Push(item)

	if !success {
		return
	}

	ps.level.SetTileAtHeight(pos.X(), pos.Y(), 1, nil)
}

func (ps *PlayingState) OnInput(e *tcell.EventKey) {
	ps.player.Input(e)

	if e.Key() == tcell.KeyEsc {
		ps.pauseGame = true
		return
	}

	if e.Key() == tcell.KeyRune && e.Rune() == 'i' {
		ps.openInventory = true
		return
	}

	if e.Key() == tcell.KeyRune && e.Rune() == 'p' {
		ps.pickUpUnderPlayer = true
		return
	}

	switch e.Key() {
	case tcell.KeyUp:
		ps.movePlayerDirection = model.DirectionUp
	case tcell.KeyDown:
		ps.movePlayerDirection = model.DirectionDown
	case tcell.KeyLeft:
		ps.movePlayerDirection = model.DirectionLeft
	case tcell.KeyRight:
		ps.movePlayerDirection = model.DirectionRight
	}
}

func (ps *PlayingState) OnTick(dt int64) GameState {
	ps.player.Tick(dt)

	if ps.pauseGame {
		return PauseGame(ps)
	}

	if ps.openInventory {
		ps.openInventory = false
		return CreateInventoryScreenState(ps.player, ps)
	}

	if ps.movePlayerDirection != model.DirectionNone {
		ps.MovePlayer()
	}

	if ps.pickUpUnderPlayer {
		ps.pickUpUnderPlayer = false
		ps.PickUpItemUnderPlayer()
	}

	return ps
}

func (ps *PlayingState) CollectDrawables() []engine.Drawable {
	return engine.Multidraw(engine.CreateDrawingInstructions(func(v views.View) {
		ps.viewport.DrawFromProvider(v, func(x, y int) (rune, tcell.Style) {
			tile := ps.level.TileAt(x, y)

			if tile != nil {
				return tile.Presentation()
			}

			return ' ', tcell.StyleDefault
		})
	}))
}
