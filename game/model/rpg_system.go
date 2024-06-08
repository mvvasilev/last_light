package model

import (
	"math/rand"
	"mvvasilev/last_light/engine"
)

type Stat int

const (
	// Used as a default value in cases where no other stat could be determined.
	// Should never be used except in cases of error handling
	Stat_NonExtant Stat = -1

	Stat_Attributes_Strength     Stat = 0
	Stat_Attributes_Dexterity    Stat = 10
	Stat_Attributes_Intelligence Stat = 20
	Stat_Attributes_Constitution Stat = 30

	Stat_PhysicalPrecisionBonus Stat = 5
	Stat_EvasionBonus           Stat = 15
	Stat_MagicPrecisionBonus    Stat = 25
	Stat_TotalPrecisionBonus    Stat = 35

	Stat_DamageBonus_Physical_Unarmed     Stat = 40
	Stat_DamageBonus_Physical_Slashing    Stat = 50
	Stat_DamageBonus_Physical_Piercing    Stat = 60
	Stat_DamageBonus_Physical_Bludgeoning Stat = 70

	Stat_DamageBonus_Magic_Fire     Stat = 80
	Stat_DamageBonus_Magic_Cold     Stat = 90
	Stat_DamageBonus_Magic_Necrotic Stat = 100
	Stat_DamageBonus_Magic_Thunder  Stat = 110
	Stat_DamageBonus_Magic_Acid     Stat = 120
	Stat_DamageBonus_Magic_Poison   Stat = 130

	Stat_ResistanceBonus_Physical_Unarmed     Stat = 200
	Stat_ResistanceBonus_Physical_Slashing    Stat = 210
	Stat_ResistanceBonus_Physical_Piercing    Stat = 220
	Stat_ResistanceBonus_Physical_Bludgeoning Stat = 230

	Stat_ResistanceBonus_Magic_Fire     Stat = 240
	Stat_ResistanceBonus_Magic_Cold     Stat = 250
	Stat_ResistanceBonus_Magic_Necrotic Stat = 260
	Stat_ResistanceBonus_Magic_Thunder  Stat = 270
	Stat_ResistanceBonus_Magic_Acid     Stat = 280
	Stat_ResistanceBonus_Magic_Poison   Stat = 290

	Stat_MaxHealthBonus Stat = 1000
	Stat_SpeedBonus     Stat = 1010
)

func StatLongName(stat Stat) string {
	switch stat {
	case Stat_Attributes_Strength:
		return "Strength"
	case Stat_Attributes_Intelligence:
		return "Intelligence"
	case Stat_Attributes_Dexterity:
		return "Dexterity"
	case Stat_Attributes_Constitution:
		return "Constitution"
	default:
		return "Unknown"
	}
}

func RandomStats(pointsAvailable int, min, max int, stats []Stat) (result map[Stat]int) {
	result = map[Stat]int{}

	for {
		if pointsAvailable == 0 {
			break
		}

		for _, s := range stats {
			if pointsAvailable == 0 {
				break
			}

			limit := pointsAvailable

			if limit > max {
				limit = max
			}

			value := engine.RandInt(min+1, limit+min)

			result[s] += value

			pointsAvailable -= value - min
		}
	}

	return
}

type StatModifierId string

type StatModifier struct {
	Id    StatModifierId
	Stat  Stat
	Bonus int
}

// RPG system is based off of dice rolls

func rollDice(times, sides int) int {
	acc := 0

	for range times {
		acc += 1 + rand.Intn(sides+1)
	}

	return acc
}

func RollD100(times int) int {
	return rollDice(times, 100)
}

func RollD20(times int) int {
	return rollDice(times, 20)
}

func RollD12(times int) int {
	return rollDice(times, 12)
}

func RollD10(times int) int {
	return rollDice(times, 10)
}

func RollD8(times int) int {
	return rollDice(times, 8)
}

func RollD6(times int) int {
	return rollDice(times, 6)
}

func RollD4(times int) int {
	return rollDice(times, 4)
}

// Contests are "meets it, beats it"
//
// Luck roll = 1d10
//
// 2 rolls per attack:
//
// BASIC ATTACKS ( spells and abilities can have special rules, or in leu of special rules, these are used ):
//
// 1. Attack roll ( determines if the attack lands ). Contest between Evasion and Precision.
// 		Evasion = Dexterity + Luck roll.
// 		Precision = ( Strength | Intelligence ) + Luck roll ( intelligence for magic, strength for melee ).
//
// 2. Damage roll ( only if the previous was successful ). Each spell, ability and weapon has its own damage calculation.

