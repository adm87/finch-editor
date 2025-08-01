package editor

import (
	"image/color"

	"github.com/adm87/finch-application/application"
	"github.com/adm87/finch-application/config"
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
			ClearColor:   color.RGBA{R: 100, G: 149, B: 237, A: 255},
		},
	}).
	WithStartup(Start).
	WithShutdown(Shutdown)

func Start(app *application.Application) error {
	// TODO: start the editor application
	return nil
}

func Shutdown(app *application.Application) error {
	// TODO: shutdown the editor application
	return nil
}
