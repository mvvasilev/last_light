package state

import (
	"mvvasilev/last_light/game/model"
	"mvvasilev/last_light/render"
	"mvvasilev/last_light/util"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
)

type PlayingState struct {
	player *model.Player
	level  *model.MultilevelMap

	viewport *render.Viewport

	movePlayerDirection model.Direction
	pauseGame           bool
	openInventory       bool
}

func BeginPlayingState() *PlayingState {
	s := new(PlayingState)

	mapSize := util.SizeOf(128, 128)

	s.player = model.CreatePlayer(40, 12)

	s.level = model.CreateMultilevelMap(
		model.CreateFlatGroundDungeonLevel(mapSize.WH()),
		model.CreateEmptyDungeonLevel(mapSize.WH()),
	)

	s.level.SetTileAtHeight(40, 12, 1, s.player)

	s.viewport = render.CreateViewport(
		util.PositionAt(0, 0),
		util.PositionAt(40, 12),
		util.SizeOf(80, 24),
		tcell.StyleDefault,
	)

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

	tileAtMovePos := ps.level.TileAt(newPlayerPos.XY())

	if tileAtMovePos.Passable() {
		ps.level.SetTileAtHeight(ps.player.Position().X(), ps.player.Position().Y(), 1, nil)
		ps.player.Move(ps.movePlayerDirection)
		ps.viewport.SetCenter(ps.player.Position())
		ps.level.SetTileAtHeight(ps.player.Position().X(), ps.player.Position().Y(), 1, ps.player)
	}

	ps.movePlayerDirection = model.DirectionNone
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

	switch e.Key() {
	case tcell.KeyUp:
		ps.movePlayerDirection = model.DirectionUp
	case tcell.KeyDown:
		ps.movePlayerDirection = model.DirectionDown
	case tcell.KeyLeft:
		ps.movePlayerDirection = model.DirectionLeft
	case tcell.KeyRight:
		ps.movePlayerDirection = model.DirectionRight
	case tcell.KeyRune:
		switch e.Rune() {
		case 'w':
			ps.movePlayerDirection = model.DirectionUp
		case 'a':
			ps.movePlayerDirection = model.DirectionLeft
		case 's':
			ps.movePlayerDirection = model.DirectionDown
		case 'd':
			ps.movePlayerDirection = model.DirectionRight
		}
	}
}

func (ps *PlayingState) OnTick(dt int64) GameState {
	ps.player.Tick(dt)

	if ps.pauseGame {
		return PauseGame(ps)
	}

	if ps.openInventory {
		return CreateInventoryScreenState(ps.player, ps)
	}

	if ps.movePlayerDirection != model.DirectionNone {
		ps.MovePlayer()
	}

	return ps
}

func (ps *PlayingState) CollectDrawables() []render.Drawable {
	return render.Multidraw(render.CreateDrawingInstructions(func(v views.View) {
		ps.viewport.DrawFromProvider(v, func(x, y int) rune {
			tile := ps.level.TileAt(x, y)

			if tile != nil {
				return tile.Presentation()
			}

			return ' '
		})
	}))
}
