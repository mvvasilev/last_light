package ecs

import (
	"reflect"
	"testing"

	"github.com/gdamore/tcell/v2"
)

func Test_ecsError(t *testing.T) {
	type args struct {
		err string
	}
	tests := []struct {
		name string
		args args
		want *ECSError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ecsError(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ecsError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestECSError_Error(t *testing.T) {
	type fields struct {
		err string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ECSError{
				err: tt.fields.err,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("ECSError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaskOf(t *testing.T) {
	type args struct {
		ids []ComponentType
	}
	tests := []struct {
		name string
		args args
		want ComponentMask
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaskOf(tt.args.ids...); got != tt.want {
				t.Errorf("MaskOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponentMask_CombinedWithMask(t *testing.T) {
	type args struct {
		mask ComponentMask
	}
	tests := []struct {
		name string
		c    ComponentMask
		args args
		want ComponentMask
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.CombinedWithMask(tt.args.mask); got != tt.want {
				t.Errorf("ComponentMask.CombinedWithMask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponentMask_CombinedWithType(t *testing.T) {
	type args struct {
		id ComponentType
	}
	tests := []struct {
		name string
		c    ComponentMask
		args args
		want ComponentMask
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.CombinedWithType(tt.args.id); got != tt.want {
				t.Errorf("ComponentMask.CombinedWithType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponentMask_Contains(t *testing.T) {
	type args struct {
		id ComponentType
	}
	tests := []struct {
		name string
		c    ComponentMask
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Contains(tt.args.id); got != tt.want {
				t.Errorf("ComponentMask.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponentMask_ContainsMultiple(t *testing.T) {
	type args struct {
		ids ComponentMask
	}
	tests := []struct {
		name string
		c    ComponentMask
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ContainsMultiple(tt.args.ids); got != tt.want {
				t.Errorf("ComponentMask.ContainsMultiple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateRandomEntityId(t *testing.T) {
	tests := []struct {
		name string
		want EntityId
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateRandomEntityId(); got != tt.want {
				t.Errorf("CreateRandomEntityId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createEntity(t *testing.T) {
	type args struct {
		components []Component
	}
	tests := []struct {
		name string
		args args
		want *BasicEntity
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createEntity(tt.args.components...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBasicEntity_Id(t *testing.T) {
	type fields struct {
		id                  EntityId
		containedComponents ComponentMask
		components          map[ComponentType]Component
	}
	tests := []struct {
		name   string
		fields fields
		want   EntityId
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ent := &BasicEntity{
				id:                  tt.fields.id,
				containedComponents: tt.fields.containedComponents,
				components:          tt.fields.components,
			}
			if got := ent.Id(); got != tt.want {
				t.Errorf("BasicEntity.Id() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBasicEntity_ContainedComponents(t *testing.T) {
	type fields struct {
		id                  EntityId
		containedComponents ComponentMask
		components          map[ComponentType]Component
	}
	tests := []struct {
		name   string
		fields fields
		want   ComponentMask
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ent := &BasicEntity{
				id:                  tt.fields.id,
				containedComponents: tt.fields.containedComponents,
				components:          tt.fields.components,
			}
			if got := ent.ContainedComponents(); got != tt.want {
				t.Errorf("BasicEntity.ContainedComponents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBasicEntity_AddComponent(t *testing.T) {
	type fields struct {
		id                  EntityId
		containedComponents ComponentMask
		components          map[ComponentType]Component
	}
	type args struct {
		c Component
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ent := &BasicEntity{
				id:                  tt.fields.id,
				containedComponents: tt.fields.containedComponents,
				components:          tt.fields.components,
			}
			ent.addComponent(tt.args.c)
		})
	}
}

func TestBasicEntity_AllComponents(t *testing.T) {
	type fields struct {
		id                  EntityId
		containedComponents ComponentMask
		components          map[ComponentType]Component
	}
	tests := []struct {
		name   string
		fields fields
		want   []Component
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ent := &BasicEntity{
				id:                  tt.fields.id,
				containedComponents: tt.fields.containedComponents,
				components:          tt.fields.components,
			}
			if got := ent.AllComponents(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BasicEntity.AllComponents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBasicEntity_QueryComponents(t *testing.T) {
	type fields struct {
		id                  EntityId
		containedComponents ComponentMask
		components          map[ComponentType]Component
	}
	type args struct {
		componentIds []ComponentType
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantComponents []Component
		wantErr        *ECSError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ent := &BasicEntity{
				id:                  tt.fields.id,
				containedComponents: tt.fields.containedComponents,
				components:          tt.fields.components,
			}
			gotComponents, gotErr := ent.QueryComponents(tt.args.componentIds...)
			if !reflect.DeepEqual(gotComponents, tt.wantComponents) {
				t.Errorf("BasicEntity.QueryComponents() gotComponents = %v, want %v", gotComponents, tt.wantComponents)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("BasicEntity.QueryComponents() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestBasicEntity_ContainsComponents(t *testing.T) {
	type fields struct {
		id                  EntityId
		containedComponents ComponentMask
		components          map[ComponentType]Component
	}
	type args struct {
		mask ComponentMask
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ent := &BasicEntity{
				id:                  tt.fields.id,
				containedComponents: tt.fields.containedComponents,
				components:          tt.fields.components,
			}
			if got := ent.ContainsComponents(tt.args.mask); got != tt.want {
				t.Errorf("BasicEntity.ContainsComponents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBasicEntity_FetchComponent(t *testing.T) {
	type fields struct {
		id                  EntityId
		containedComponents ComponentMask
		components          map[ComponentType]Component
	}
	type args struct {
		id ComponentType
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantComponent Component
		wantErr       *ECSError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ent := &BasicEntity{
				id:                  tt.fields.id,
				containedComponents: tt.fields.containedComponents,
				components:          tt.fields.components,
			}
			gotComponent, gotErr := ent.FetchComponent(tt.args.id)
			if !reflect.DeepEqual(gotComponent, tt.wantComponent) {
				t.Errorf("BasicEntity.FetchComponent() gotComponent = %v, want %v", gotComponent, tt.wantComponent)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("BasicEntity.FetchComponent() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestCreateWorld(t *testing.T) {
	tests := []struct {
		name string
		want *World
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateWorld(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateWorld() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorld_QueryComponents(t *testing.T) {
	type fields struct {
		registeredComponentTypes ComponentMask
		registeredComponentNames map[ComponentType]string
		entities                 map[EntityId]*BasicEntity
		components               map[ComponentType][]Component
		systems                  []System
	}
	type args struct {
		componentIds []ComponentType
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantComponents map[ComponentType][]Component
		wantErr        *ECSError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &World{
				registeredComponentTypes: tt.fields.registeredComponentTypes,
				registeredComponentNames: tt.fields.registeredComponentNames,
				entities:                 tt.fields.entities,
				components:               tt.fields.components,
				systems:                  tt.fields.systems,
			}
			gotComponents, gotErr := w.QueryComponents(tt.args.componentIds...)
			if !reflect.DeepEqual(gotComponents, tt.wantComponents) {
				t.Errorf("World.QueryComponents() gotComponents = %v, want %v", gotComponents, tt.wantComponents)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("World.QueryComponents() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestWorld_RegisterComponentType(t *testing.T) {
	type fields struct {
		registeredComponentTypes ComponentMask
		registeredComponentNames map[ComponentType]string
		entities                 map[EntityId]*BasicEntity
		components               map[ComponentType][]Component
		systems                  []System
	}
	type args struct {
		t    ComponentType
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr *ECSError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &World{
				registeredComponentTypes: tt.fields.registeredComponentTypes,
				registeredComponentNames: tt.fields.registeredComponentNames,
				entities:                 tt.fields.entities,
				components:               tt.fields.components,
				systems:                  tt.fields.systems,
			}
			if gotErr := w.RegisterComponentType(tt.args.t, tt.args.name); !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("World.RegisterComponentType() = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestWorld_AddComponentToEntity(t *testing.T) {
	type fields struct {
		registeredComponentTypes ComponentMask
		registeredComponentNames map[ComponentType]string
		entities                 map[EntityId]*BasicEntity
		components               map[ComponentType][]Component
		systems                  []System
	}
	type args struct {
		ent  *BasicEntity
		comp Component
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantModifiedEntity *BasicEntity
		wantErr            *ECSError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &World{
				registeredComponentTypes: tt.fields.registeredComponentTypes,
				registeredComponentNames: tt.fields.registeredComponentNames,
				entities:                 tt.fields.entities,
				components:               tt.fields.components,
				systems:                  tt.fields.systems,
			}
			gotModifiedEntity, gotErr := w.AddComponentToEntity(tt.args.ent, tt.args.comp)
			if !reflect.DeepEqual(gotModifiedEntity, tt.wantModifiedEntity) {
				t.Errorf("World.AddComponentToEntity() gotModifiedEntity = %v, want %v", gotModifiedEntity, tt.wantModifiedEntity)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("World.AddComponentToEntity() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestWorld_AddSystem(t *testing.T) {
	type fields struct {
		registeredComponentTypes ComponentMask
		registeredComponentNames map[ComponentType]string
		entities                 map[EntityId]*BasicEntity
		components               map[ComponentType][]Component
		systems                  []System
	}
	type args struct {
		s System
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr *ECSError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &World{
				registeredComponentTypes: tt.fields.registeredComponentTypes,
				registeredComponentNames: tt.fields.registeredComponentNames,
				entities:                 tt.fields.entities,
				components:               tt.fields.components,
				systems:                  tt.fields.systems,
			}
			if gotErr := w.AddSystem(tt.args.s); !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("World.AddSystem() = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestWorld_CreateEntity(t *testing.T) {
	type fields struct {
		registeredComponentTypes ComponentMask
		registeredComponentNames map[ComponentType]string
		entities                 map[EntityId]*BasicEntity
		components               map[ComponentType][]Component
		systems                  []System
	}
	type args struct {
		comps []Component
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BasicEntity
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &World{
				registeredComponentTypes: tt.fields.registeredComponentTypes,
				registeredComponentNames: tt.fields.registeredComponentNames,
				entities:                 tt.fields.entities,
				components:               tt.fields.components,
				systems:                  tt.fields.systems,
			}
			if got := w.CreateEntity(tt.args.comps...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("World.CreateEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorld_FindEntityById(t *testing.T) {
	type fields struct {
		registeredComponentTypes ComponentMask
		registeredComponentNames map[ComponentType]string
		entities                 map[EntityId]*BasicEntity
		components               map[ComponentType][]Component
		systems                  []System
	}
	type args struct {
		id EntityId
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantEntity *BasicEntity
		wantErr    *ECSError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &World{
				registeredComponentTypes: tt.fields.registeredComponentTypes,
				registeredComponentNames: tt.fields.registeredComponentNames,
				entities:                 tt.fields.entities,
				components:               tt.fields.components,
				systems:                  tt.fields.systems,
			}
			gotEntity, gotErr := w.FindEntityById(tt.args.id)
			if !reflect.DeepEqual(gotEntity, tt.wantEntity) {
				t.Errorf("World.FindEntityById() gotEntity = %v, want %v", gotEntity, tt.wantEntity)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("World.FindEntityById() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestWorld_RemoveEntity(t *testing.T) {
	type fields struct {
		registeredComponentTypes ComponentMask
		registeredComponentNames map[ComponentType]string
		entities                 map[EntityId]*BasicEntity
		components               map[ComponentType][]Component
		systems                  []System
	}
	type args struct {
		id EntityId
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &World{
				registeredComponentTypes: tt.fields.registeredComponentTypes,
				registeredComponentNames: tt.fields.registeredComponentNames,
				entities:                 tt.fields.entities,
				components:               tt.fields.components,
				systems:                  tt.fields.systems,
			}
			w.RemoveEntity(tt.args.id)
		})
	}
}

func TestWorld_Tick(t *testing.T) {
	type fields struct {
		registeredComponentTypes ComponentMask
		registeredComponentNames map[ComponentType]string
		entities                 map[EntityId]*BasicEntity
		components               map[ComponentType][]Component
		systems                  []System
	}
	type args struct {
		dt int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &World{
				registeredComponentTypes: tt.fields.registeredComponentTypes,
				registeredComponentNames: tt.fields.registeredComponentNames,
				entities:                 tt.fields.entities,
				components:               tt.fields.components,
				systems:                  tt.fields.systems,
			}
			w.Tick(tt.args.dt)
		})
	}
}

func TestWorld_Input(t *testing.T) {
	type fields struct {
		registeredComponentTypes ComponentMask
		registeredComponentNames map[ComponentType]string
		entities                 map[EntityId]*BasicEntity
		components               map[ComponentType][]Component
		systems                  []System
	}
	type args struct {
		e tcell.EventKey
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &World{
				registeredComponentTypes: tt.fields.registeredComponentTypes,
				registeredComponentNames: tt.fields.registeredComponentNames,
				entities:                 tt.fields.entities,
				components:               tt.fields.components,
				systems:                  tt.fields.systems,
			}
			w.Input(tt.args.e)
		})
	}
}

func TestWorld_FindEntitiesWithComponents(t *testing.T) {
	type fields struct {
		registeredComponentTypes ComponentMask
		registeredComponentNames map[ComponentType]string
		entities                 map[EntityId]*BasicEntity
		components               map[ComponentType][]Component
		systems                  []System
	}
	type args struct {
		comps ComponentMask
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*BasicEntity
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &World{
				registeredComponentTypes: tt.fields.registeredComponentTypes,
				registeredComponentNames: tt.fields.registeredComponentNames,
				entities:                 tt.fields.entities,
				components:               tt.fields.components,
				systems:                  tt.fields.systems,
			}
			if got := w.FindEntitiesWithComponents(tt.args.comps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("World.FindEntitiesWithComponents() = %v, want %v", got, tt.want)
			}
		})
	}
}
