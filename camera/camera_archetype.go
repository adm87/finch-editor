package camera

import (
	fcam "github.com/adm87/finch-core/components/camera"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/errors"
	"github.com/adm87/finch-core/hash"
)

var (
	ErrCameraNotFound       = errors.NewNotFoundError("camera entity not found")
	ErrMultipleCamerasFound = errors.NewAmbiguousError("multiple camera entities found")
)

func NewCamera(world *ecs.World) (ecs.Entity, error) {
	return world.NewEntityWithComponents(
		fcam.NewCameraComponent(),
		NewCameraDragComponent(),
		NewCameraPanComponent(),
	)
}

func FindCameraEntities(world *ecs.World) hash.HashSet[ecs.Entity] {
	return world.FilterEntitiesByComponents(
		fcam.CameraComponentType,
		CameraDragComponentType,
		CameraPanComponentType,
	)
}

func FindCameraEntity(world *ecs.World) (ecs.Entity, error) {
	entities := FindCameraEntities(world)
	if entities.IsEmpty() {
		return ecs.NilEntity, ErrCameraNotFound
	}
	if len(entities) > 1 {
		return ecs.NilEntity, ErrMultipleCamerasFound
	}
	entity, ok := entities.First()
	if !ok {
		return ecs.NilEntity, ErrCameraNotFound
	}
	return entity, nil
}

func FindCameraComponent(world *ecs.World) (*fcam.CameraComponent, error) {
	entity, err := FindCameraEntity(world)
	if err != nil {
		return nil, err
	}
	cameraComponent, found, err := ecs.GetComponent[*fcam.CameraComponent](world, entity, fcam.CameraComponentType)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ErrCameraNotFound
	}
	return cameraComponent, nil
}

func FindCameraDragComponent(world *ecs.World) (*CameraDragComponent, error) {
	entity, err := FindCameraEntity(world)
	if err != nil {
		return nil, err
	}
	cameraDragComponent, found, err := ecs.GetComponent[*CameraDragComponent](world, entity, CameraDragComponentType)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ErrCameraNotFound
	}
	return cameraDragComponent, nil
}

func FindCameraPanComponent(world *ecs.World) (*CameraPanComponent, error) {
	entity, err := FindCameraEntity(world)
	if err != nil {
		return nil, err
	}
	cameraPanComponent, found, err := ecs.GetComponent[*CameraPanComponent](world, entity, CameraPanComponentType)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ErrCameraNotFound
	}
	return cameraPanComponent, nil
}
