package commands

import (
	"fmt"
	"path/filepath"
	"strings"

	finch "github.com/adm87/finch-application/application"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/errors"
	"github.com/adm87/finch-core/fsys"
	"github.com/adm87/finch-editor/tilemaps"
	"github.com/sqweek/dialog"
)

type PromptLoadTilemap struct {
	app          *finch.Application
	world        *ecs.World
	resourcePath string
}

func NewPromptLoadTilemap(app *finch.Application, world *ecs.World) *PromptLoadTilemap {
	return &PromptLoadTilemap{
		app:          app,
		world:        world,
		resourcePath: app.Config().Resources.Path,
	}
}

func (c *PromptLoadTilemap) Execute() error {
	path, err := get_tilemap_load_path(c.resourcePath)
	if err != nil {
		if err.Error() == "Cancelled" {
			return nil
		}
		return err
	}
	path = filepath.ToSlash(path)

	if !fsys.Exists(path) {
		return errors.NewNotFoundError(fmt.Sprintf("file does not exist at path: %s", path))
	}

	c.resourcePath = filepath.Dir(path)

	tilemapID := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	if tilemapID == "" {
		return errors.NewConflictError("invalid tilemap file")
	}

	return c.app.CommandStack().ExecuteCommand(tilemaps.NewLoadMapCommand(c.world, tilemapID))
}

func get_tilemap_load_path(startDir string) (string, error) {
	return dialog.
		File().
		SetStartDir(startDir).
		Title("Load Tilemap").
		Filter("Tilemap Files", "tilemap").
		Load()
}
