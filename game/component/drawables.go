package component

import (
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/engine/ecs"
)

const ComponentType_RenderableComponent = ecs.ComponentType_0

type DrawablesComponent struct {
	Priority  int
	Drawables []engine.Drawable
}

func (rc *DrawablesComponent) Type() ecs.ComponentType {
	return ComponentType_RenderableComponent
}
