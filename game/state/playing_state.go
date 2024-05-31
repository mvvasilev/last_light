package state

import (
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/input"
	"mvvasilev/last_light/game/npc"
	"mvvasilev/last_light/game/player"
	"mvvasilev/last_light/game/rpg"
	"mvvasilev/last_light/game/turns"
	"mvvasilev/last_light/game/ui"
	"mvvasilev/last_light/game/world"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
)

type PlayingState struct {
	turnSystem  *turns.TurnSystem
	inputSystem *input.InputSystem

	player  *player.Player
	someNPC *npc.BasicNPC

	eventLog   *engine.GameEventLog
	uiEventLog *ui.UIEventLog

	healthBar *ui.UIHealthBar

	dungeon *world.Dungeon

	viewport *engine.Viewport

	viewShortLogs bool

	nextGameState GameState
}

func CreatePlayingState(turnSystem *turns.TurnSystem, inputSystem *input.InputSystem, playerStats map[rpg.Stat]int) *PlayingState {
	turnSystem.Clear()

	s := new(PlayingState)

	s.turnSystem = turnSystem
	s.inputSystem = inputSystem

	mapSize := engine.SizeOf(128, 128)

	s.dungeon = world.CreateDungeon(mapSize.Width(), mapSize.Height(), 1)

	s.player = player.CreatePlayer(
		s.dungeon.CurrentLevel().PlayerSpawnPoint().X(),
		s.dungeon.CurrentLevel().PlayerSpawnPoint().Y(),
		playerStats,
	)
	s.player.Heal(rpg.BaseMaxHealth(s.player))

	s.turnSystem.Schedule(10, func() (complete bool, requeue bool) {
		requeue = true
		complete = false

		switch inputSystem.NextAction() {
		case input.InputAction_PauseGame:
			s.nextGameState = PauseGame(s, s.turnSystem, s.inputSystem)
		case input.InputAction_OpenInventory:
			s.nextGameState = CreateInventoryScreenState(s.inputSystem, s.turnSystem, s.player, s)
		case input.InputAction_PickUpItem:
			s.PickUpItemUnderPlayer()
			complete = true
		case input.InputAction_Interact:
			s.InteractBelowPlayer()
			complete = true
		case input.InputAction_OpenLogs:
			s.viewShortLogs = !s.viewShortLogs
		case input.InputAction_MovePlayer_East:
			s.MovePlayer(npc.East)
			complete = true
		case input.InputAction_MovePlayer_West:
			s.MovePlayer(npc.West)
			complete = true
		case input.InputAction_MovePlayer_North:
			s.MovePlayer(npc.North)
			complete = true
		case input.InputAction_MovePlayer_South:
			s.MovePlayer(npc.South)
			complete = true
		default:
		}

		return
	})

	s.someNPC = npc.CreateNPC(s.dungeon.CurrentLevel().NextLevelStaircase())

	s.turnSystem.Schedule(20, func() (complete bool, requeue bool) {
		s.CalcPathToPlayerAndMove()

		return true, true
	})

	s.eventLog = engine.CreateGameEventLog(100)

	s.uiEventLog = ui.CreateUIEventLog(0, 17, 80, 7, s.eventLog, tcell.StyleDefault)
	s.healthBar = ui.CreateHealthBar(68, 0, 12, 3, s.player.CurrentHealth(), rpg.BaseMaxHealth(s.player), tcell.StyleDefault)

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

func (s *PlayingState) InputContext() input.Context {
	return input.InputContext_Play
}

func (ps *PlayingState) MovePlayer(direction npc.Direction) {
	if direction == npc.DirectionNone {
		return
	}

	newPlayerPos := ps.player.Position().WithOffset(npc.MovementDirectionOffset(direction))

	if ps.dungeon.CurrentLevel().IsTilePassable(newPlayerPos.XY()) {
		dx, dy := npc.MovementDirectionOffset(direction)
		ps.dungeon.CurrentLevel().MoveEntity(ps.player.UniqueId(), dx, dy)
		ps.viewport.SetCenter(ps.player.Position())
	}

	ps.eventLog.Log("You moved " + npc.DirectionName(direction))
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
			ps.inputSystem,
			ps.turnSystem,
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
			ps.inputSystem,
			ps.turnSystem,
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
		return
	}

	itemName, _ := item.Name()

	ps.eventLog.Log("You picked up " + itemName)
}

func (ps *PlayingState) HasLineOfSight(start, end engine.Position) bool {
	positions := engine.CastRay(start, end)

	for _, p := range positions {
		if ps.dungeon.CurrentLevel().IsGroundTileOpaque(p.XY()) {
			return false
		}
	}

	return true
}

func (ps *PlayingState) CalcPathToPlayerAndMove() {
	playerVisibleAndInRange := false

	if ps.someNPC.Position().Distance(ps.player.Position()) < 20 && ps.HasLineOfSight(ps.someNPC.Position(), ps.player.Position()) {
		playerVisibleAndInRange = true
	}

	if !playerVisibleAndInRange {
		randomMove := npc.Direction(engine.RandInt(int(npc.DirectionNone), int(npc.East)))

		nextPos := ps.someNPC.Position()

		switch randomMove {
		case npc.North:
			nextPos = nextPos.WithOffset(0, -1)
		case npc.South:
			nextPos = nextPos.WithOffset(0, +1)
		case npc.West:
			nextPos = nextPos.WithOffset(-1, 0)
		case npc.East:
			nextPos = nextPos.WithOffset(+1, 0)
		default:
			return
		}

		if ps.dungeon.CurrentLevel().IsTilePassable(nextPos.XY()) {
			ps.dungeon.CurrentLevel().MoveEntityTo(
				ps.someNPC.UniqueId(),
				nextPos.X(),
				nextPos.Y(),
			)
		}

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

func (ps *PlayingState) OnTick(dt int64) (nextState GameState) {
	ps.nextGameState = ps

	ps.turnSystem.NextTurn()

	return ps.nextGameState
}

func (ps *PlayingState) CollectDrawables() []engine.Drawable {
	mainCameraDrawingInstructions := engine.CreateDrawingInstructions(func(v views.View) {
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
	})

	drawables := []engine.Drawable{}

	drawables = append(drawables, mainCameraDrawingInstructions)

	if ps.viewShortLogs {
		drawables = append(drawables, ps.uiEventLog)
	}

	drawables = append(drawables, ps.healthBar)

	return drawables
}
