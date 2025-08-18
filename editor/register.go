package editor

import (
	finapp "github.com/adm87/finch-application/application"
	finmsg "github.com/adm87/finch-application/messages"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/keys"
	"github.com/adm87/finch-editor/camera"
	"github.com/adm87/finch-editor/debug"
	"github.com/adm87/finch-editor/grid"
	"github.com/adm87/finch-editor/selection"
	"github.com/adm87/finch-rendering/rendering"
	"github.com/hajimehoshi/ebiten/v2"
)

func Register(app *finapp.Application, world *ecs.World) error {
	if err := RegisterECSSystems(app, world); err != nil {
		return err
	}
	if err := RegisterDebugSystems(app, world); err != nil {
		return err
	}
	if err := RegisterMessageHandlers(app, world); err != nil {
		return err
	}
	if err := RegisterKeyCommands(app, world); err != nil {
		return err
	}
	return nil
}

func RegisterECSSystems(app *finapp.Application, world *ecs.World) error {
	return world.RegisterSystems(map[ecs.System]int{
		// =================================================================
		// Early Update Systems
		// =================================================================
		selection.NewSelectionBoxUpdate(): -2,
		selection.NewSelectionSystem():    -1,

		camera.NewCameraDrag(): 0,
		camera.NewCameraPan():  1,
		camera.NewCameraZoom(): 2,

		// =================================================================
		// Fixed Update Systems
		// =================================================================

		// =================================================================
		// Late Update Systems
		// =================================================================

		// =================================================================
		// Rendering Systems
		// =================================================================
		grid.NewGridLineRenderer(app.Config().Window): -1000,

		rendering.NewRenderSystem(): 0,
	})
}

func RegisterDebugSystems(app *finapp.Application, world *ecs.World) error {
	return world.RegisterSystems(map[ecs.System]int{
		debug.NewDebugBoundsRenderer(): 10000,
	})
}

func RegisterMessageHandlers(app *finapp.Application, world *ecs.World) error {
	if err := finmsg.ApplicationResize.Subscribe(camera.NewCameraResizeHandler(world)); err != nil {
		return err
	}
	return nil
}

func RegisterKeyCommands(app *finapp.Application, world *ecs.World) error {
	if err := keys.RegisterAction(ebiten.KeyF1, keys.KeyPhaseRelease, grid.NewGridLineToggle(world)); err != nil {
		return err
	}
	if err := keys.RegisterAction(ebiten.KeyF2, keys.KeyPhaseRelease, debug.NewDebugBoundsToggle(world)); err != nil {
		return err
	}
	return nil
}
