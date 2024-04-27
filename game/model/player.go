package model

import (
	"mvvasilev/last_light/util"

	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
)

type Player struct {
	id       uuid.UUID
	position util.Position
}

func CreatePlayer(x, y int) *Player {
	p := new(Player)

	p.id = uuid.New()
	p.position = util.PositionAt(x, y)

	return p
}

func (p *Player) UniqueId() uuid.UUID {
	return p.id
}

func (p *Player) Position() util.Position {
	return p.position
}

func (p *Player) Move(dir Direction) {
	p.position = p.Position().WithOffset(MovementDirectionOffset(dir))
}

func (p *Player) Presentation() rune {
	return '@'
}

func (p *Player) Passable() bool {
	return false
}

func (p *Player) Input(e *tcell.EventKey) {
}

func (p *Player) Tick(dt int64) {

}
