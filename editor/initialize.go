package editor

import (
	finch "github.com/adm87/finch-application/application"
	"github.com/adm87/finch-core/components/transform"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-editor/editor/components"
)

func Initialize(app *finch.Application, world *ecs.ECSWorld) error {
	if err := internal_setup_camera(world); err != nil {
		return err
	}
	if err := internal_setup_test_sprite(world); err != nil {
		return err
	}
	return nil
}

func internal_setup_camera(world *ecs.ECSWorld) error {
	entity, err := world.NewEntityWithComponents(
		components.NewCameraComponent(),
		components.NewGridComponent(),
	)
	cameraComponent, _, _ := ecs.GetComponent[*components.CameraComponent](world, entity, components.CameraComponentType)
	cameraComponent.ZoomFactor = 0.1
	return err
}

func internal_setup_test_sprite(world *ecs.ECSWorld) error {
	_, err := world.NewEntityWithComponents(
		transform.NewTransformComponent(),
	)
	return err
}
