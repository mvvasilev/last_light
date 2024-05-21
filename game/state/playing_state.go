package state

import (
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/model"
	"mvvasilev/last_light/game/player"
	"mvvasilev/last_light/game/ui"
	"mvvasilev/last_light/game/world"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
)

type PlayingState struct {
	player  *player.Player
	someNPC *model.BasicNPC

	dungeon *world.Dungeon

	viewport *engine.Viewport

	movePlayerDirection model.Direction
	pauseGame           bool
	openInventory       bool
	pickUpUnderPlayer   bool
	interact            bool
	moveEntities        bool

	nextGameState GameState
}

func BeginPlayingState() *PlayingState {
	s := new(PlayingState)

	mapSize := engine.SizeOf(128, 128)

	s.dungeon = world.CreateDungeon(mapSize.Width(), mapSize.Height(), 1)

	s.player = player.CreatePlayer(s.dungeon.CurrentLevel().PlayerSpawnPoint().XY())

	s.someNPC = model.CreateNPC(s.dungeon.CurrentLevel().NextLevelStaircase())

	s.dungeon.CurrentLevel().AddEntity(s.player, '@', tcell.StyleDefault)
	s.dungeon.CurrentLevel().AddEntity(s.someNPC, 'N', tcell.StyleDefault)

	s.viewport = engine.CreateViewport(
		engine.PositionAt(0, 0),
		s.dungeon.CurrentLevel().PlayerSpawnPoint(),
		engine.SizeOf(80, 24),
		tcell.StyleDefault,
	)

	s.nextGameState = s

	return s
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

	if ps.dungeon.CurrentLevel().IsTilePassable(newPlayerPos.XY()) {
		dx, dy := model.MovementDirectionOffset(ps.movePlayerDirection)
		ps.dungeon.CurrentLevel().MoveEntity(ps.player.UniqueId(), dx, dy)
		ps.viewport.SetCenter(ps.player.Position())
	}

	ps.movePlayerDirection = model.DirectionNone
}

func (ps *PlayingState) InteractBelowPlayer() {
	playerPos := ps.player.Position()

	if playerPos == ps.dungeon.CurrentLevel().NextLevelStaircase() {
		ps.SwitchToNextLevel()
		return
	}

	if playerPos == ps.dungeon.CurrentLevel().PreviousLevelStaircase() {
		ps.SwitchToPreviousLevel()
		return
	}
}

func (ps *PlayingState) SwitchToNextLevel() {
	if !ps.dungeon.HasNextLevel() {
		ps.nextGameState = CreateDialogState(
			ui.CreateOkDialog(
				"The Unknown Depths",
				"The staircases descent down to the lower levels is seemingly blocked by multiple large boulders. They appear immovable.",
				"Continue",
				40,
				func() {
					ps.nextGameState = ps
				},
			),
			ps,
		)

		return
	}

	ps.dungeon.CurrentLevel().DropEntity(ps.player.UniqueId())

	ps.dungeon.MoveToNextLevel()

	ps.player.MoveTo(ps.dungeon.CurrentLevel().PlayerSpawnPoint())

	ps.viewport = engine.CreateViewport(
		engine.PositionAt(0, 0),
		ps.dungeon.CurrentLevel().PlayerSpawnPoint(),
		engine.SizeOf(80, 24),
		tcell.StyleDefault,
	)

	ps.dungeon.CurrentLevel().AddEntity(ps.player, '@', tcell.StyleDefault)
}

