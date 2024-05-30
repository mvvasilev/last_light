package ui

import (
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/input"
)

type UIElement interface {
	MoveTo(x, y int)
	Position() engine.Position
	Size() engine.Size
	Input(inputAction input.InputAction)

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
