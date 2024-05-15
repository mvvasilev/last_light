package ecs

import (
	"slices"
	"time"

	"github.com/gdamore/tcell/v2"
)

/* ===================== ECSError ===================== */

type ECSError struct {
	err string
}

func ecsError(err string) *ECSError {
	return &ECSError{
		err: err,
	}
}

func (e *ECSError) Error() string {
	return e.err
}

/* ==================================================== */

/* ================== ComponentType =================== */

type ComponentType uint64

func TypeFrom(powerOf2 uint64) (ComponentType, *ECSError) {
	if powerOf2 > 63 {
		return 0, ecsError("Failure: Provided number is too high ( component types must be represented by a power of 2, up to 63 )")
	}

	t := uint64(0)

	for i := range powerOf2 {
		t *= i
	}

	return ComponentType(t), nil
}

/* ==================================================== */

/* ====================== Types ======================= */

type SystemOrder uint8

type EntityId uint64

type ComponentMask uint64

/* ==================================================== */

/* ================== ComponentMask =================== */

func MaskOf(ids ...ComponentType) ComponentMask {
	mask := uint64(0)

	for _, id := range ids {
		mask |= uint64(id)
	}

	return ComponentMask(mask)
}

func (c ComponentMask) CombinedWithMask(mask ComponentMask) ComponentMask {
	return ComponentMask(uint64(c) | uint64(mask))
}

func (c ComponentMask) CombinedWithType(id ComponentType) ComponentMask {
	return ComponentMask(uint64(c) | uint64(id))
}

func (c ComponentMask) Contains(id ComponentType) bool {
	return (uint64(id) & uint64(c)) == uint64(id)
}

func (c ComponentMask) ContainsMultiple(ids ComponentMask) bool {
	return (uint64(ids) & uint64(c)) == uint64(ids)
}

/* ==================================================== */

/* ===================== System ======================= */

type System interface {
	Name() string
	Order() int
	Input(world *World, e tcell.EventKey)
	Tick(world *World, deltaTime int64)
}

/* ==================================================== */

/* =================== Component ====================== */

type Component interface {
	Type() ComponentType
}

/* ==================================================== */

/* ================== BasicEntity ===================== */

type BasicEntity struct {
	id                  EntityId
	containedComponents ComponentMask
	components          map[ComponentType]Component
}

func CreateRandomEntityId() EntityId {
	return EntityId(time.Now().UnixNano())
}

func createEntity(components ...Component) *BasicEntity {
	ent := &BasicEntity{
		id:         CreateRandomEntityId(),
		components: make(map[ComponentType]Component, 0),
	}

	for _, c := range components {
		ent.components[c.Type()] = c
	}

	return ent
}

func (ent *BasicEntity) Id() EntityId {
	return ent.id
}

func (ent *BasicEntity) ContainedComponents() ComponentMask {
	return ComponentMask(ent.containedComponents)
}

func (ent *BasicEntity) addComponent(c Component) {
	ent.containedComponents = ent.containedComponents.CombinedWithType(c.Type())
	ent.components[c.Type()] = c
}

func (ent *BasicEntity) AllComponents() []Component {
	vals := make([]Component, len(ent.components))

	for _, v := range ent.components {
		vals = append(vals, v)
	}

	return vals
}

func (ent *BasicEntity) QueryComponents(componentIds ...ComponentType) (components []Component, err *ECSError) {
	comps := make([]Component, len(componentIds))

	for _, id := range componentIds {
		comp := ent.components[id]

		if comp == nil {
			return nil, ecsError("Failure: Entity does not contain all of requested types")
		}

		comps = append(comps, comp)
	}

	return comps, nil
}

func (ent *BasicEntity) ContainsComponents(mask ComponentMask) bool {
	return ent.containedComponents.ContainsMultiple(mask)
}

func (ent *BasicEntity) FetchComponent(id ComponentType) (component Component, err *ECSError) {
	comp := ent.components[id]

	if comp == nil {
		return nil, ecsError("Failure: Entity does not contain requested component")
	}

	return comp, nil
}

/* ==================================================== */

/* ==================== World ========================= */

type World struct {
	registeredComponentTypes ComponentMask
	registeredComponentNames map[ComponentType]string

	entities   map[EntityId]*BasicEntity
	components map[ComponentType][]Component
	systems    []System
}

func CreateWorld() *World {
	return &World{
		entities:   make(map[EntityId]*BasicEntity, 0),
		systems:    make([]System, 0),
		components: make(map[ComponentType][]Component, 0),
	}
}

func (w *World) QueryComponents(componentIds ...ComponentType) (components map[ComponentType][]Component, err *ECSError) {
	if !w.registeredComponentTypes.ContainsMultiple(MaskOf(componentIds...)) {
		return nil, ecsError("Failure: One of the provided queries component types has not been registered")
	}

	comps := make(map[ComponentType][]Component, 0)

	for _, id := range componentIds {
		comp := w.components[id]

		// No components of the requested type exist, that's ok, add an empty slice for that type
		if comp == nil {
			comps[id] = make([]Component, 0)
			continue
		}

		comps[id] = comp
	}

	return comps, nil
}

func (w *World) RegisterComponentType(t ComponentType, name string) (err *ECSError) {
	if w.registeredComponentTypes.Contains(t) {
		return ecsError("Failure: ComponentType conflict, another component already exists with that type number")
	}

	w.registeredComponentTypes = w.registeredComponentTypes.CombinedWithType(t)
	w.registeredComponentNames[t] = name

	return nil
}

func (w *World) FindEntitiesWithComponents(comps ComponentMask) []*BasicEntity {
	ents := make([]*BasicEntity, 16)

	for _, v := range w.entities {
		if v.ContainsComponents(comps) {
			ents = append(ents, v)
		}
	}

	return ents
}

func (w *World) AddComponentToEntity(ent *BasicEntity, comp Component) (modifiedEntity *BasicEntity, err *ECSError) {
	if !w.registeredComponentTypes.Contains(comp.Type()) {
		return nil, ecsError("Failure: Attempting to add unknown component to an entity.")
	}

	if ent.ContainsComponents(ComponentMask(comp.Type())) {
		return nil, ecsError("Failure: Entity already contains component")
	}

	ent.addComponent(comp)

	if w.components[comp.Type()] == nil {
		w.components[comp.Type()] = make([]Component, 0)
	}

	w.components[comp.Type()] = append(w.components[comp.Type()], comp)

	return ent, nil
}

func (w *World) AddSystem(s System) (err *ECSError) {
	w.systems = append(w.systems, s)

	slices.SortFunc(w.systems, func(a System, b System) int { return a.Order() - b.Order() })

	return nil
}

func (w *World) CreateEntity(comps ...Component) *BasicEntity {
	ent := createEntity(comps...)

	w.entities[ent.Id()] = ent

	return ent
}

func (w *World) FindEntityById(id EntityId) (entity *BasicEntity, err *ECSError) {
	ent := w.entities[id]

	if ent == nil {
		return nil, ecsError("Failure: No entity with request id exists")
	}

	return ent, nil
}

func (w *World) RemoveEntity(id EntityId) {
	delete(w.entities, id)
}

func (w *World) Tick(dt int64) {
	for _, s := range w.systems {
		s.Tick(w, dt)
	}
}

func (w *World) Input(e tcell.EventKey) {
	for _, s := range w.systems {
		s.Input(w, e)
	}
}

/* ==================================================== */