func (ps *PlayingState) SwitchToPreviousLevel() {
	if !ps.dungeon.HasPreviousLevel() {
		ps.nextGameState = CreateDialogState(
			ui.CreateOkDialog(
				"The Surface",
				"You feel the gentle, yet chilling breeze of the surface make its way through the weaving cavern tunnels, the very same you had to make your way through to get where you are. There is nothing above that you need. Find the last light, or die trying.",
				"Continue",
				40,
				func() {
					ps.nextGameState = ps
				},
			),
			ps,
		)

		return
	}

	ps.dungeon.CurrentLevel().DropEntity(ps.player.UniqueId())

	ps.dungeon.MoveToPreviousLevel()

	ps.player.MoveTo(ps.dungeon.CurrentLevel().NextLevelStaircase())

	ps.viewport = engine.CreateViewport(
		engine.PositionAt(0, 0),
		ps.dungeon.CurrentLevel().NextLevelStaircase(),
		engine.SizeOf(80, 24),
		tcell.StyleDefault,
	)

	ps.dungeon.CurrentLevel().AddEntity(ps.player, '@', tcell.StyleDefault)
}

func (ps *PlayingState) PickUpItemUnderPlayer() {
	pos := ps.player.Position()
	item := ps.dungeon.CurrentLevel().RemoveItemAt(pos.XY())

	if item == nil {
		return
	}

	success := ps.player.Inventory().Push(item)

	if !success {
		ps.dungeon.CurrentLevel().SetItemAt(pos.X(), pos.Y(), item)
	}
}

func (ps *PlayingState) CalcPathToPlayerAndMove() {
	distanceToPlayer := ps.someNPC.Position().Distance(ps.player.Position())

	if distanceToPlayer > 20 {
		return
	}

	pathToPlayer := engine.FindPath(
		ps.someNPC.Position(),
		ps.player.Position(),
		func(x, y int) bool {
			if x == ps.player.Position().X() && y == ps.player.Position().Y() {
				return true
			}

			return ps.dungeon.CurrentLevel().IsTilePassable(x, y)
		},
	)

	nextPos, hasNext := pathToPlayer.Next()

	if !hasNext {
		return
	}

	if nextPos.Equals(ps.player.Position()) {
		return
	}

	ps.dungeon.CurrentLevel().MoveEntityTo(ps.someNPC.UniqueId(), nextPos.X(), nextPos.Y())
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

	if e.Key() == tcell.KeyRune && e.Rune() == 'e' {
		ps.interact = true
		return
	}

	switch e.Key() {
	case tcell.KeyUp:
		ps.movePlayerDirection = model.DirectionUp
		ps.moveEntities = true
	case tcell.KeyDown:
		ps.movePlayerDirection = model.DirectionDown
		ps.moveEntities = true
	case tcell.KeyLeft:
		ps.movePlayerDirection = model.DirectionLeft
		ps.moveEntities = true
	case tcell.KeyRight:
		ps.movePlayerDirection = model.DirectionRight
		ps.moveEntities = true
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

	if ps.interact {
		ps.interact = false
		ps.InteractBelowPlayer()
	}

	if ps.moveEntities {
		ps.moveEntities = false
		ps.CalcPathToPlayerAndMove()
	}

	return ps.nextGameState
}

func (ps *PlayingState) CollectDrawables() []engine.Drawable {
	return engine.Multidraw(engine.CreateDrawingInstructions(func(v views.View) {
		visibilityMap := engine.ComputeFOV(
			func(x, y int) world.Tile {
				ps.dungeon.CurrentLevel().Flatten().MarkExplored(x, y)

				return ps.dungeon.CurrentLevel().TileAt(x, y)
			},
			func(x, y int) bool { return ps.dungeon.CurrentLevel().Flatten().IsInBounds(x, y) },
			func(x, y int) bool { return ps.dungeon.CurrentLevel().Flatten().TileAt(x, y).Opaque() },
			ps.player.Position().X(), ps.player.Position().Y(),
			13,
		)

		ps.viewport.DrawFromProvider(v, func(x, y int) (rune, tcell.Style) {
			tile := visibilityMap[engine.PositionAt(x, y)]

			if tile != nil {
				return tile.Presentation()
			}

			explored := ps.dungeon.CurrentLevel().Flatten().ExploredTileAt(x, y)

			if explored != nil {
				return explored.Presentation()
			}

			return ' ', tcell.StyleDefault
		})
	}))
}
