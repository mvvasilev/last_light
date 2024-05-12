package logic

import (
	"mvvasilev/last_light/game/model"

	"github.com/gdamore/tcell/v2"
)

type LogicSnippet[T model.Entity] interface {
	Input(e *tcell.EventKey)
	Tick(dt int64, entity T)
}
