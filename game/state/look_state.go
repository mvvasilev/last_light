package state

import (
	"fmt"
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/model"
	"mvvasilev/last_light/game/systems"
	"mvvasilev/last_light/game/ui"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
)

const CursorRune = '+'
const CursorBlinkTime = 200 // Blink cursor every 200ms, showing what's under it

type LookState struct {
	prevState GameState

	inputSystem *systems.InputSystem
	turnSystem  *systems.TurnSystem
	eventLog    *engine.GameEventLog
	player      *model.Player
	dungeon     *model.Dungeon

	showLog    bool
	uiEventLog *ui.UIEventLog

	showCursor          bool
	cursorPos           engine.Position
	lastCursorBlinkTime time.Time

	nextGameState GameState
}

func CreateLookState(prevState GameState, eventLog *engine.GameEventLog, dungeon *model.Dungeon, inputSystem *systems.InputSystem, turnSystem *systems.TurnSystem, player *model.Player) *LookState {
	return &LookState{
		prevState:           prevState,
		inputSystem:         inputSystem,
		turnSystem:          turnSystem,
		dungeon:             dungeon,
		player:              player,
		eventLog:            eventLog,
		cursorPos:           engine.PositionAt(0, 0),
		lastCursorBlinkTime: time.Now(),
		showLog:             true,
		uiEventLog:          ui.CreateUIEventLog(0, 17, 80, 7, eventLog, tcell.StyleDefault),
	}
}

func (ls *LookState) InputContext() systems.InputContext {
	return systems.InputContext_Look
}

func (ls *LookState) OnTick(dt int64) GameState {
	ls.nextGameState = ls

	switch ls.inputSystem.NextAction() {
	case systems.InputAction_Move_North:
		ls.cursorPos = ls.cursorPos.WithOffset(model.MovementDirectionOffset(model.North))
	case systems.InputAction_Move_South:
		ls.cursorPos = ls.cursorPos.WithOffset(model.MovementDirectionOffset(model.South))
	case systems.InputAction_Move_East:
		ls.cursorPos = ls.cursorPos.WithOffset(model.MovementDirectionOffset(model.East))
	case systems.InputAction_Move_West:
		ls.cursorPos = ls.cursorPos.WithOffset(model.MovementDirectionOffset(model.West))
	case systems.InputAction_OpenLogs:
		ls.showLog = !ls.showLog
	case systems.InputAction_Describe:
		ls.Describe()
	case systems.InputAction_Shoot:
		ls.ShootEquippedWeapon()
	case systems.InputAction_Menu_Exit:
		ls.nextGameState = ls.prevState
	}

	return ls.nextGameState
}

func (ls *LookState) ShootEquippedWeapon() {
	weapon := ls.player.Inventory().AtSlot(model.EquippedSlotDominantHand)

	if weapon == nil {
		ls.eventLog.Log("You don't have anything equipped!")

		return
	}

	if weapon.Damaging() == nil {
		ls.eventLog.Log("Item unusable")

		return
	}

	damaging := weapon.Damaging()

	if !damaging.IsRanged {
		ls.eventLog.Log("Equipped weapon is not ranged!")

		return
	}

	// TODO: Projectiles
	dX, dY := ls.lookCursorCoordsToDungeonCoords()

	distance := engine.PositionAt(dX, dY).Distance(ls.player.Position())

	if distance >= 12 {
		ls.eventLog.Log("Can't see in the dark that far")

		return
	}

	path := engine.LinePath(
		ls.player.Position(),
		engine.PositionAt(dX, dY),
	)

	if path == nil {
		ls.eventLog.Log("Can't shoot there, something is in the way")
		return
	}

	projectile := model.Entity_ArrowProjectile(ls.player, path, ls.eventLog, ls.dungeon)

	ls.turnSystem.Schedule(
		projectile.Behavior().Speed,
		projectile.Behavior().Behavior,
	)

	ls.player.SkipNextTurn(true)

	ls.nextGameState = ls.prevState
}

func (ls *LookState) Describe() {
	dX, dY := ls.lookCursorCoordsToDungeonCoords()

	distance := engine.PositionAt(dX, dY).Distance(ls.player.Position())

	if distance >= 12 {
		ls.eventLog.Log("Can't see in the dark that far")

		return
	}

	isVisibleFromPlayer, lastTile := model.HasLineOfSight(ls.dungeon, ls.player.Position(), engine.PositionAt(dX, dY))

	if !isVisibleFromPlayer {
		materialName, _ := materialToDescription(lastTile.Material())

		ls.eventLog.Log(fmt.Sprintf("%s obscures your view", materialName))
		return
	}

	tile := ls.dungeon.CurrentLevel().TileAt(dX, dY)

	entity := tile.Entity()

	if entity != nil {
		ls.DescribeEntity(entity.Entity)

		return
	}

	item := tile.Item()

	if item != nil {
		ls.DescribeItem(item.Item)

		return
	}

	materialName, materialDesc := materialToDescription(tile.Material())

	ls.eventLog.Log(fmt.Sprintf("%s: %s", materialName, materialDesc))
}

func (ls *LookState) DescribeEntity(entity model.Entity) {
	if entity == nil {
		return
	}

	if entity == ls.player {
		ls.eventLog.Log("You")

		return
	}

	if entity.Named() == nil {
		return
	}

	if entity.Described() != nil {
		ls.eventLog.Log(fmt.Sprintf("%s: %s", entity.Named().Name, entity.Described().Description))
	} else {
		ls.eventLog.Log(entity.Named().Name)
	}
}

func (ls *LookState) DescribeItem(item model.Item) {
	if item == nil {
		return
	}

	if item.Named() == nil {
		return
	}

	if item.Described() != nil {
		ls.eventLog.Log(fmt.Sprintf("%s: %s", item.Named().Name, item.Described().Description))
	} else {
		ls.eventLog.Log(item.Named().Name)
	}
}

func materialToDescription(material model.Material) (name, description string) {
	switch material {
	case model.MaterialVoid:
		return "Void", "Who knows what lurks here..."
	case model.MaterialWall:
		return "Wall", "Mediocre masonry"
	case model.MaterialGround:
		return "Ground", "Try not to trip"
	}

	return "Void", "Who knows what lurks here..."
}

func (ls *LookState) lookCursorCoordsToScreenCoords() (sX, xY int) {
	x, y := ls.cursorPos.XY()
	middleOfScreenX, middleOfScreenY := engine.TERMINAL_SIZE_WIDTH/2, engine.TERMINAL_SIZE_HEIGHT/2
	return middleOfScreenX + x, middleOfScreenY + y
}

func (ls *LookState) lookCursorCoordsToDungeonCoords() (sX, xY int) {
	x, y := ls.cursorPos.XY()
	playerX, playerY := ls.player.Position().XY()
	return playerX + x, playerY + y
}

func (ls *LookState) CollectDrawables() []engine.Drawable {
	drawables := append(ls.prevState.CollectDrawables(), engine.CreateDrawingInstructions(func(v views.View) {
		if time.Since(ls.lastCursorBlinkTime).Milliseconds() >= CursorBlinkTime {
			ls.showCursor = !ls.showCursor
			ls.lastCursorBlinkTime = time.Now()
		}

		if ls.showCursor {
			x, y := ls.lookCursorCoordsToScreenCoords()
			v.SetContent(x, y, CursorRune, nil, tcell.StyleDefault)
		}
	}))

	if ls.showLog {
		drawables = append(drawables, ls.uiEventLog)
	}

	return drawables
}
