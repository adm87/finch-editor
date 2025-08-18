package grid

import (
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/errors"
	"github.com/adm87/finch-core/hash"
)

var (
	ErrGridLineEntityNotFound = errors.NewNotFoundError("editor grid line entity")
)

func NewGridLines(world *ecs.World) (ecs.Entity, error) {
	return world.NewEntityWithComponents(
		NewGridLineComponent(),
	)
}

func FindGridLineEntities(world *ecs.World) hash.HashSet[ecs.Entity] {
	return world.FilterEntitiesByComponents(
		GridLineComponentType,
	)
}

func FindGridLineEntity(world *ecs.World) (ecs.Entity, error) {
	entities := FindGridLineEntities(world)
	if entities.IsEmpty() {
		return ecs.NilEntity, ErrGridLineEntityNotFound
	}
	entity, ok := entities.First()
	if !ok {
		return ecs.NilEntity, ErrGridLineEntityNotFound
	}
	return entity, nil
}

func FindGridLineComponent(world *ecs.World) (*GridLineComponent, error) {
	entity, err := FindGridLineEntity(world)
	if err != nil {
		return nil, err
	}
	gridComponent, found, err := ecs.GetComponent[*GridLineComponent](world, entity, GridLineComponentType)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ErrGridLineEntityNotFound
	}
	return gridComponent, nil
}
