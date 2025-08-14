package grid

import (
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/errors"
	"github.com/adm87/finch-core/hash"
)

var (
	ErrGridEntityNotFound = errors.NewNotFoundError("editor grid entity")
)

func NewEditorGridEntity(world *ecs.ECSWorld) (ecs.Entity, error) {
	return world.NewEntityWithComponents(
		NewGridComponent(),
	)
}

func FindEditorGridEntities(world *ecs.ECSWorld) hash.HashSet[ecs.Entity] {
	return world.FilterEntitiesByComponents(
		GridComponentType,
	)
}

func FindEditorGridEntity(world *ecs.ECSWorld) (ecs.Entity, error) {
	entities := FindEditorGridEntities(world)
	if entities.IsEmpty() {
		return ecs.NilEntity, ErrGridEntityNotFound
	}
	entity, ok := entities.First()
	if !ok {
		return ecs.NilEntity, ErrGridEntityNotFound
	}
	return entity, nil
}

func FindGridComponent(world *ecs.ECSWorld) (*GridComponent, error) {
	entity, err := FindEditorGridEntity(world)
	if err != nil {
		return nil, err
	}
	gridComponent, found, err := ecs.GetComponent[*GridComponent](world, entity, GridComponentType)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ErrGridEntityNotFound
	}
	return gridComponent, nil
}
