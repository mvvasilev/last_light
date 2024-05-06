package model

import (
	"mvvasilev/last_light/util"

	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
)

type Direction int

const (
	DirectionNone Direction = iota
	DirectionUp
	DirectionDown
	DirectionLeft
	DirectionRight
)

func MovementDirectionOffset(dir Direction) (int, int) {
	switch dir {
	case DirectionUp:
		return 0, -1
	case DirectionDown:
		return 0, 1
	case DirectionLeft:
		return -1, 0
	case DirectionRight:
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
	Position() util.Position
	MoveTo(newPosition util.Position)

	Entity
}
