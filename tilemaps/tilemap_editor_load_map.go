package tilemaps

import (
	"github.com/adm87/finch-core/ecs"
)

type TilemapEditorLoadMap struct {
	world        *ecs.World
	newTilemapID string
	oldTilemapID string
}

func NewLoadMapCommand(world *ecs.World, newTilemapID string) *TilemapEditorLoadMap {
	return &TilemapEditorLoadMap{
		world:        world,
		newTilemapID: newTilemapID,
	}
}

func (c *TilemapEditorLoadMap) Execute() error {
	tilemapEditor, err := FindTilemapEditorComponent(c.world)
	if err != nil {
		return err
	}
	c.oldTilemapID = tilemapEditor.LoadedTilemap
	tilemapEditor.LoadedTilemap = c.newTilemapID
	return nil
}

func (c *TilemapEditorLoadMap) Undo() error {
	tilemapEditor, err := FindTilemapEditorComponent(c.world)
	if err != nil {
		return err
	}
	tilemapEditor.LoadedTilemap = c.oldTilemapID
	return nil
}

func (c *TilemapEditorLoadMap) Redo() error {
	tilemapEditor, err := FindTilemapEditorComponent(c.world)
	if err != nil {
		return err
	}
	tilemapEditor.LoadedTilemap = c.newTilemapID
	return nil
}
