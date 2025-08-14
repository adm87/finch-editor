package editor

import (
	finch "github.com/adm87/finch-application/application"
	"github.com/adm87/finch-core/components/camera"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-editor/editor/components"
)

func Initialize(app *finch.Application, world *ecs.ECSWorld) error {
	if err := InitializeCamera(world); err != nil {
		return err
	}
	if err := InitializeGrid(world); err != nil {
		return err
	}
	return nil
}

func InitializeCamera(world *ecs.ECSWorld) error {
	entity, err := world.NewEntityWithComponents(
		camera.NewCameraComponent(),
	)
	cameraComponent, _, _ := ecs.GetComponent[*camera.CameraComponent](world, entity, camera.CameraComponentType)
	cameraComponent.ZoomFactor = 0.1
	return err
}

func InitializeGrid(world *ecs.ECSWorld) error {
	_, err := world.NewEntityWithComponents(
		components.NewGridComponent(),
	)
	return err
}
