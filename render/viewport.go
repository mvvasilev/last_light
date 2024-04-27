package render

import (
	"mvvasilev/last_light/util"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	"github.com/google/uuid"
)

type Viewport struct {
	id             uuid.UUID
	screenLocation util.Position

	viewportCenter util.Position
	viewportSize   util.Size

	style tcell.Style
}

func CreateViewport(screenLoc, viewportCenter util.Position, size util.Size, style tcell.Style) *Viewport {
	v := new(Viewport)

	v.id = uuid.New()
	v.screenLocation = screenLoc
	v.viewportCenter = viewportCenter
	v.viewportSize = size
	v.style = style

	return v
}

func (vp *Viewport) UniqueId() uuid.UUID {
	return vp.id
}

func (vp *Viewport) Center() util.Position {
	return vp.viewportCenter
}

func (vp *Viewport) SetCenter(pos util.Position) {
	vp.viewportCenter = pos
}

func (vp *Viewport) Size() util.Size {
	return vp.viewportSize
}

func (vp *Viewport) ScreenLocation() util.Position {
	return vp.screenLocation
}

func (vp *Viewport) DrawFromProvider(v views.View, provider func(x, y int) rune) {
	width, height := vp.viewportSize.WH()
	originX, originY := vp.viewportCenter.WithOffset(-width/2, -height/2).XY()
	screenX, screenY := vp.screenLocation.XY()

	for h := originY; h < originY+height; h++ {
		for w := originX; w < originX+width; w++ {
			v.SetContent(screenX, screenY, provider(w, h), nil, vp.style)

			screenX += 1
		}

		screenX = 0
		screenY += 1
	}
}

func (vp *Viewport) Draw(v views.View, buffer [][]rune) {
	width, height := vp.viewportSize.WH()
	originX, originY := vp.viewportCenter.WithOffset(-width/2, -height/2).XY()
	screenX, screenY := vp.screenLocation.XY()

	for h := originY; h < originY+height; h++ {

		if h < 0 || h >= len(buffer) {
			screenY += 1
			continue
		}

		for w := originX; w < originX+width; w++ {
			if w < 0 || w >= len(buffer[h]) {
				screenX += 1
				continue
			}

			v.SetContent(screenX, screenY, buffer[h][w], nil, vp.style)

			screenX += 1
		}

		screenX = 0
		screenY += 1
	}
}
