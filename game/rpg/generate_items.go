package rpg

import (
	"math/rand"
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/item"

	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
)

type ItemSupplier func() item.Item

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

func (igt *LootTable) Generate() item.Item {
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

func generateItemName(itemType RPGItemType, rarity ItemRarity) (string, tcell.Style) {
	switch rarity {
	case ItemRarity_Common:
		return itemType.Name(), tcell.StyleDefault
	case ItemRarity_Uncommon:
		return itemType.Name(), tcell.StyleDefault.Foreground(tcell.ColorLime)
	case ItemRarity_Rare:
		return itemType.Name(), tcell.StyleDefault.Foreground(tcell.ColorBlue)
	case ItemRarity_Epic:
		return itemType.Name(), tcell.StyleDefault.Foreground(tcell.ColorPurple)
	case ItemRarity_Legendary:
		return itemType.Name(), tcell.StyleDefault.Foreground(tcell.ColorOrange).Attributes(tcell.AttrBold)
	default:
		return itemType.Name(), tcell.StyleDefault
	}
}

func randomStat() Stat {
	stats := []Stat{
		Stat_Attributes_Strength,
		Stat_Attributes_Dexterity,
		Stat_Attributes_Intelligence,
		Stat_Attributes_Constitution,
		Stat_PhysicalPrecisionBonus,
		Stat_EvasionBonus,
		Stat_MagicPrecisionBonus,
		Stat_TotalPrecisionBonus,
		Stat_DamageBonus_Physical_Unarmed,
		Stat_DamageBonus_Physical_Slashing,
		Stat_DamageBonus_Physical_Piercing,
		Stat_DamageBonus_Physical_Bludgeoning,
		Stat_DamageBonus_Magic_Fire,
		Stat_DamageBonus_Magic_Cold,
		Stat_DamageBonus_Magic_Necrotic,
		Stat_DamageBonus_Magic_Thunder,
		Stat_DamageBonus_Magic_Acid,
		Stat_DamageBonus_Magic_Poison,
		Stat_MaxHealthBonus,
	}

	return stats[rand.Intn(len(stats))]
}

func generateItemStatModifiers(rarity ItemRarity) []StatModifier {
	points := pointPerRarity(rarity)
	modifiers := []StatModifier{}

	for {
		if points <= 0 {
			break
		}

		modAmount := engine.RandInt(-points/2, points)

		if modAmount == 0 {
			continue
		}

		modifiers = append(modifiers, StatModifier{
			Id:    StatModifierId(uuid.New().String()),
			Stat:  randomStat(),
			Bonus: modAmount,
		})

		points -= modAmount
	}

	return modifiers
}

// Each rarity gets an amount of generation points, the higher the rarity, the more points
// Each stat modifier consumes points. The higher the stat bonus, the more points it consumes.
func GenerateItemOfTypeAndRarity(itemType RPGItemType, rarity ItemRarity) RPGItem {
	// points := pointPerRarity(rarity)
	name, style := generateItemName(itemType, rarity)

	return CreateRPGItem(
		name,
		style,
		itemType,
		generateItemStatModifiers(rarity),
	)
}
