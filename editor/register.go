package editor

import (
	fapp "github.com/adm87/finch-application/application"
	fmsg "github.com/adm87/finch-application/messages"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-editor/camera"
	"github.com/adm87/finch-editor/grid"
)

func Register(app *fapp.Application, world *ecs.ECSWorld) error {
	if err := RegisterSystems(app, world); err != nil {
		return err
	}
	if err := RegisterMessageHandlers(app, world); err != nil {
		return err
	}
	return nil
}

func RegisterSystems(app *fapp.Application, world *ecs.ECSWorld) error {
	return world.RegisterSystems(map[ecs.System]int{
		// =================================================================
		// Early Update Systems
		// =================================================================
		camera.NewCameraZoom(): 0,
		camera.NewCameraPan():  1,

		// =================================================================
		// Fixed Update Systems
		// =================================================================

		// =================================================================
		// Late Update Systems
		// =================================================================

		// =================================================================
		// Rendering Systems
		// =================================================================
		grid.NewGridRenderingSystem(app.Config().Window): 0,
	})
}

func RegisterMessageHandlers(app *fapp.Application, world *ecs.ECSWorld) error {
	if err := fmsg.ApplicationResize.Subscribe(camera.NewCameraResizeHandler(world)); err != nil {
		return err
	}
	return nil
}