type DamageType int

const (
	DamageType_Physical_Unarmed     DamageType = 0
	DamageType_Physical_Slashing    DamageType = 1
	DamageType_Physical_Piercing    DamageType = 2
	DamageType_Physical_Bludgeoning DamageType = 3

	DamageType_Magic_Fire     DamageType = 4
	DamageType_Magic_Cold     DamageType = 5
	DamageType_Magic_Necrotic DamageType = 6
	DamageType_Magic_Thunder  DamageType = 7
	DamageType_Magic_Acid     DamageType = 8
	DamageType_Magic_Poison   DamageType = 9
)

func DamageTypeName(dmgType DamageType) string {
	switch dmgType {
	case DamageType_Physical_Unarmed:
		return "Unarmed"
	case DamageType_Physical_Slashing:
		return "Slashing"
	case DamageType_Physical_Piercing:
		return "Piercing"
	case DamageType_Physical_Bludgeoning:
		return "Bludgeoning"
	case DamageType_Magic_Fire:
		return "Fire"
	case DamageType_Magic_Cold:
		return "Cold"
	case DamageType_Magic_Necrotic:
		return "Necrotic"
	case DamageType_Magic_Thunder:
		return "Thunder"
	case DamageType_Magic_Acid:
		return "Acid"
	case DamageType_Magic_Poison:
		return "Poison"
	default:
		return "Unknown"
	}
}

func DamageTypeToBonusStat(dmgType DamageType) Stat {
	switch dmgType {
	case DamageType_Physical_Unarmed:
		return Stat_DamageBonus_Physical_Unarmed
	case DamageType_Physical_Slashing:
		return Stat_DamageBonus_Physical_Slashing
	case DamageType_Physical_Piercing:
		return Stat_DamageBonus_Physical_Piercing
	case DamageType_Physical_Bludgeoning:
		return Stat_DamageBonus_Physical_Bludgeoning
	case DamageType_Magic_Fire:
		return Stat_DamageBonus_Magic_Fire
	case DamageType_Magic_Cold:
		return Stat_DamageBonus_Magic_Cold
	case DamageType_Magic_Necrotic:
		return Stat_DamageBonus_Magic_Necrotic
	case DamageType_Magic_Thunder:
		return Stat_DamageBonus_Magic_Thunder
	case DamageType_Magic_Acid:
		return Stat_DamageBonus_Magic_Acid
	case DamageType_Magic_Poison:
		return Stat_DamageBonus_Magic_Poison
	default:
		return Stat_NonExtant
	}
}

func DamageTypeToResistanceStat(dmgType DamageType) Stat {
	switch dmgType {
	case DamageType_Physical_Unarmed:
		return Stat_ResistanceBonus_Physical_Unarmed
	case DamageType_Physical_Slashing:
		return Stat_ResistanceBonus_Physical_Slashing
	case DamageType_Physical_Piercing:
		return Stat_ResistanceBonus_Physical_Piercing
	case DamageType_Physical_Bludgeoning:
		return Stat_ResistanceBonus_Physical_Bludgeoning
	case DamageType_Magic_Fire:
		return Stat_ResistanceBonus_Magic_Fire
	case DamageType_Magic_Cold:
		return Stat_ResistanceBonus_Magic_Cold
	case DamageType_Magic_Necrotic:
		return Stat_ResistanceBonus_Magic_Necrotic
	case DamageType_Magic_Thunder:
		return Stat_ResistanceBonus_Magic_Thunder
	case DamageType_Magic_Acid:
		return Stat_ResistanceBonus_Magic_Acid
	case DamageType_Magic_Poison:
		return Stat_ResistanceBonus_Magic_Poison
	default:
		return Stat_NonExtant
	}
}

func LuckRoll() int {
	return RollD10(1)
}

func TotalModifierForStat(stats *Item_StatModifierComponent, stat Stat) int {
	agg := 0

	for _, m := range stats.StatModifiers {
		if m.Stat == stat {
			agg += m.Bonus
		}
	}

	return agg
}

