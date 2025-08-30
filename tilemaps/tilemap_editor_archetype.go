package tilemaps

import (
	"github.com/adm87/finch-core/components/transform"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/errors"

	tm "github.com/adm87/finch-tilemap/tilemaps"
)

func NewTilemapEditor(world *ecs.World) (ecs.Entity, error) {
	return world.NewEntityWithComponents(
		NewTilemapEditorComponent(),
		tm.NewEmptyTilemapComponent(),
		tm.NewTilemapRenderComponent(0),
		transform.NewTransformComponent(),
	)
}

func FindTilemapEditorEntity(world *ecs.World) (ecs.Entity, error) {
	entities := world.FilterEntitiesByComponents(tm.TilemapComponentType)
	if len(entities) == 0 {
		return ecs.NilEntity, errors.NewNotFoundError("tilemap editor entity not found")
	}
	entity, ok := entities.First()
	if !ok {
		return ecs.NilEntity, errors.NewNotFoundError("tilemap editor entity not found")
	}
	return entity, nil
}

func FindTilemapComponentEntity(world *ecs.World) (ecs.Entity, error) {
	entities := world.FilterEntitiesByComponents(tm.TilemapComponentType)
	if len(entities) == 0 {
		return ecs.NilEntity, errors.NewNotFoundError("tilemap component entity not found")
	}
	entity, ok := entities.First()
	if !ok {
		return ecs.NilEntity, errors.NewNotFoundError("tilemap component entity not found")
	}
	return entity, nil
}

func FindTilemapComponent(world *ecs.World) (*tm.TilemapComponent, error) {
	entity, err := FindTilemapComponentEntity(world)
	if err != nil {
		return nil, err
	}
	component, found, _ := ecs.GetComponent[*tm.TilemapComponent](world, entity, tm.TilemapComponentType)
	if !found {
		return nil, errors.NewNotFoundError("tilemap component not found")
	}
	return component, nil
}

func FindTilemapEditorComponent(world *ecs.World) (*TilemapEditorComponent, error) {
	entity, err := FindTilemapEditorEntity(world)
	if err != nil {
		return nil, err
	}
	component, found, _ := ecs.GetComponent[*TilemapEditorComponent](world, entity, TilemapEditorComponentType)
	if !found {
		return nil, errors.NewNotFoundError("tilemap editor component not found")
	}
	return component, nil
}

func FindTilemapEditorTransform(world *ecs.World) (*transform.TransformComponent, error) {
	entity, err := FindTilemapEditorEntity(world)
	if err != nil {
		return nil, err
	}
	component, found, _ := ecs.GetComponent[*transform.TransformComponent](world, entity, transform.TransformComponentType)
	if !found {
		return nil, errors.NewNotFoundError("tilemap editor transform component not found")
	}
	return component, nil
}
