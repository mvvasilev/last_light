package state

import (
	"fmt"
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

	s.turnSystem.Schedule(10, func() (complete bool, requeue bool) {
		requeue = true
		complete = false

		if s.player.HealthData().IsDead {
			s.nextGameState = CreateGameOverState(inputSystem)
		}

		switch inputSystem.NextAction() {
		case systems.InputAction_PauseGame:
			s.nextGameState = PauseGame(s, s.turnSystem, s.inputSystem)
		case systems.InputAction_OpenInventory:
			s.nextGameState = CreateInventoryScreenState(s.eventLog, s.dungeon, s.inputSystem, s.turnSystem, s.player, s)
		case systems.InputAction_PickUpItem:
			complete = PickUpItemUnderPlayer(s.eventLog, s.dungeon, s.player)
		case systems.InputAction_Interact:
			complete = s.InteractBelowPlayer()
		case systems.InputAction_OpenLogs:
			s.viewShortLogs = !s.viewShortLogs
		case systems.InputAction_MovePlayer_East:
			complete = s.MovePlayer(model.East)
		case systems.InputAction_MovePlayer_West:
			complete = s.MovePlayer(model.West)
		case systems.InputAction_MovePlayer_North:
			complete = s.MovePlayer(model.North)
		case systems.InputAction_MovePlayer_South:
			complete = s.MovePlayer(model.South)
		default:
		}

		return
	})

	// s.someNPC = model.CreateEntity(
	// 	model.WithPosition(s.dungeon.CurrentLevel().Ground().NextLevelStaircase().Position),
	// 	model.WithName("NPC"),
	// 	model.WithPresentation('n', tcell.StyleDefault),
	// 	model.WithStats(model.RandomStats(21, 1, 20, []model.Stat{model.Stat_Attributes_Strength, model.Stat_Attributes_Constitution, model.Stat_Attributes_Intelligence, model.Stat_Attributes_Dexterity})),
	// 	model.WithHealthData(20, 20, false),
	// )

	// s.turnSystem.Schedule(20, func() (complete bool, requeue bool) {
	// 	s.CalcPathToPlayerAndMove()

	// 	return true, true
	// })

	s.eventLog = engine.CreateGameEventLog(100)

	s.uiEventLog = ui.CreateUIEventLog(0, 17, 80, 7, s.eventLog, tcell.StyleDefault)
	s.healthBar = ui.CreateHealthBar(68, 0, 12, 3, s.player, tcell.StyleDefault)

	s.dungeon.CurrentLevel().AddEntity(s.player)

	entityTable := model.CreateEntityTable()

	entityTable.Add(1, func(x, y int) model.Entity { return model.Entity_Imp(x, y) })

	s.npcs = SpawnNPCs(s.dungeon, 7, entityTable)

	for _, npc := range s.npcs {
		speed := 10

		if npc.Speed() != nil {
			speed = npc.Speed().Speed
		}

		s.turnSystem.Schedule(speed, func() (complete bool, requeue bool) {
			CalcPathToPlayerAndMove(25, s.eventLog, s.dungeon, npc, s.player)

			return true, true
		})
	}

	// s.dungeon.CurrentLevel().AddEntity(s.someNPC)

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

		ExecuteAttack(ps.eventLog, ps.player, ent)

		return true
	}

	if ps.dungeon.CurrentLevel().IsTilePassable(newPlayerPos.XY()) {
		ps.dungeon.CurrentLevel().MoveEntityTo(ps.player.UniqueId(), newPlayerPos.X(), newPlayerPos.Y())
		ps.viewport.SetCenter(ps.player.Position())

		ps.eventLog.Log("You moved " + model.DirectionName(direction))

		return true
	} else {
		ps.eventLog.Log("You bump into an impassable object" + model.DirectionName(direction))

		return false
	}
}

