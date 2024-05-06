package ui

import (
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/util"

	"github.com/gdamore/tcell/v2"
)

type UIElement interface {
	MoveTo(x, y int)
	Position() util.Position
	Size() util.Size
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
	UIHighlightableElement
}
