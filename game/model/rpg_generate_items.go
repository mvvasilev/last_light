package model

import (
	"math/rand"
	"mvvasilev/last_light/engine"
	"slices"
	"unicode"
	"unicode/utf8"

	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
)

const MaxNumberOfModifiers = 6

type ItemSupplier func() Item_V2

type LootTable struct {
	table []ItemSupplier
}

func CreateLootTable() *LootTable {
	return &LootTable{
		table: make([]ItemSupplier, 0),
	}
}

func (igt *LootTable) Add(weight int, createItemFunction ItemSupplier) {
	for range weight {
		igt.table = append(igt.table, createItemFunction)
	}
}

func (igt *LootTable) Generate() Item_V2 {
	return igt.table[rand.Intn(len(igt.table))]()
}

type ItemRarity int

const (
	ItemRarity_Common    ItemRarity = 0
	ItemRarity_Uncommon  ItemRarity = 1
	ItemRarity_Rare      ItemRarity = 2
	ItemRarity_Epic      ItemRarity = 3
	ItemRarity_Legendary ItemRarity = 4
)

func pointPerRarity(rarity ItemRarity) int {
	switch rarity {
	case ItemRarity_Common:
		return 0
	case ItemRarity_Uncommon:
		return 3
	case ItemRarity_Rare:
		return 5
	case ItemRarity_Epic:
		return 8
	case ItemRarity_Legendary:
		return 13
	default:
		return 0
	}
}

func generateUniqueItemName() string {
	starts := []string{
		"du", "nol", "ma", "re",
		"ka", "gro", "hru", "lo",
		"ara", "ke", "ko", "uro",
		"ne", "pe", "pa", "pho",
	}

	middles := []string{
		"kora", "duru", "kolku", "dila",
		"luio", "ghro", "kelma", "riga",
		"fela", "fiya", "numa", "ruta",
	}

	end := []string{
		"dum", "dor", "dar", "thar",
		"thor", "thum", "hor", "hum",
		"her", "kom", "kur", "kyr",
		"mor", "mar", "man", "kum",
		"tum",
	}

	name := starts[rand.Intn(len(starts))] + middles[rand.Intn(len(middles))] + end[rand.Intn(len(end))]

	r, size := utf8.DecodeRuneInString(name)

	return string(unicode.ToUpper(r)) + name[size:]
}

func randomAdjective() string {
	adjectives := []string{
		"shiny", "gruesome", "sharp", "tattered",
		"mediocre", "unusual", "bright", "rusty",
		"dreadful", "exceptional", "old", "bent",
		"ancient", "crude", "dented", "cool",
	}

	adj := adjectives[rand.Intn(len(adjectives))]

	r, size := utf8.DecodeRuneInString(adj)

	return string(unicode.ToUpper(r)) + adj[size:]
}

func randomSuffix() string {
	suffixes := []string{
		"of the Monkey", "of the Tiger", "of the Elephant", "of the Slug",
		"of Elven Make",
	}

	return suffixes[rand.Intn(len(suffixes))]
}

func generateItemName(existingItemName string, rarity ItemRarity) (string, tcell.Style) {
	switch rarity {
	case ItemRarity_Common:
		return existingItemName, tcell.StyleDefault
	case ItemRarity_Uncommon:
		return randomAdjective() + " " + existingItemName, tcell.StyleDefault.Foreground(tcell.ColorLime)
	case ItemRarity_Rare:
		return existingItemName + " " + randomSuffix(), tcell.StyleDefault.Foreground(tcell.ColorBlue)
	case ItemRarity_Epic:
		return randomAdjective() + " " + existingItemName + " " + randomSuffix(), tcell.StyleDefault.Foreground(tcell.ColorPurple)
	case ItemRarity_Legendary:
		return generateUniqueItemName() + ", Legendary " + existingItemName, tcell.StyleDefault.Foreground(tcell.ColorOrange).Attributes(tcell.AttrBold)
	default:
		return existingItemName, tcell.StyleDefault
	}
}