func ExecuteAttack(eventLog *engine.GameEventLog, attacker, victim model.Entity) {
	hit, precision, evasion, dmg, dmgType := CalculateAttack(attacker, victim)

	attackerName := "Unknown"

	if attacker.Named() != nil {
		attackerName = attacker.Named().Name
	}

	victimName := "Unknown"

	if victim.Named() != nil {
		victimName = victim.Named().Name
	}

	if !hit {
		eventLog.Log(fmt.Sprintf("%s attacked %s, but missed ( %v Evasion vs %v Precision)", attackerName, victimName, evasion, precision))
		return
	}

	victim.HealthData().Health -= dmg

	if victim.HealthData().Health <= 0 {
		victim.HealthData().IsDead = true
		eventLog.Log(fmt.Sprintf("%s attacked %s, and was victorious ( %v Evasion vs %v Precision)", attackerName, victimName, evasion, precision))
		return
	}

	eventLog.Log(fmt.Sprintf("%s attacked %s, and hit for %v %v damage", attackerName, victimName, dmg, model.DamageTypeName(dmgType)))
}

func CalculateAttack(attacker, victim model.Entity) (hit bool, precisionRoll, evasionRoll int, damage int, damageType model.DamageType) {
	if attacker.Equipped() != nil && attacker.Equipped().Inventory.AtSlot(model.EquippedSlotDominantHand) != nil {
		weapon := attacker.Equipped().Inventory.AtSlot(model.EquippedSlotDominantHand)

		return model.PhysicalWeaponAttack(attacker, weapon, victim)
	} else {
		return model.UnarmedAttack(attacker, victim)
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

func HasLineOfSight(dungeon *model.Dungeon, start, end engine.Position) bool {
	positions := engine.CastRay(start, end)

	for _, p := range positions {
		if dungeon.CurrentLevel().IsGroundTileOpaque(p.XY()) {
			return false
		}
	}

	return true
}

func WithinHitRange(pos engine.Position, otherPos engine.Position) bool {
	return pos.WithOffset(-1, 0) == otherPos || pos.WithOffset(+1, 0) == otherPos || pos.WithOffset(0, -1) == otherPos || pos.WithOffset(0, +1) == otherPos
}

func CalcPathToPlayerAndMove(simulationDistance int, eventLog *engine.GameEventLog, dungeon *model.Dungeon, npc model.Entity, player *model.Player) {
	if npc.Positioned().Position.DistanceSquared(player.Position()) > simulationDistance*simulationDistance {
		return
	}

	if npc.HealthData().IsDead {
		dungeon.CurrentLevel().DropEntity(npc.UniqueId())
		return
	}

	playerVisibleAndInRange := false

	if npc.Positioned().Position.DistanceSquared(player.Position()) < 144 && HasLineOfSight(dungeon, npc.Positioned().Position, player.Position()) {
		playerVisibleAndInRange = true
	}

	if !playerVisibleAndInRange {
		randomMove := model.Direction(engine.RandInt(int(model.DirectionNone), int(model.East)))

		nextPos := npc.Positioned().Position

		switch randomMove {
		case model.North:
			nextPos = nextPos.WithOffset(0, -1)
		case model.South:
			nextPos = nextPos.WithOffset(0, +1)
		case model.West:
			nextPos = nextPos.WithOffset(-1, 0)
		case model.East:
			nextPos = nextPos.WithOffset(+1, 0)
		default:
			return
		}

		if dungeon.CurrentLevel().IsTilePassable(nextPos.XY()) {
			dungeon.CurrentLevel().MoveEntityTo(
				npc.UniqueId(),
				nextPos.X(),
				nextPos.Y(),
			)
		}

		return
	}

	if WithinHitRange(npc.Positioned().Position, player.Position()) {
		ExecuteAttack(eventLog, npc, player)
	}

	pathToPlayer := engine.FindPath(
		npc.Positioned().Position,
		player.Position(),
		12,
		func(x, y int) bool {
			if x == player.Position().X() && y == player.Position().Y() {
				return true
			}

			return dungeon.CurrentLevel().IsTilePassable(x, y)
		},
	)

	if pathToPlayer == nil {
		return
	}

	nextPos, hasNext := pathToPlayer.Next()

	if !hasNext {
		return
	}

	if nextPos.Equals(player.Position()) {
		return
	}

	dungeon.CurrentLevel().MoveEntityTo(npc.UniqueId(), nextPos.X(), nextPos.Y())
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
