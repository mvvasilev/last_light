package model

import (
	"mvvasilev/last_light/engine"

	"github.com/gdamore/tcell/v2"
)

type Player struct {
	Entity
}

func CreatePlayer(x, y int, playerBaseStats map[Stat]int) *Player {
	p := &Player{
		Entity: CreateEntity(
			WithName("Player"),
			WithPosition(engine.PositionAt(x, y)),
			WithPresentation('@', tcell.StyleDefault),
			WithInventory(CreateEquippedInventory()),
			WithStats(playerBaseStats),
			WithHealthData(0, 0, false),
			WithSpeed(10),
		),
	}

	p.HealthData().MaxHealth = BaseMaxHealth(p)
	p.HealthData().Health = p.HealthData().MaxHealth

	return p
}

func (p *Player) Inventory() *EquippedInventory {
	return p.Entity.Equipped().Inventory
}

func (p *Player) Position() engine.Position {
	return p.Entity.Positioned().Position
}

func (p *Player) Presentation() (rune, tcell.Style) {
	return p.Presentable().Rune, p.Presentable().Style
}

func (p *Player) Stats() *Entity_StatsHolderComponent {
	return p.Entity.Stats()
}

func (p *Player) HealthData() *Entity_HealthComponent {
	return p.Entity.HealthData()
}
