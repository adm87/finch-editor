package tilemaps

import (
	"github.com/adm87/finch-core/errors"
	tm "github.com/adm87/finch-tilemap/tilemaps"
)

type TilemapEditorTilePlacement struct {
	tilemapID     string
	placementInfo []*TilePlacementInfo
}

func NewTilemapEditorTilePlacement(tilemapID string) *TilemapEditorTilePlacement {
	return &TilemapEditorTilePlacement{
		tilemapID:     tilemapID,
		placementInfo: make([]*TilePlacementInfo, 0),
	}
}

func (c *TilemapEditorTilePlacement) AddPlacement(info *TilePlacementInfo) {
	c.placementInfo = append(c.placementInfo, info)
}

func (c *TilemapEditorTilePlacement) Execute() error {
	if len(c.placementInfo) == 0 {
		return errors.NewInvalidArgumentError("No tile placement information available")
	}
	if c.tilemapID == "" {
		return errors.NewInvalidArgumentError("Tilemap ID is required")
	}
	return nil
}

func (c *TilemapEditorTilePlacement) Undo() error {
	tilemap, err := tm.Cache().Get(c.tilemapID)
	if err != nil {
		return err
	}
	for _, info := range c.placementInfo {
		set_tile_at(info.Position, info.OldTileID, tilemap)
	}
	return nil
}

func (c *TilemapEditorTilePlacement) Redo() error {
	tilemap, err := tm.Cache().Get(c.tilemapID)
	if err != nil {
		return err
	}
	for _, info := range c.placementInfo {
		set_tile_at(info.Position, info.NewTileID, tilemap)
	}
	return nil
}
