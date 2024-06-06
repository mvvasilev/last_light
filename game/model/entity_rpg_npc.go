package model

// import (
// 	"mvvasilev/last_light/engine"
// 	"mvvasilev/last_light/game/item"
// 	"mvvasilev/last_light/game/rpg"

// 	"github.com/gdamore/tcell/v2"
// )

// type RPGNPC interface {
// 	NPC
// 	rpg.RPGEntity
// 	EquippedEntity
// }

// type BasicRPGNPC struct {
// 	inventory *item.EquippedInventory

// 	*BasicNPC
// 	*rpg.BasicRPGEntity
// }

// func CreateRPGNPC(x, y int, name string, representation rune, style tcell.Style, stats map[rpg.Stat]int) *BasicRPGNPC {
// 	rpgnpc := &BasicRPGNPC{
// 		inventory: item.CreateEquippedInventory(),
// 		BasicNPC: CreateNPC(
// 			engine.PositionAt(x, y),
// 			name,
// 			representation,
// 			style,
// 		),
// 		BasicRPGEntity: rpg.CreateBasicRPGEntity(
// 			stats,
// 			map[rpg.Stat][]rpg.StatModifier{},
// 		),
// 	}

// 	rpgnpc.Heal(rpg.BaseMaxHealth(rpgnpc))

// 	return rpgnpc
// }

// func (rnpc *BasicRPGNPC) Inventory() *item.EquippedInventory {
// 	return rnpc.inventory
// }

// func (p *BasicRPGNPC) CalculateAttack(other rpg.RPGEntity) (hit bool, precisionRoll, evasionRoll int, damage int, damageType rpg.DamageType) {
// 	mainHand := p.inventory.AtSlot(item.EquippedSlotDominantHand)

// 	switch mh := mainHand.(type) {
// 	case rpg.RPGItem:
// 		return rpg.PhysicalWeaponAttack(p, mh, other)
// 	default:
// 		return rpg.UnarmedAttack(p, other)
// 	}
// }
