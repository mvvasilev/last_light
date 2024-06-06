package engine

import (
	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	"github.com/google/uuid"
)

type Raw struct {
	id       uuid.UUID
	buffer   [][]rune
	position Position
	style    tcell.Style
}

func CreateRawDrawable(x, y int, style tcell.Style, buffer ...string) *Raw {
	r := new(Raw)

	r.position = PositionAt(x, y)
	r.buffer = make([][]rune, 0)

	for _, row := range buffer {
		r.buffer = append(r.buffer, []rune(row))
	}

	r.style = style

	return r
}

func CreateRawDrawableFromBuffer(x, y int, style tcell.Style, buffer [][]rune) *Raw {
	r := new(Raw)

	r.position = PositionAt(x, y)
	r.buffer = buffer

	return r
}

func (r *Raw) UniqueId() uuid.UUID {
	return r.id
}

func (r *Raw) DrawWithin(screenX, screenY, originX, originY, width, height int, v views.View) {
	for h := originY; h < originY+height; h++ {

		if h < 0 || h >= len(r.buffer) {
			screenY += 1
			continue
		}

		for w := originX; w < originX+width; w++ {
			if w < 0 || w >= len(r.buffer[h]) {
				screenX += 1
				continue
			}

			v.SetContent(screenX, screenY, r.buffer[h][w], nil, r.style)

			screenX += 1
		}

		screenX = 0
		screenY += 1
	}
}

func (r *Raw) Draw(v views.View) {
	x := r.position.X()
	y := r.position.Y()

	for h, row := range r.buffer {
		for i, ru := range row {
			v.SetContent(x+i, y+h, ru, nil, r.style)
		}
	}
}
