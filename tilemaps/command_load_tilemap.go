package tilemaps

import (
	finch "github.com/adm87/finch-application/application"
	"github.com/adm87/finch-core/ecs"
)

type TilemapEditorLoadMap struct {
	app          *finch.Application
	world        *ecs.World
	newTilemapID string
}

func NewLoadMapCommand(app *finch.Application, world *ecs.World, newTilemapID string) *TilemapEditorLoadMap {
	return &TilemapEditorLoadMap{
		app:          app,
		world:        world,
		newTilemapID: newTilemapID,
	}
}

func (c *TilemapEditorLoadMap) Execute() error {
	tilemapComp, err := FindTilemapComponent(c.world)
	if err != nil {
		return err
	}
	tilemapComp.TilemapID = c.newTilemapID

	c.app.CommandStack().Clear()
	return nil
}
