package npc

import (
	"mvvasilev/last_light/engine"

	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
)

type BasicNPC struct {
	id uuid.UUID
	engine.Positioned
}

func CreateNPC(pos engine.Position) *BasicNPC {
	return &BasicNPC{
		id:         uuid.New(),
		Positioned: engine.WithPosition(pos),
	}
}

func (c *BasicNPC) MoveTo(newPosition engine.Position) {
	c.Positioned.SetPosition(newPosition)
}

func (c *BasicNPC) UniqueId() uuid.UUID {
	return c.id
}

func (c *BasicNPC) Input(e *tcell.EventKey) {
}

func (c *BasicNPC) Tick(dt int64) {
}
