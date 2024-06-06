package model

import (
	"mvvasilev/last_light/engine"

	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
)

type Direction int

const (
	DirectionNone Direction = iota
	North
	South
	West
	East
)

func DirectionName(dir Direction) string {
	switch dir {
	case North:
		return "North"
	case South:
		return "South"
	case West:
		return "West"
	case East:
		return "East"
	default:
		return "Unknown"
	}
}

func MovementDirectionOffset(dir Direction) (int, int) {
	switch dir {
	case North:
		return 0, -1
	case South:
		return 0, 1
	case West:
		return -1, 0
	case East:
		return 1, 0
	}

	return 0, 0
}

type Entity_NamedComponent struct {
	Name string
}

type Entity_DescribedComponent struct {
	Description string
}

type Entity_PresentableComponent struct {
	Rune  rune
	Style tcell.Style
}

type Entity_PositionedComponent struct {
	Position engine.Position
}

type Entity_EquippedComponent struct {
	Inventory *EquippedInventory
}

type Entity_StatsHolderComponent struct {
	BaseStats map[Stat]int
	// StatModifiers []StatModifier
}

type Entity_HealthComponent struct {
	Health    int
	MaxHealth int
	IsDead    bool
}

type Entity interface {
	UniqueId() uuid.UUID

	Named() *Entity_NamedComponent
	Described() *Entity_DescribedComponent
	Presentable() *Entity_PresentableComponent
	Positioned() *Entity_PositionedComponent
	Equipped() *Entity_EquippedComponent
	Stats() *Entity_StatsHolderComponent
	HealthData() *Entity_HealthComponent
}

type BaseEntity struct {
	id uuid.UUID

	named       *Entity_NamedComponent
	described   *Entity_DescribedComponent
	presentable *Entity_PresentableComponent
	positioned  *Entity_PositionedComponent
	equipped    *Entity_EquippedComponent
	stats       *Entity_StatsHolderComponent
	damageable  *Entity_HealthComponent
}

func (be *BaseEntity) UniqueId() uuid.UUID {
	return be.id
}

func (be *BaseEntity) Named() *Entity_NamedComponent {
	return be.named
}

func (be *BaseEntity) Described() *Entity_DescribedComponent {
	return be.described
}

func (be *BaseEntity) Presentable() *Entity_PresentableComponent {
	return be.presentable
}

func (be *BaseEntity) Positioned() *Entity_PositionedComponent {
	return be.positioned
}

func (be *BaseEntity) Equipped() *Entity_EquippedComponent {
	return be.equipped
}

func (be *BaseEntity) Stats() *Entity_StatsHolderComponent {
	return be.stats
}

func (be *BaseEntity) HealthData() *Entity_HealthComponent {
	return be.damageable
}

func CreateEntity(components ...func(*BaseEntity)) *BaseEntity {
	e := &BaseEntity{
		id: uuid.New(),
	}

	for _, comp := range components {
		comp(e)
	}

	return e
}

func WithName(name string) func(*BaseEntity) {
	return func(e *BaseEntity) {
		e.named = &Entity_NamedComponent{
			Name: name,
		}
	}
}

func WithDescription(description string) func(e *BaseEntity) {
	return func(e *BaseEntity) {
		e.described = &Entity_DescribedComponent{
			Description: description,
		}
	}
}

func WithPresentation(symbol rune, style tcell.Style) func(e *BaseEntity) {
	return func(e *BaseEntity) {
		e.presentable = &Entity_PresentableComponent{
			Rune:  symbol,
			Style: style,
		}
	}
}

func WithPosition(pos engine.Position) func(e *BaseEntity) {
	return func(e *BaseEntity) {
		e.positioned = &Entity_PositionedComponent{
			Position: pos,
		}
	}
}

func WithInventory(inv *EquippedInventory) func(e *BaseEntity) {
	return func(e *BaseEntity) {
		e.equipped = &Entity_EquippedComponent{
			Inventory: inv,
		}
	}
}

func WithStats(baseStats map[Stat]int, statModifiers ...StatModifier) func(e *BaseEntity) {
	return func(e *BaseEntity) {
		e.stats = &Entity_StatsHolderComponent{
			BaseStats: baseStats,
			// StatModifiers: statModifiers,
		}
	}
}

func WithHealthData(health, maxHealth int, isDead bool) func(e *BaseEntity) {
	return func(e *BaseEntity) {
		e.damageable = &Entity_HealthComponent{
			Health:    health,
			MaxHealth: maxHealth,
			IsDead:    isDead,
		}
	}
}
