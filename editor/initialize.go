package editor

import (
	finapp "github.com/adm87/finch-application/application"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-editor/camera"
	"github.com/adm87/finch-editor/grid"
	"github.com/adm87/finch-editor/selection"
)

func Initialize(app *finapp.Application, world *ecs.World) error {
	if _, err := grid.NewGridLines(world); err != nil {
		return err
	}
	if _, err := selection.NewSelectionBox(world); err != nil {
		return err
	}
	if _, err := camera.NewCamera(world); err != nil {
		return err
	}
	return nil
}
