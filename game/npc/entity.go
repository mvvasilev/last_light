package npc

import (
	"mvvasilev/last_light/engine"

	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
)

type Direction int

const (
	DirectionNone Direction = iota
	North
	South
	West
	East
)

func DirectionName(dir Direction) string {
	switch dir {
	case North:
		return "North"
	case South:
		return "South"
	case West:
		return "West"
	case East:
		return "East"
	default:
		return "Unknown"
	}
}

func MovementDirectionOffset(dir Direction) (int, int) {
	switch dir {
	case North:
		return 0, -1
	case South:
		return 0, 1
	case West:
		return -1, 0
	case East:
		return 1, 0
	}

	return 0, 0
}

type Entity interface {
	UniqueId() uuid.UUID
	Input(e *tcell.EventKey)
	Tick(dt int64)
}

type MovableEntity interface {
	Position() engine.Position
	MoveTo(newPosition engine.Position)

	Entity
}
