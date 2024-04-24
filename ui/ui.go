package ui

import (
	"mvvasilev/last_light/util"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	"github.com/google/uuid"
)

type UIElement interface {
	UniqueId() uuid.UUID
	MoveTo(x, y uint16)
	Position() util.Position
	Size() util.Size
	Draw(v views.View)
	Input(e *tcell.EventKey)
}

type UIHighlightableElement interface {
	IsHighlighted() bool
	Highlight()
	Unhighlight()
	SetHighlighted(highlighted bool)
	UIElement
}

type UISelectableElement interface {
	Select()
	OnSelect(func())
	UIHighlightableElement
}
