package model

import (
	"mvvasilev/last_light/engine"

	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
)

type NPC struct {
	id uuid.UUID
	engine.Positioned
}

func CreateNPC(pos engine.Position) *NPC {
	return &NPC{
		id:         uuid.New(),
		Positioned: engine.WithPosition(pos),
	}
}

func (c *NPC) MoveTo(newPosition engine.Position) {
	c.Positioned.SetPosition(newPosition)
}

func (c *NPC) UniqueId() uuid.UUID {
	return c.id
}

func (c *NPC) Input(e *tcell.EventKey) {
}

func (c *NPC) Tick(dt int64) {
}
