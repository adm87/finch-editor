package tilemaps

import (
	tm "github.com/adm87/finch-tilemap/tilemaps"
	ts "github.com/adm87/finch-tilemap/tilesets"
)

func get_tilemap_info(tilemapID string) (*tm.Tilemap, *ts.Tileset, error) {
	tilemap, err := tm.Cache().Get(tilemapID)
	if err != nil {
		return nil, nil, err
	}
	tileset, err := ts.Cache().Get(tilemap.TilesetID)
	if err != nil {
		return nil, nil, err
	}
	return tilemap, tileset, nil
}

func set_tile_at(pos int, tileID int, tilemap *tm.Tilemap) {
	tilemap.Data[pos] = tileID
	tilemap.IsDirty = true
}
