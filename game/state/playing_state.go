package state

import (
	"math/rand"
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/model"
	"mvvasilev/last_light/game/systems"
	"mvvasilev/last_light/game/ui"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
)

type PlayingState struct {
	turnSystem  *systems.TurnSystem
	inputSystem *systems.InputSystem

	player *model.Player
	npcs   []model.Entity

	eventLog   *engine.GameEventLog
	uiEventLog *ui.UIEventLog

	healthBar *ui.UIHealthBar

	dungeon *model.Dungeon

	viewport *engine.Viewport

	viewShortLogs bool

	nextGameState GameState
}

func CreatePlayingState(turnSystem *systems.TurnSystem, inputSystem *systems.InputSystem, playerStats map[model.Stat]int) *PlayingState {
	turnSystem.Clear()

	s := new(PlayingState)

	s.turnSystem = turnSystem
	s.inputSystem = inputSystem

	mapSize := engine.SizeOf(128, 128)

	s.dungeon = model.CreateDungeon(mapSize.Width(), mapSize.Height(), 1)

	s.player = model.CreatePlayer(
		s.dungeon.CurrentLevel().Ground().PlayerSpawnPoint().Position.X(),
		s.dungeon.CurrentLevel().Ground().PlayerSpawnPoint().Position.Y(),
		playerStats,
	)

	s.turnSystem.Schedule(s.player.DefaultSpeed().Speed, func() (complete bool, requeue bool) {
		requeue = true
		complete = false

		if s.player.HealthData().Health <= 0 || s.player.HealthData().IsDead {
			s.nextGameState = CreateGameOverState(inputSystem)
		}

		switch inputSystem.NextAction() {
		case systems.InputAction_PauseGame:
			s.nextGameState = PauseGame(s, s.turnSystem, s.inputSystem)
		case systems.InputAction_OpenInventory:
			s.nextGameState = CreateInventoryScreenState(s.eventLog, s.dungeon, s.inputSystem, s.turnSystem, s.player, s)
		case systems.InputAction_EnterLookMode:
			s.viewShortLogs = false
			s.nextGameState = CreateLookState(s, s.eventLog, s.dungeon, s.inputSystem, s.turnSystem, s.player)
		case systems.InputAction_PickUpItem:
			complete = PickUpItemUnderPlayer(s.eventLog, s.dungeon, s.player)
		case systems.InputAction_Interact:
			complete = s.InteractBelowPlayer()
		case systems.InputAction_OpenLogs:
			s.viewShortLogs = !s.viewShortLogs
		case systems.InputAction_Move_East:
			complete = s.MovePlayer(model.East)
		case systems.InputAction_Move_West:
			complete = s.MovePlayer(model.West)
		case systems.InputAction_Move_North:
			complete = s.MovePlayer(model.North)
		case systems.InputAction_Move_South:
			complete = s.MovePlayer(model.South)
		default:
		}

		if s.player.IsNextTurnSkipped() {
			s.player.SkipNextTurn(false)
			complete = true
		}

		return
	})

	s.eventLog = engine.CreateGameEventLog(100)

	s.uiEventLog = ui.CreateUIEventLog(0, 17, 80, 7, s.eventLog, tcell.StyleDefault)
	s.healthBar = ui.CreateHealthBar(68, 0, 12, 3, s.player, tcell.StyleDefault)

	s.dungeon.CurrentLevel().AddEntity(s.player)

	entityTable := model.CreateEntityTable()

	entityTable.Add(1, func(x, y int) model.Entity {
		return model.Entity_Imp(x, y, model.HostileNPCBehavior(s.eventLog, s.dungeon, s.player))
	})
	entityTable.Add(1, func(x, y int) model.Entity {
		return model.Entity_SkeletalKnight(x, y, model.HostileNPCBehavior(s.eventLog, s.dungeon, s.player))
	})
	entityTable.Add(1, func(x, y int) model.Entity {
		return model.Entity_SkeletalWarrior(x, y, model.HostileNPCBehavior(s.eventLog, s.dungeon, s.player))
	})

	s.npcs = SpawnNPCs(s.dungeon, 7, entityTable)

	for _, npc := range s.npcs {
		if npc.Behavior() != nil {
			speed := npc.Behavior().Speed
			s.turnSystem.Schedule(speed, npc.Behavior().Behavior)
		}
	}

	s.viewport = engine.CreateViewport(
		engine.PositionAt(0, 0),
		s.dungeon.CurrentLevel().Ground().PlayerSpawnPoint().Position,
		engine.SizeOf(80, 24),
		tcell.StyleDefault,
	)

	s.nextGameState = s

	return s
}

func (s *PlayingState) InputContext() systems.InputContext {
	return systems.InputContext_Play
}

func SpawnNPCs(dungeon *model.Dungeon, number int, genTable *model.EntityTable) []model.Entity {
	rooms := dungeon.CurrentLevel().Ground().Rooms().Rooms

	entities := make([]model.Entity, 0, number)

	for range number {
		r := rooms[rand.Intn(len(rooms))]

		x, y := engine.RandInt(r.Position().X()+1, r.Position().X()+r.Size().Width()-1), engine.RandInt(r.Position().Y()+1, r.Position().Y()+r.Size().Height()-1)

		entity := genTable.Generate(x, y)

		entities = append(entities, entity)

		dungeon.CurrentLevel().AddEntity(entity)
	}

	return entities
}

