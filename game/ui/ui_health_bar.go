package ui

import (
	"fmt"
	"math"
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/model"
	"mvvasilev/last_light/game/systems"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	"github.com/google/uuid"
)

type UIHealthBar struct {
	id     uuid.UUID
	player *model.Player_V2

	window *UIWindow

	style tcell.Style
}

func CreateHealthBar(x, y, w, h int, player *model.Player_V2, style tcell.Style) *UIHealthBar {
	return &UIHealthBar{
		window: CreateWindow(x, y, w, h, "HP", style),
		player: player,
		style:  style,
	}
}

func (uihp *UIHealthBar) MoveTo(x int, y int) {
	uihp.window.MoveTo(x, y)
}

func (uihp *UIHealthBar) Position() engine.Position {
	return uihp.window.Position()
}

func (uihp *UIHealthBar) Size() engine.Size {
	return uihp.window.Size()
}

func (uihp *UIHealthBar) Input(inputAction systems.InputAction) {
}

func (uihp *UIHealthBar) UniqueId() uuid.UUID {
	return uihp.id
}

func (uihp *UIHealthBar) Draw(v views.View) {
	x, y, w, h := uihp.Position().X(), uihp.Position().Y(), uihp.Size().Width(), uihp.Size().Height()

	uihp.window.Draw(v)

	stages := []rune{'█', '▓', '▒', '░'} // 0 = 1.0, 1 = 0.75, 2 = 0.5, 3 = 0.25

	percentage := (float64(w) - 2.0) * (float64(uihp.player.HealthData().Health) / float64(uihp.player.HealthData().MaxHealth))

	whole := math.Trunc(percentage)
	last := percentage - whole

	hpStyle := tcell.StyleDefault.Foreground(tcell.ColorIndianRed)

	for i := range int(whole) {
		v.SetContent(x+1+i, y+1, stages[0], nil, hpStyle)
	}

	if last > 0.0 {
		if last <= 0.25 {
			v.SetContent(x+1+int(whole), y+1, stages[3], nil, hpStyle)
		}

		if last <= 0.50 {
			v.SetContent(x+1+int(whole), y+1, stages[2], nil, hpStyle)
		}

		if last <= 0.75 {
			v.SetContent(x+1+int(whole), y+1, stages[1], nil, hpStyle)
		}
	}

	hpText := fmt.Sprintf("%v/%v", uihp.player.HealthData().Health, uihp.player.HealthData().MaxHealth)

	engine.DrawText(
		x+w/2-len(hpText)/2,
		y+h-1,
		hpText,
		hpStyle,
		v,
	)
}
