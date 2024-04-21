package ui

import (
	"mvvasilev/last_light/util"

	"github.com/google/uuid"
)

type UIElement interface {
	UniqueId() uuid.UUID
	Position() util.Position
}
