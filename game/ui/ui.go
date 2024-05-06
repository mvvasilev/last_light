package ui

import (
	"mvvasilev/last_light/engine"
	engine1 "mvvasilev/last_light/engine"

	"github.com/gdamore/tcell/v2"
)

type UIElement interface {
	MoveTo(x, y int)
	Position() engine1.Position
	Size() engine1.Size
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
