package model

// type RPGItemType interface {
// 	RollDamage(victim, attacker RPGEntity) (damage int, dmgType DamageType)
// 	Use(eventLogger *engine.GameEventLog, user RPGEntity)
// 	MetaTypes() []ItemMetaType

// 	item.ItemType
// }

// type BasicRPGItemType struct {
// 	damageRollFunc func(victim, attacker RPGEntity) (damage int, dmgType DamageType)
// 	useFunc        func(eventLogger *engine.GameEventLog, user RPGEntity)

// 	metaTypes []ItemMetaType

// 	*item.BasicItemType
// }

// func (it *BasicRPGItemType) Use(eventLogger *engine.GameEventLog, user RPGEntity) {
// 	if it.useFunc == nil {
// 		return
// 	}

// 	it.useFunc(eventLogger, user)
// }

// func (it *BasicRPGItemType) RollDamage(victim, attacker RPGEntity) (damage int, dmgType DamageType) {
// 	if it.damageRollFunc == nil {
// 		return 0, DamageType_Physical_Unarmed
// 	}

// 	return it.damageRollFunc(victim, attacker)
// }

// func (it *BasicRPGItemType) MetaTypes() []ItemMetaType {
// 	return it.metaTypes
// }

// type RPGItem interface {
// 	Modifiers() []StatModifier
// 	RPGType() RPGItemType

// 	item.Item
// }

// type BasicRPGItem struct {
// 	modifiers []StatModifier
// 	rpgType   RPGItemType

// 	item.BasicItem
// }

// func (i *BasicRPGItem) Modifiers() []StatModifier {
// 	return i.modifiers
// }

// func (i *BasicRPGItem) RPGType() RPGItemType {
// 	return i.rpgType
// }

// func CreateRPGItem(name string, style tcell.Style, itemType RPGItemType, modifiers []StatModifier) RPGItem {
// 	return &BasicRPGItem{
// 		modifiers: modifiers,
// 		rpgType:   itemType,
// 		BasicItem: item.CreateBasicItemWithName(
// 			name,
// 			style,
// 			itemType,
// 			1,
// 		),
// 	}
// }
