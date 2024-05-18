package system

import (
	"fmt"
	"math"
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/engine/ecs"
	"mvvasilev/last_light/game/component"
	"slices"

	"github.com/gdamore/tcell/v2"
)

type RenderingSystem struct {
	engineContext *engine.EngineContext
}

func CreateRenderingSystem(renderContext *engine.EngineContext) *RenderingSystem {
	return &RenderingSystem{
		engineContext: renderContext,
	}
}

func (rs *RenderingSystem) Name() string {
	return "RenderingSystem"
}

func (rs *RenderingSystem) Order() int {
	return math.MaxInt
}

func (rs *RenderingSystem) Tick(world *ecs.World, deltaTime int64) {
	comps, err := world.QueryComponents(component.ComponentType_RenderableComponent)

	if err != nil {
		// Skip this frame since an error occured // TODO: error logging
		return
	}

	components := comps[component.ComponentType_RenderableComponent]

	slices.SortFunc(components, func(a ecs.Component, b ecs.Component) int {
		aDrawable := a.(*component.DrawablesComponent)
		bDrawable := b.(*component.DrawablesComponent)

		return aDrawable.Priority - bDrawable.Priority
	})

	fps := 1_000 / deltaTime

	msPerFrame := float32(fps) / 1000.0

	fpsText := engine.CreateText(0, 0, 16, 1, fmt.Sprintf("%vms", msPerFrame), tcell.StyleDefault)

	rs.engineContext.Clear()

	for _, c := range components {
		drawables := c.(*component.DrawablesComponent).Drawables
		rs.engineContext.Draw(drawables)
	}

	rs.engineContext.Draw([]engine.Drawable{fpsText})

	rs.engineContext.Show()
}
