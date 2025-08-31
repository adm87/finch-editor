package tilemaps

import tm "github.com/adm87/finch-tilemap/tilemaps"

type TilemapEditorPlaceTile struct {
	tilemapID string
	newTileID int
	oldTileID int
	tilePos   int
}

func NewPlaceTileCommand(tilemapID string, tileID int, position int) *TilemapEditorPlaceTile {
	return &TilemapEditorPlaceTile{
		tilemapID: tilemapID,
		newTileID: tileID,
		oldTileID: tileID,
		tilePos:   position,
	}
}

func (c *TilemapEditorPlaceTile) Execute() error {
	tilemap, err := tm.Cache().Get(c.tilemapID)
	if err != nil {
		return err
	}
	c.oldTileID = tilemap.Data[c.tilePos]
	c.set_tile_at(c.tilePos, c.newTileID, tilemap)
	return nil
}

func (c *TilemapEditorPlaceTile) Undo() error {
	tilemap, err := tm.Cache().Get(c.tilemapID)
	if err != nil {
		return err
	}
	c.set_tile_at(c.tilePos, c.oldTileID, tilemap)
	return nil
}

func (c *TilemapEditorPlaceTile) Redo() error {
	tilemap, err := tm.Cache().Get(c.tilemapID)
	if err != nil {
		return err
	}
	c.set_tile_at(c.tilePos, c.newTileID, tilemap)
	return nil
}

func (c *TilemapEditorPlaceTile) set_tile_at(i int, tile int, tilemap *tm.Tilemap) {
	tilemap.Data[i] = tile
	tilemap.IsDirty = true
}
