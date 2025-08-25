package editor

import (
	"math/rand"

	finapp "github.com/adm87/finch-application/application"
	"github.com/adm87/finch-core/components/transform"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/linq"
	"github.com/adm87/finch-editor/camera"
	"github.com/adm87/finch-editor/grid"
	"github.com/adm87/finch-editor/selection"
	"github.com/adm87/finch-resources/storage"
	"github.com/adm87/finch-tilemap/tilemaps"
	"github.com/adm87/finch-tilemap/tilesets"
)

func Initialize(app *finapp.Application, world *ecs.World) error {
	if err := storage.Load("tileset_0000", "tilemap_0000", "img_tileset_0000"); err != nil {
		return err
	}
	if _, err := grid.NewGridLines(world); err != nil {
		return err
	}
	if _, err := selection.NewSelectionBox(world); err != nil {
		return err
	}
	if _, err := camera.NewCamera(world); err != nil {
		return err
	}
	if _, err := new_tilemap(world); err != nil {
		return err
	}
	return nil
}

func new_tilemap(world *ecs.World) (ecs.Entity, error) {
	tilemap, err := tilemaps.Cache().Get("tilemap_0000")
	if err != nil {
		return ecs.NilEntity, err
	}
	tileset, err := tilesets.Cache().Get(tilemap.TilesetID)
	if err != nil {
		return ecs.NilEntity, err
	}

	tilesetMax := tileset.Rows * tileset.Columns

	linq.FillFunc(tilemap.Data, func(i int) int {
		return rand.Intn(tilesetMax)
	})
	tilemap.IsDirty = true
	return world.NewEntityWithComponents(
		tilemaps.NewTilemapComponent("tilemap_0000", 0),
		transform.NewTransformComponent(),
	)
}
