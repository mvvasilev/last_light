package model

import (
	"mvvasilev/last_light/util"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	"github.com/google/uuid"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Entity interface {
	UniqueId() uuid.UUID
	Draw(v views.View)
	Input(e *tcell.EventKey)
	Tick(dt int64)
}

type MovableEntity interface {
	Position() util.Position
	Move(dir Direction)
}
