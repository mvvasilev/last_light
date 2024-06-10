package model

import (
	"mvvasilev/last_light/engine"

	"github.com/gdamore/tcell/v2"
)

type Player struct {
	Entity

	inLookState  bool
	skipNextTurn bool
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
			WithBehavior(100, nil),
		),
	}

	p.Inventory().Push(Item_Bow())

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

func (p *Player) DefaultSpeed() *Entity_BehaviorComponent {
	return p.Entity.Behavior()
}

func (p *Player) SkipNextTurn(skip bool) {
	p.skipNextTurn = skip
}

func (p *Player) IsNextTurnSkipped() bool {
	return p.skipNextTurn
}

func (p *Player) IsInLookState() bool {
	return p.inLookState
}

func (p *Player) SetInLookState(lookState bool) {
	p.inLookState = lookState
}
