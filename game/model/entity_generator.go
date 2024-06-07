package model

import "math/rand"

type EntitySupplier func(x, y int) Entity

type EntityTable struct {
	table []EntitySupplier
}

func CreateEntityTable() *EntityTable {
	return &EntityTable{
		table: make([]EntitySupplier, 0),
	}
}

func (igt *EntityTable) Add(weight int, createItemFunction EntitySupplier) {
	for range weight {
		igt.table = append(igt.table, createItemFunction)
	}
}

func (igt *EntityTable) Generate(x, y int) Entity {
	return igt.table[rand.Intn(len(igt.table))](x, y)
}
