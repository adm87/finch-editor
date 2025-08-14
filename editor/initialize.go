package editor

import (
	fin "github.com/adm87/finch-application/application"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-editor/camera"
	"github.com/adm87/finch-editor/grid"
)

func Initialize(app *fin.Application, world *ecs.World) error {
	if _, err := camera.NewCameraEntity(world); err != nil {
		return err
	}
	if _, err := grid.NewEditorGridEntity(world); err != nil {
		return err
	}
	return nil
}
