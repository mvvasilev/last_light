package model

import (
	"mvvasilev/last_light/util"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	"github.com/google/uuid"
)

type Player struct {
	id       uuid.UUID
	position util.Position
	style    tcell.Style
}

func CreatePlayer(x, y uint16, style tcell.Style) *Player {
	p := new(Player)

	p.id = uuid.New()
	p.position = util.PositionAt(x, y)

	return p
}

func (p *Player) UniqueId() uuid.UUID {
	return p.id
}

func (p *Player) Move(dir Direction) {
	x, y := p.position.XYUint16()

	switch dir {
	case Up:
		p.position = util.PositionAt(x, y-1)
	case Down:
		p.position = util.PositionAt(x, y+1)
	case Left:
		p.position = util.PositionAt(x-1, y)
	case Right:
		p.position = util.PositionAt(x+1, y)
	}
}

func (p *Player) Draw(v views.View) {
	x, y := p.position.XY()
	v.SetContent(x, y, '@', nil, p.style)
}

func (p *Player) Input(e *tcell.EventKey) {
	switch e.Key() {
	case tcell.KeyUp:
		p.Move(Up)
	case tcell.KeyDown:
		p.Move(Down)
	case tcell.KeyLeft:
		p.Move(Left)
	case tcell.KeyRight:
		p.Move(Right)
	case tcell.KeyRune:
		switch e.Rune() {
		case 'w':
			p.Move(Up)
		case 'a':
			p.Move(Left)
		case 's':
			p.Move(Down)
		case 'd':
			p.Move(Right)
		}
	}
}

func (p *Player) Tick(dt int64) {

}
