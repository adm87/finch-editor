package editor

import (
	finch "github.com/adm87/finch-application/application"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-editor/editor/systems"
)

func Register(app *finch.Application, world *ecs.ECSWorld) error {
	if err := RegisterSystems(app, world); err != nil {
		return err
	}
	return nil
}

func RegisterSystems(app *finch.Application, world *ecs.ECSWorld) error {
	return world.RegisterSystems(map[ecs.System]int{
		// =================================================================
		// Early Update Systems
		// =================================================================
		systems.NewCameraZoom(): 0,
		systems.NewCameraPan():  1,

		// =================================================================
		// Fixed Update Systems
		// =================================================================

		// =================================================================
		// Late Update Systems
		// =================================================================

		// =================================================================
		// Rendering Systems
		// =================================================================
		systems.NewGridRenderingSystem(app.Config().Window): 0,
	})
}
