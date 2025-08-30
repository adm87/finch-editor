package editor

import (
	finch "github.com/adm87/finch-application/application"

	"github.com/adm87/finch-application/messages"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/keys"
	"github.com/adm87/finch-editor/camera"
	"github.com/adm87/finch-editor/commands"
	"github.com/adm87/finch-editor/grid"
	"github.com/adm87/finch-editor/tilemaps"
	"github.com/adm87/finch-rendering/rendering"
	"github.com/hajimehoshi/ebiten/v2"
)

func Register(app *finch.Application, world *ecs.World) error {
	if err := RegisterECSSystems(app, world); err != nil {
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

func RegisterECSSystems(app *finch.Application, world *ecs.World) error {
	return world.RegisterSystems(map[ecs.System]int{
		// =================================================================
		// Early Update Systems
		// =================================================================
		camera.NewCameraDrag(): 0,
		camera.NewCameraPan():  1,
		camera.NewCameraZoom(): 2,

		tilemaps.NewTilemapEditorSystem(app): 3,

		// =================================================================
		// Fixed Update Systems
		// =================================================================

		// =================================================================
		// Late Update Systems
		// =================================================================

		// =================================================================
		// Rendering Systems
		// =================================================================
		grid.NewGridLineRenderer(app.Config().Window): -1,

		rendering.NewRenderingSystem(): 0,

		tilemaps.NewTilemapEditorRender(): 1,
	})
}

func RegisterMessageHandlers(app *finch.Application, world *ecs.World) error {
	if err := messages.ApplicationResize.Subscribe(camera.NewCameraResizeHandler(world)); err != nil {
		return err
	}
	return nil
}

func RegisterKeyCommands(app *finch.Application, world *ecs.World) error {
	if err := keys.RegisterAction(ebiten.KeyEscape, keys.KeyPhaseRelease, commands.NewCloseApplication(app)); err != nil {
		return err
	}
	if err := keys.RegisterAction(ebiten.KeyF5, keys.KeyPhaseRelease, commands.NewPromptSaveTilemap(app)); err != nil {
		return err
	}
	if err := keys.RegisterAction(ebiten.KeyF8, keys.KeyPhaseRelease, commands.NewPromptLoadTilemap(app.Config().Resources.Path, world)); err != nil {
		return err
	}
	if err := keys.RegisterAction(ebiten.KeyF9, keys.KeyPhaseRelease, commands.NewToggleGridLines(world)); err != nil {
		return err
	}
	if err := keys.RegisterAction(ebiten.KeyF11, keys.KeyPhaseRelease, commands.NewFullscreenToggle(app.Config().Window)); err != nil {
		return err
	}
	return nil
}
