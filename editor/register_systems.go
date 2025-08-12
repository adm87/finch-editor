package editor

import (
	finch "github.com/adm87/finch-application/application"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-editor/editor/camera"
	"github.com/adm87/finch-editor/editor/grid"
	"github.com/adm87/finch-rendering/rendering"
)

func RegisterSystems(app *finch.Application, world *ecs.ECSWorld) error {
	if err := internal_register_late_updates(app, world); err != nil {
		return err
	}
	if err := internal_register_renders(app, world); err != nil {
		return err
	}
	return nil
}

func internal_register_late_updates(app *finch.Application, world *ecs.ECSWorld) error {
	return world.RegisterSystems(map[ecs.System]int{
		camera.NewCameraFinalizationUpdate(app.Config().Window): 1000, // Update camera component late into the LateUpdate phase
	})
}

func internal_register_renders(app *finch.Application, world *ecs.ECSWorld) error {
	return world.RegisterSystems(map[ecs.System]int{
		grid.NewGridRenderingSystem(app.Config().Window): 0, // Render editor grid early in the Render phase
		rendering.NewRenderSystem():                      1,
	})
}
