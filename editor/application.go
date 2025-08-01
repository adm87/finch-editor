package editor

import (
	"image/color"

	"github.com/adm87/finch-application/application"
	"github.com/adm87/finch-application/config"
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
	app.World().RegisterSystems(map[ecs.System]int{
		systems.NewEditorGridRenderer(): 0,
	})
	return nil
}

func Shutdown(app *application.Application) error {
	// TODO: shutdown the editor application
	return nil
}
