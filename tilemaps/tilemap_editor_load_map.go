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
	tilemapComp, err := FindTilemapComponent(c.world)
	if err != nil {
		return err
	}
	c.oldTilemapID = tilemapComp.TilemapID
	tilemapComp.TilemapID = c.newTilemapID
	return nil
}

func (c *TilemapEditorLoadMap) Undo() error {
	tilemapComp, err := FindTilemapComponent(c.world)
	if err != nil {
		return err
	}
	tilemapComp.TilemapID = c.oldTilemapID
	return nil
}

func (c *TilemapEditorLoadMap) Redo() error {
	tilemapComp, err := FindTilemapComponent(c.world)
	if err != nil {
		return err
	}
	tilemapComp.TilemapID = c.newTilemapID
	return nil
}
