package render

import (
	"mvvasilev/last_light/util"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	"github.com/google/uuid"
)

type Raw struct {
	id       uuid.UUID
	buffer   [][]rune
	position util.Position
	style    tcell.Style
}

func CreateRawDrawable(x, y uint16, style tcell.Style, buffer ...string) *Raw {
	r := new(Raw)

	r.position = util.PositionAt(x, y)
	r.buffer = make([][]rune, 0)

	for _, row := range buffer {
		r.buffer = append(r.buffer, []rune(row))
	}

	r.style = style

	return r
}

func (r *Raw) UniqueId() uuid.UUID {
	return r.id
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
