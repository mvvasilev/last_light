package npc

import (
	"mvvasilev/last_light/engine"

	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
)

type NPC interface {
	Name() string

	MovableEntity
}

type BasicNPC struct {
	id           uuid.UUID
	name         string
	presentation rune
	style        tcell.Style
	engine.Positioned
}

func CreateNPC(pos engine.Position, name string, presentation rune, style tcell.Style) *BasicNPC {
	return &BasicNPC{
		id:           uuid.New(),
		name:         name,
		presentation: presentation,
		style:        style,
		Positioned:   engine.WithPosition(pos),
	}
}

func (c *BasicNPC) Name() string {
	return c.name
}

func (c *BasicNPC) MoveTo(newPosition engine.Position) {
	c.Positioned.SetPosition(newPosition)
}

func (c *BasicNPC) UniqueId() uuid.UUID {
	return c.id
}

func (c *BasicNPC) Presentation() (rune, tcell.Style) {
	return c.presentation, c.style
}
