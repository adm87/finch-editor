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
	tilemap.Data[c.tilePos] = c.newTileID
	tilemap.IsDirty = true
	return nil
}

func (c *TilemapEditorPlaceTile) Undo() error {
	tilemap, err := tm.Cache().Get(c.tilemapID)
	if err != nil {
		return err
	}
	tilemap.Data[c.tilePos] = c.oldTileID
	tilemap.IsDirty = true
	return nil
}

func (c *TilemapEditorPlaceTile) Redo() error {
	return c.Execute()
}
