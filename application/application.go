package application

import (
	"image/color"

	finch "github.com/adm87/finch-application/application"
	"github.com/adm87/finch-application/config"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-editor/editor"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	world = ecs.NewWorld()
)

var Application = finch.NewApplicationWithConfig(
	&finch.ApplicationConfig{
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
			Width:           1240,
			Height:          720,
			ScreenWidth:     800,
			ScreenHeight:    600,
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

func Start(app *finch.Application) error {
	// =================================================================
	// Registration
	// =================================================================

	if err := editor.Register(app, world); err != nil {
		return err
	}

	// =================================================================
	// Initialization
	// =================================================================

	if err := editor.Initialize(app, world); err != nil {
		return err
	}

	return nil
}

func Shutdown(app *finch.Application) error {
	// TODO: shutdown the editor application
	return nil
}

// Draw processes the world and UI render systems
func Draw(app *finch.Application, screen *ebiten.Image) error {
	if err := world.ProcessRenderSystems(screen); err != nil {
		return err
	}
	return nil
}

// Update processes the world and UI update systems
func Update(app *finch.Application, deltaSeconds, fixedDeltaSeconds float64, frames int) error {
	if err := world.ProcessUpdateSystems(deltaSeconds, fixedDeltaSeconds, frames); err != nil {
		return err
	}

	return nil
}
