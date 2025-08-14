package editor

import (
	"path/filepath"
	"strings"

	fin "github.com/adm87/finch-application/application"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/linq"
	"github.com/adm87/finch-editor/camera"
	"github.com/adm87/finch-editor/grid"
	"github.com/adm87/finch-resources/manifest"
)

func Initialize(app *fin.Application, world *ecs.World) error {
	if err := LoadDefaultResources(app); err != nil {
		return err
	}
	if _, err := camera.NewCameraEntity(world); err != nil {
		return err
	}
	if _, err := grid.NewEditorGridEntity(world); err != nil {
		return err
	}
	return nil
}

func LoadDefaultResources(app *fin.Application) error {
	embeddedManifest, err := manifest.GetSubManifest(app.Cache().Manifest(), "embedded")
	if err != nil {
		return err
	}
	names := linq.SelectKeys(embeddedManifest, func(key string, value manifest.ResourceMetadata) bool {
		parts := strings.Split(value.Path, string(filepath.Separator))
		return len(parts) > 1 && parts[0] == "defaults"
	})
	return app.Cache().Load(names...)
}
