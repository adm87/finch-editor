package application

import (
	"image/color"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	finch "github.com/adm87/finch-application/application"
	"github.com/adm87/finch-application/config"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/linq"
	"github.com/adm87/finch-editor/data"
	"github.com/adm87/finch-editor/editor"
	"github.com/adm87/finch-resources/manifest"
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
			ClearColor:      color.RGBA{R: 29, G: 33, B: 41, A: 255},
		},
	}).
	WithStartup(Start).
	WithShutdown(Shutdown).
	WithDraw(Draw).
	WithUpdate(Update)

func Start(app *finch.Application) error {
	if err := AddResourceFilesystems(app, world); err != nil {
		return err
	}
	if err := LoadDefaultResources(app); err != nil {
		return err
	}
	if err := editor.Register(app, world); err != nil {
		return err
	}
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

func AddResourceFilesystems(app *finch.Application, world *ecs.World) error {
	resourcePath := app.Config().Resources.Path
	return app.Cache().AddFilesystems(map[string]fs.FS{
		"embedded": data.Embedded,
		"assets":   os.DirFS(filepath.Join(resourcePath, "assets")),
	})
}

func LoadDefaultResources(app *finch.Application) error {
	embeddedManifest, err := manifest.GetSubManifest(app.Cache().Manifest(), "embedded")
	if err != nil {
		return err
	}
	names := linq.SelectKeys(embeddedManifest, func(key string, value manifest.ResourceMetadata) bool {
		parts := strings.Split(value.Path, string(filepath.Separator))
		return len(parts) > 1 && parts[0] == "defaults"
	})
	if err := app.Cache().Load(names...); err != nil {
		return err
	}
	for _, name := range names {
		if err := app.Cache().SetFallback(name); err != nil {
			return err
		}
	}
	return nil
}