func statValue(stats *Entity_StatsHolderComponent, stat Stat) int {
	return stats.BaseStats[stat]
}

func statModifierValue(statModifiers []StatModifier, stat Stat) int {
	for _, sm := range statModifiers {
		if sm.Stat == stat {
			return sm.Bonus
		}
	}

	return 0
}

// Base Max Health is determined from constitution:
// 5*Constitution + Max Health Bonus
func BaseMaxHealth(entity Entity) int {
	stats := entity.Stats()

	if stats == nil {
		return 0
	}

	return 5*statValue(stats, Stat_Attributes_Constitution) + statValue(stats, Stat_MaxHealthBonus)
}

// Dexterity + Evasion bonus + luck roll
func EvasionRoll(victim Entity) int {
	if victim.Stats() == nil {
		return 0
	}

	return statValue(victim.Stats(), Stat_Attributes_Dexterity) + statValue(victim.Stats(), Stat_EvasionBonus) + LuckRoll()
}

// Strength + Precision bonus ( melee + total ) + luck roll
func PhysicalPrecisionRoll(attacker Entity) int {
	if attacker.Stats() == nil {
		return 0
	}

	return statValue(attacker.Stats(), Stat_Attributes_Strength) + statValue(attacker.Stats(), Stat_PhysicalPrecisionBonus) + statValue(attacker.Stats(), Stat_TotalPrecisionBonus) + LuckRoll()
}

// Intelligence + Precision bonus ( magic + total ) + luck roll
func MagicPrecisionRoll(attacker Entity) int {
	if attacker.Stats() == nil {
		return 0
	}

	return statValue(attacker.Stats(), Stat_Attributes_Intelligence) + statValue(attacker.Stats(), Stat_MagicPrecisionBonus) + statValue(attacker.Stats(), Stat_TotalPrecisionBonus) + LuckRoll()
}

// true = hit lands, false = hit does not land
func MagicHitRoll(attacker Entity, victim Entity) bool {
	return hitRoll(EvasionRoll(victim), MagicPrecisionRoll(attacker))
}

// true = hit lands, false = hit does not land
func PhysicalHitRoll(attacker Entity, victim Entity) (hit bool, evasion, precision int) {
	evasion = EvasionRoll(victim)
	precision = PhysicalPrecisionRoll(attacker)
	hit = hitRoll(evasion, precision)

	return
}

func hitRoll(evasionRoll, precisionRoll int) bool {
	return evasionRoll < precisionRoll
}

func UnarmedDamage(attacker Entity) int {
	if attacker.Stats() == nil {
		return 0
	}

	return RollD4(1) + statValue(attacker.Stats(), Stat_DamageBonus_Physical_Unarmed)
}

func PhysicalWeaponDamage(attacker Entity, weapon Item, victim Entity) (totalDamage int, dmgType DamageType) {
	if attacker.Stats() == nil || weapon.Damaging() == nil || victim.Stats() == nil {
		return UnarmedDamage(attacker), DamageType_Physical_Unarmed
	}

	totalDamage, dmgType = weapon.Damaging().DamageRoll()

	bonusDmgStat := DamageTypeToBonusStat(dmgType)
	dmgResistStat := DamageTypeToResistanceStat(dmgType)

	totalDamage = totalDamage + statValue(attacker.Stats(), bonusDmgStat) - statValue(victim.Stats(), dmgResistStat)

	if weapon.StatModifier() != nil {
		totalDamage += statModifierValue(weapon.StatModifier().StatModifiers, bonusDmgStat)
	}

	if totalDamage <= 0 {
		return 0, dmgType
	}

	return
}

func UnarmedAttack(attacker Entity, victim Entity) (hit bool, precisionRoll, evasionRoll int, damage int, damageType DamageType) {
	hit, evasionRoll, precisionRoll = PhysicalHitRoll(attacker, victim)

	if !hit {
		return
	}

	damage = UnarmedDamage(attacker)
	damageType = DamageType_Physical_Unarmed

	return
}

func PhysicalWeaponAttack(attacker Entity, weapon Item, victim Entity) (hit bool, precisionRoll, evasionRoll int, damage int, damageType DamageType) {
	hit, evasionRoll, precisionRoll = PhysicalHitRoll(attacker, victim)

	if !hit {
		return
	}

	damage, damageType = PhysicalWeaponDamage(attacker, weapon, victim)

	return
}
