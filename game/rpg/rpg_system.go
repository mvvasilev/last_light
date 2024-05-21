package rpg

import (
	"math/rand"
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

	Stat_MaxHealthBonus Stat = 140
)

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
		return Stat_DamageBonus_Magic_Fire
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

func LuckRoll() int {
	return RollD10(1)
}

func TotalModifierForStat(entity RPGEntity, stat Stat) int {
	agg := 0

	for _, m := range entity.CollectModifiersForStat(stat) {
		agg += m.Bonus
	}

	return agg
}

func StatValue(entity RPGEntity, stat Stat) int {
	return entity.BaseStat(stat) + TotalModifierForStat(entity, stat)
}

// Base Max Health is determined from constitution:
// Constitution + Max Health Bonus + 10
func BaseMaxHealth(entity RPGEntity) int {
	return StatValue(entity, Stat_Attributes_Constitution) + StatValue(entity, Stat_MaxHealthBonus) + 10
}

// Dexterity + Evasion bonus + luck roll
func EvasionRoll(victim RPGEntity) int {
	return StatValue(victim, Stat_Attributes_Dexterity) + StatValue(victim, Stat_EvasionBonus) + LuckRoll()
}

// Strength + Precision bonus ( melee + total ) + luck roll
func PhysicalPrecisionRoll(attacker RPGEntity) int {
	return StatValue(attacker, Stat_Attributes_Strength) + StatValue(attacker, Stat_PhysicalPrecisionBonus) + StatValue(attacker, Stat_TotalPrecisionBonus) + LuckRoll()
}

// Intelligence + Precision bonus ( magic + total ) + luck roll
func MagicPrecisionRoll(attacker RPGEntity) int {
	return StatValue(attacker, Stat_Attributes_Intelligence) + StatValue(attacker, Stat_MagicPrecisionBonus) + StatValue(attacker, Stat_TotalPrecisionBonus) + LuckRoll()
}

// true = hit lands, false = hit does not land
func MagicHitRoll(attacker RPGEntity, victim RPGEntity) bool {
	return hitRoll(EvasionRoll(victim), MagicPrecisionRoll(attacker))
}

// true = hit lands, false = hit does not land
func PhysicalHitRoll(attacker RPGEntity, victim RPGEntity) bool {
	return hitRoll(EvasionRoll(victim), PhysicalPrecisionRoll(attacker))
}

func hitRoll(evasionRoll, precisionRoll int) bool {
	if evasionRoll == 20 && precisionRoll == 20 {
		return true
	}

	if evasionRoll == 20 {
		return false
	}

	if precisionRoll == 20 {
		return true
	}

	return evasionRoll < precisionRoll
}

func UnarmedDamage(attacker RPGEntity) int {
	return RollD4(1) + StatValue(attacker, Stat_DamageBonus_Physical_Unarmed)
}
