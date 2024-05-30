package rpg

type RPGEntity interface {
	BaseStat(stat Stat) int
	SetBaseStat(stat Stat, value int)

	CollectModifiersForStat(stat Stat) []StatModifier
	AddStatModifier(modifier StatModifier)
	RemoveStatModifier(id StatModifierId)

	CurrentHealth() int
	Heal(health int)
	Damage(damage int)
}

type BasicRPGEntity struct {
	stats map[Stat]int

	statModifiers map[Stat][]StatModifier

	currentHealth int
}

func CreateBasicRPGEntity(baseStats map[Stat]int, statModifiers map[Stat][]StatModifier) *BasicRPGEntity {
	return &BasicRPGEntity{
		stats:         baseStats,
		statModifiers: statModifiers,
		currentHealth: 0,
	}
}

func (brpg *BasicRPGEntity) BaseStat(stat Stat) int {
	return brpg.stats[stat]
}

func (brpg *BasicRPGEntity) SetBaseStat(stat Stat, value int) {
	brpg.stats[stat] = value
}

func (brpg *BasicRPGEntity) CollectModifiersForStat(stat Stat) []StatModifier {
	modifiers := brpg.statModifiers[stat]

	if modifiers == nil {
		return []StatModifier{}
	}

	return modifiers
}

func (brpg *BasicRPGEntity) AddStatModifier(modifier StatModifier) {
	existing := brpg.statModifiers[modifier.Stat]

	if existing == nil {
		existing = make([]StatModifier, 0)
	}

	existing = append(existing, modifier)

	brpg.statModifiers[modifier.Stat] = existing
}

func (brpg *BasicRPGEntity) RemoveStatModifier(id StatModifierId) {
	// TODO
}

func (brpg *BasicRPGEntity) CurrentHealth() int {
	return brpg.currentHealth
}

func (brpg *BasicRPGEntity) Heal(health int) {
	maxHealth := BaseMaxHealth(brpg)

	if brpg.currentHealth+health > maxHealth {
		brpg.currentHealth = maxHealth

		return
	}

	brpg.currentHealth += health
}

func (brpg *BasicRPGEntity) Damage(damage int) {
	if brpg.currentHealth-damage < 0 {
		brpg.currentHealth = 0

		return
	}

	brpg.currentHealth -= damage
}
