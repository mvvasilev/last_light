package render

import (
	"github.com/gdamore/tcell/v2/views"
	"github.com/google/uuid"
)

type ArbitraryDrawable struct {
	id                  uuid.UUID
	drawingInstructions func(v views.View)
}

func CreateDrawingInstructions(instructions func(v views.View)) *ArbitraryDrawable {
	a := new(ArbitraryDrawable)

	a.id = uuid.New()
	a.drawingInstructions = instructions

	return a
}

func (ab *ArbitraryDrawable) UniqueId() uuid.UUID {
	return ab.id
}

func (ab *ArbitraryDrawable) Draw(v views.View) {
	ab.drawingInstructions(v)
}
