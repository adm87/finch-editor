package tilemaps

import (
	"github.com/adm87/finch-core/ecs"
	tm "github.com/adm87/finch-tilemap/tilemaps"
	ts "github.com/adm87/finch-tilemap/tilesets"
	"github.com/sqweek/dialog"
)

const (
	DefaultTilemapID = "Untitled Tilemap"
)

func get_tilemap_info(tilemapID string) (*tm.Tilemap, *ts.Tileset, error) {
	tilemap, err := tm.Storage().Get(tilemapID)
	if err != nil {
		return nil, nil, err
	}
	tileset, err := ts.Storage().Get(tilemap.TilesetID)
	if err != nil {
		return nil, nil, err
	}
	return tilemap, tileset, nil
}

func check_unsaved_changes(world *ecs.World) (bool, error) {
	editorComp, err := FindTilemapEditorComponent(world)
	if err != nil {
		return false, err
	}
	if editorComp.IsDirty {
		return true, nil
	}
	return false, nil
}

func confirm_discard_changes() bool {
	return dialog.Message("You have unsaved changes. Discard changes?").Title("Unsaved Changes").YesNo()
}
