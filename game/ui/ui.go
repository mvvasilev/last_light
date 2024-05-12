package ui

import (
	"mvvasilev/last_light/engine"

	"github.com/gdamore/tcell/v2"
)

type UIElement interface {
	MoveTo(x, y int)
	Position() engine.Position
	Size() engine.Size
	Input(e *tcell.EventKey)

	engine.Drawable
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
	UIElement
}
