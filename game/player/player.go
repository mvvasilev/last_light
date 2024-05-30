package player

import (
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/item"
	"mvvasilev/last_light/game/rpg"

	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
)

type Player struct {
	id       uuid.UUID
	position engine.Position

	inventory *item.EquippedInventory

	*rpg.BasicRPGEntity
}

func CreatePlayer(x, y int) *Player {
	p := new(Player)

	p.id = uuid.New()
	p.position = engine.PositionAt(x, y)
	p.inventory = item.CreateEquippedInventory()
	p.BasicRPGEntity = rpg.CreateBasicRPGEntity(
		map[rpg.Stat]int{
			rpg.Stat_Attributes_Constitution: 10,
			rpg.Stat_Attributes_Dexterity:    10,
			rpg.Stat_Attributes_Strength:     10,
			rpg.Stat_Attributes_Intelligence: 10,
		},
		map[rpg.Stat][]rpg.StatModifier{},
	)

	return p
}

func (p *Player) UniqueId() uuid.UUID {
	return p.id
}

func (p *Player) Position() engine.Position {
	return p.position
}

func (p *Player) MoveTo(newPos engine.Position) {
	p.position = newPos
}

func (p *Player) Presentation() (rune, tcell.Style) {
	return '@', tcell.StyleDefault
}

func (p *Player) Passable() bool {
	return false
}

func (p *Player) Transparent() bool {
	return false
}

func (p *Player) Inventory() *item.EquippedInventory {
	return p.inventory
}

func (p *Player) Input(e *tcell.EventKey) {
}

func (p *Player) Tick(dt int64) {

}