func (ps *PlayingState) MovePlayer(direction model.Direction) (success bool) {
	if direction == model.DirectionNone {
		return true
	}

	newPlayerPos := ps.player.Position().WithOffset(model.MovementDirectionOffset(direction))

	ent := ps.dungeon.CurrentLevel().EntityAt(newPlayerPos.XY())

	// We are moving into an entity with health data. Attack it.
	if ent != nil && ent.HealthData() != nil {
		if ent.HealthData().IsDead {
			// TODO: If the entity is dead, the player should be able to move through it.
			return false
		}

		model.ExecuteAttack(ps.eventLog, ps.player, ent)

		return true
	}

	if ps.dungeon.CurrentLevel().IsTilePassable(newPlayerPos.XY()) {
		ps.dungeon.CurrentLevel().MoveEntityTo(ps.player.UniqueId(), newPlayerPos.X(), newPlayerPos.Y())
		ps.viewport.SetCenter(ps.player.Position())

		ps.eventLog.Log("You moved " + model.DirectionName(direction))

		return true
	} else {
		ps.eventLog.Log("You bump into an impassable object")

		return false
	}
}

func (ps *PlayingState) InteractBelowPlayer() (success bool) {
	playerPos := ps.player.Position()

	if playerPos == ps.dungeon.CurrentLevel().Ground().NextLevelStaircase().Position {
		ps.SwitchToNextLevel()
		return true
	}

	if playerPos == ps.dungeon.CurrentLevel().Ground().PreviousLevelStaircase().Position {
		ps.SwitchToPreviousLevel()
		return true
	}

	ps.eventLog.Log("There is nothing to interact with here")

	return false
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

	ps.player.Positioned().Position = ps.dungeon.CurrentLevel().Ground().PlayerSpawnPoint().Position

	ps.viewport = engine.CreateViewport(
		engine.PositionAt(0, 0),
		ps.dungeon.CurrentLevel().Ground().PlayerSpawnPoint().Position,
		engine.SizeOf(80, 24),
		tcell.StyleDefault,
	)

	ps.dungeon.CurrentLevel().AddEntity(ps.player)
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

	ps.player.Positioned().Position = ps.dungeon.CurrentLevel().Ground().NextLevelStaircase().Position

	ps.viewport = engine.CreateViewport(
		engine.PositionAt(0, 0),
		ps.dungeon.CurrentLevel().Ground().NextLevelStaircase().Position,
		engine.SizeOf(80, 24),
		tcell.StyleDefault,
	)

	ps.dungeon.CurrentLevel().AddEntity(ps.player)
}

func PickUpItemUnderPlayer(eventLog *engine.GameEventLog, dungeon *model.Dungeon, player *model.Player) (success bool) {
	pos := player.Position()
	item := dungeon.CurrentLevel().RemoveItemAt(pos.XY())

	if item == nil {
		eventLog.Log("There is no item to pick up here")
		return false
	}

	success = player.Inventory().Push(item)

	if !success {
		eventLog.Log("Unable to pick up item")
		dungeon.CurrentLevel().SetItemAt(pos.X(), pos.Y(), item)
		return
	}

	if item.Named() != nil {
		itemName := item.Named().Name
		eventLog.Log("You picked up " + itemName)
	} else {
		eventLog.Log("You picked up an item")
	}

	return true
}

func (ps *PlayingState) OnTick(dt int64) (nextState GameState) {
	ps.nextGameState = ps

	ps.turnSystem.NextTurn()

	return ps.nextGameState
}

func (ps *PlayingState) CollectDrawables() []engine.Drawable {
	mainCameraDrawingInstructions := engine.CreateDrawingInstructions(func(v views.View) {
		visibilityMap := engine.ComputeFOV(
			func(x, y int) model.Tile {
				model.Map_MarkExplored(ps.dungeon.CurrentLevel().Ground(), x, y)

				return ps.dungeon.CurrentLevel().TileAt(x, y)
			},
			func(x, y int) bool { return model.Map_IsInBounds(ps.dungeon.CurrentLevel().Ground(), x, y) },
			func(x, y int) bool { return ps.dungeon.CurrentLevel().TileAt(x, y).Opaque() },
			ps.player.Position().X(), ps.player.Position().Y(),
			13,
		)

		ps.viewport.DrawFromProvider(v, func(x, y int) (rune, tcell.Style) {
			tile := visibilityMap[engine.PositionAt(x, y)]

			if tile != nil {

				if tile.Entity() != nil {
					return tile.Entity().Entity.Presentable().Rune, tile.Entity().Entity.Presentable().Style
				}

				if tile.Item() != nil {
					return tile.Item().Item.TileIcon(), tile.Item().Item.Style()
				}

				return tile.DefaultPresentation()
			}

			explored := model.Map_ExploredTileAt(ps.dungeon.CurrentLevel().Ground(), x, y)

			if explored != nil {
				return explored.DefaultPresentation()
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
