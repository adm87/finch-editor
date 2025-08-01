package editor

import (
	"image/color"

	"github.com/adm87/finch-application/application"
	"github.com/adm87/finch-application/config"
	"github.com/adm87/finch-common/camera"
	"github.com/adm87/finch-common/transform"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-editor/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

var Application = application.NewApplicationWithConfig(
	&application.ApplicationConfig{
		Metadata: &config.Metadata{
			Name:    "Finch Editor",
			Version: "0.0.0-unreleased",
			Root:    ".",
		},
		Resources: &config.Resources{
			Path:         "data/",
			ManifestName: "manifest.json",
		},
		Window: &config.Window{
			Title:        "Finch Editor",
			Width:        800,
			Height:       600,
			ScreenWidth:  800,
			ScreenHeight: 600,
			ResizeMode:   ebiten.WindowResizingModeEnabled,
			RenderScale:  1.0,
			Fullscreen:   false,
			ClearColor:   color.RGBA{R: 30, G: 30, B: 30, A: 255},
		},
	}).
	WithStartup(Start).
	WithShutdown(Shutdown)

func Start(app *application.Application) error {
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

func RegisterSystems(app *application.Application) error {
	// Register Update Systems
	if _, err := app.World().RegisterSystems(map[ecs.System]int{
		camera.NewCameraLateUpdate(app.World(), app.Config().Window): 1000,
	}); err != nil {
		return err
	}

	// Register Rendering Systems
	if _, err := app.World().RegisterSystems(map[ecs.System]int{
		systems.NewEditorGridRenderer(app.World(), app.Config().Window): -1,
	}); err != nil {
		return err
	}

	return nil
}

func SetupElements(app *application.Application) error {
	cameraEntity, err := ecs.NewEntity().AddComponents(
		camera.NewCameraComponent(),
		transform.NewTransformComponent(),
	)
	if err != nil {
		return err
	}
	if _, err := app.World().AddEntities(cameraEntity); err != nil {
		return err
	}
	return nil
}
