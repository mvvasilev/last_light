package component

import (
	"mvvasilev/last_light/engine/ecs"

	"github.com/gdamore/tcell/v2"
)

const ComponentType_InputComponent = ecs.ComponentType_1

type InputComponent struct {
	KeyEvents []*tcell.EventKey
}

func (ic *InputComponent) Type() ecs.ComponentType {
	return ComponentType_InputComponent
}
