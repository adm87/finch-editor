package editor

import (
	"image/color"
	"os"
	"path/filepath"

	"github.com/adm87/finch-application/application"
	"github.com/adm87/finch-application/config"
	"github.com/adm87/finch-core/components/transform"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/errors"
	"github.com/adm87/finch-core/geometry"
	"github.com/adm87/finch-editor/components"
	"github.com/adm87/finch-editor/systems"
	"github.com/adm87/finch-rendering/rendering"
	"github.com/adm87/finch-rendering/sprites"
	"github.com/hajimehoshi/ebiten/v2"
)

var world = ecs.NewWorld()

var Application = application.NewApplicationWithConfig(
	&application.ApplicationConfig{
		Metadata: &config.Metadata{
			Name:      "Finch Editor",
			Version:   "0.0.0-unreleased",
			Root:      ".",
			TargetFps: 60,
		},
		Resources: &config.Resources{
			Path:         "data/",
			ManifestName: "manifest.json",
		},
		Window: &config.Window{
			Title:           "Finch Editor",
			Width:           800,
			Height:          600,
			ResizeMode:      ebiten.WindowResizingModeEnabled,
			RenderScale:     1.0,
			Fullscreen:      false,
			ClearBackground: true,
			ClearColor:      color.RGBA{R: 30, G: 30, B: 30, A: 255},
		},
	}).
	WithStartup(Start).
	WithShutdown(Shutdown).
	WithDraw(Draw).
	WithUpdate(Update)

func Start(app *application.Application) error {
	if err := RegisterApplicationResources(app); err != nil {
		return err
	}
	if err := RegisterSystems(app); err != nil {
		return err
	}
	if err := SetupElements(app); err != nil {
		return err
	}
	return nil
}

func Shutdown(app *application.Application) error {
	// TODO: shutdown the editor application
	return nil
}

func RegisterApplicationResources(app *application.Application) error {
	resources := app.Config().Resources
	if resources == nil {
		return errors.NewNotFoundError("application resource config not found")
	}
	if err := app.Cache().AddFilesystem("assets", os.DirFS(filepath.Join(resources.Path, "assets"))); err != nil {
		return err
	}
	if err := app.Cache().Load("tile_0000"); err != nil {
		return err
	}
	return nil
}

func RegisterSystems(app *application.Application) error {
	// Register EarlyUpdate systems

	// Register LateUpdate systems
	if err := world.RegisterSystems(map[ecs.System]int{
		systems.NewCameraLateUpdateSystem(app): 0,
	}); err != nil {
		return err
	}

	// Register Rendering systems
	if err := world.RegisterSystems(map[ecs.System]int{
		systems.NewEditorGridRenderer(app.Config().Window): 0,
		rendering.NewRenderSystem():                        1,
	}); err != nil {
		return err
	}

	return nil
}

func SetupElements(app *application.Application) error {
	img, err := app.Cache().Images().Get("tile_0000")
	if err != nil {
		return err
	}
	tile0000Renderer := sprites.NewSpriteRenderer(img, geometry.Point64{X: 0.5, Y: 0.5})

	// Test sprite entity
	if _, err := world.NewEntityWithComponents(
		transform.NewTransformComponent(),
		rendering.NewRenderComponent(tile0000Renderer, 0),
	); err != nil {
		return err
	}

	// Camera entity
	if _, err := world.NewEntityWithComponents(
		components.NewCameraComponent(),
	); err != nil {
		return err
	}

	return nil
}

func Draw(app *application.Application, screen *ebiten.Image) error {
	return world.ProcessRenderSystems(screen)
}

func Update(app *application.Application, deltaSeconds, fixedDeltaSeconds float64, frames int) error {
	return world.ProcessUpdateSystems(deltaSeconds, fixedDeltaSeconds, frames)
}
