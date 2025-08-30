package commands

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/errors"
	"github.com/adm87/finch-core/fsys"
	"github.com/adm87/finch-editor/tilemaps"
	"github.com/sqweek/dialog"

	tm "github.com/adm87/finch-tilemap/tilemaps"
)

type PromptLoadTilemap struct {
	world        *ecs.World
	resourcePath string
}

func NewPromptLoadTilemap(resourcePath string, world *ecs.World) *PromptLoadTilemap {
	return &PromptLoadTilemap{
		world:        world,
		resourcePath: resourcePath,
	}
}

func (c *PromptLoadTilemap) Execute() error {
	path, err := get_tilemap_path(c.resourcePath)
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

	filename := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	if filename == "" {
		return errors.NewConflictError("invalid tilemap file")
	}

	tilemapEditor, err := tilemaps.FindTilemapEditorComponent(c.world)
	if err != nil {
		return err
	}
	tilemapEditor.LoadedTilemap = filename

	tilemap, err := tm.Cache().Get(filename)
	if err != nil {
		return err
	}

	tilemap.IsDirty = true

	return nil
}

func get_tilemap_path(startDir string) (string, error) {
	return dialog.
		File().
		SetStartDir(startDir).
		Title("Load Tilemap").
		Filter("Tilemap Files", "tilemap").
		Load()
}
