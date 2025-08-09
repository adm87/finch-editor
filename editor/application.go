package editor

import (
	"image/color"
	"os"
	"path/filepath"

	"github.com/adm87/finch-application/application"
	"github.com/adm87/finch-application/config"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/errors"
	"github.com/adm87/finch-editor/systems"
	"github.com/adm87/finch-rendering/rendering"
	"github.com/hajimehoshi/ebiten/v2"
)

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
	WithShutdown(Shutdown)

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
	// Register Rendering Systems
	if _, err := app.World().RegisterSystems(map[ecs.System]int{
		systems.NewEditorGridRenderer(app.World(), app.Config().Window): 0,
		rendering.NewRenderSystem():                                     1,
	}); err != nil {
		return err
	}

	return nil
}

func SetupElements(app *application.Application) error {
	return nil
}