func randomStat(metaItemTypes []ItemMetaType) Stat {
	stats := make(map[ItemMetaType][]Stat, 0)

	stats[MetaItemType_Weapon] = []Stat{
		Stat_Attributes_Strength,
		Stat_Attributes_Dexterity,
		Stat_Attributes_Intelligence,
		Stat_Attributes_Constitution,
		Stat_TotalPrecisionBonus,
	}

	stats[MetaItemType_Physical_Weapon] = []Stat{
		Stat_PhysicalPrecisionBonus,
		Stat_DamageBonus_Physical_Slashing,
		Stat_DamageBonus_Physical_Piercing,
		Stat_DamageBonus_Physical_Bludgeoning,
		Stat_DamageBonus_Magic_Fire,
		Stat_DamageBonus_Magic_Cold,
		Stat_DamageBonus_Magic_Necrotic,
		Stat_DamageBonus_Magic_Thunder,
		Stat_DamageBonus_Magic_Acid,
		Stat_DamageBonus_Magic_Poison,
	}

	stats[MetaItemType_Magic_Weapon] = []Stat{
		Stat_MagicPrecisionBonus,
		Stat_DamageBonus_Magic_Fire,
		Stat_DamageBonus_Magic_Cold,
		Stat_DamageBonus_Magic_Necrotic,
		Stat_DamageBonus_Magic_Thunder,
		Stat_DamageBonus_Magic_Acid,
		Stat_DamageBonus_Magic_Poison,
	}

	stats[MetaItemType_Armour] = []Stat{
		Stat_EvasionBonus,
		Stat_DamageBonus_Physical_Unarmed,
		Stat_MaxHealthBonus,
	}

	stats[MetaItemType_Magic_Armour] = []Stat{
		Stat_MagicPrecisionBonus,
		Stat_DamageBonus_Magic_Fire,
		Stat_DamageBonus_Magic_Cold,
		Stat_DamageBonus_Magic_Necrotic,
		Stat_DamageBonus_Magic_Thunder,
		Stat_DamageBonus_Magic_Acid,
		Stat_DamageBonus_Magic_Poison,
		Stat_ResistanceBonus_Magic_Acid,
		Stat_ResistanceBonus_Magic_Cold,
		Stat_ResistanceBonus_Magic_Fire,
		Stat_ResistanceBonus_Magic_Necrotic,
		Stat_ResistanceBonus_Magic_Poison,
		Stat_ResistanceBonus_Magic_Thunder,
	}

	stats[MetaItemType_Physical_Armour] = []Stat{
		Stat_PhysicalPrecisionBonus,
		Stat_DamageBonus_Physical_Slashing,
		Stat_DamageBonus_Physical_Piercing,
		Stat_DamageBonus_Physical_Bludgeoning,
		Stat_ResistanceBonus_Physical_Bludgeoning,
		Stat_ResistanceBonus_Physical_Piercing,
		Stat_ResistanceBonus_Physical_Slashing,
		Stat_ResistanceBonus_Physical_Unarmed,
	}

	possibleStats := make([]Stat, 0, 10)

	for _, mt := range metaItemTypes {
		possibleStats = append(possibleStats, stats[mt]...)
	}

	return slices.Compact(possibleStats)[rand.Intn(len(stats))]
}

func generateItemStatModifiers(itemMetaTypes []ItemMetaType, rarity ItemRarity) []StatModifier {
	points := pointPerRarity(rarity)
	modifiers := make(map[Stat]*StatModifier, 0)

	for {
		// If no points remain, or if the number of modifiers on the item reaches the maximum
		if points <= 0 || len(modifiers) == MaxNumberOfModifiers {
			break
		}

		// Random chance to increase or decrease a stat
		modAmount := engine.RandInt(-points, points)

		if modAmount == 0 {
			continue
		}

		stat := randomStat(itemMetaTypes)

		existingForStat := modifiers[stat]

		// If this stat modifier already exists on the item, add the new modification amount to the old
		if existingForStat != nil {
			existingForStat.Bonus += modAmount

			// If the added amount is 0, remove the modifier
			if existingForStat.Bonus == 0 {
				delete(modifiers, stat)
			} else {
				modifiers[stat] = existingForStat
			}

		} else {
			// Otherwise, append a new stat modifier
			modifiers[stat] = &StatModifier{
				Id:    StatModifierId(uuid.New().String()),
				Stat:  stat,
				Bonus: modAmount,
			}
		}

		// Decrease amount of points left by absolute value
		points -= engine.AbsInt(modAmount)
	}

	vals := make([]StatModifier, 0, len(modifiers))

	for _, v := range modifiers {
		vals = append(vals, *v)
	}

	return vals
}

// Each rarity gets an amount of generation points, the higher the rarity, the more points
// Each stat modifier consumes points. The higher the stat bonus, the more points it consumes.
func GenerateItemOfTypeAndRarity(prototype Item_V2, rarity ItemRarity) Item_V2 {
	if prototype.Named() == nil {
		return prototype
	}

	if prototype.MetaTypes() == nil {
		return prototype
	}

	existingName := prototype.Named().Name
	metaTypes := prototype.MetaTypes().Types

	// points := pointPerRarity(rarity)
	name, style := generateItemName(existingName, rarity)
	statModifiers := generateItemStatModifiers(metaTypes, rarity)

	return createBaseItem(
		prototype.Type(),
		prototype.TileIcon(),
		prototype.Icon(),
		prototype.Style(),
		item_WithName(name, style),
		item_WithDescription(prototype.Described().Description, prototype.Described().Style),
		item_WithDamaging(prototype.Damaging().DamageRoll),
		item_WithEquippable(prototype.Equippable().Slot),
		item_WithStatModifiers(statModifiers),
		item_WithMetaTypes(metaTypes),
	)
}
