package model

import (
	"mvvasilev/last_light/engine"

	"github.com/gdamore/tcell/v2"
)

type Player_V2 struct {
	Entity_V2
}

func CreatePlayer_V2(x, y int, playerBaseStats map[Stat]int) *Player_V2 {
	p := &Player_V2{
		Entity_V2: CreateEntity(
			WithName("Player"),
			WithPosition(engine.PositionAt(x, y)),
			WithPresentation('@', tcell.StyleDefault),
			WithInventory(CreateEquippedInventory()),
			WithStats(playerBaseStats),
			WithHealthData(0, 0, false),
		),
	}

	p.HealthData().MaxHealth = BaseMaxHealth(p)
	p.HealthData().Health = p.HealthData().MaxHealth

	return p
}

func (p *Player_V2) Inventory() *EquippedInventory {
	return p.Entity_V2.Equipped().Inventory
}

func (p *Player_V2) Position() engine.Position {
	return p.Entity_V2.Positioned().Position
}

func (p *Player_V2) Presentation() (rune, tcell.Style) {
	return p.Presentable().Rune, p.Presentable().Style
}

func (p *Player_V2) Stats() *Entity_StatsHolderComponent {
	return p.Entity_V2.Stats()
}

func (p *Player_V2) HealthData() *Entity_HealthComponent {
	return p.Entity_V2.HealthData()
}
