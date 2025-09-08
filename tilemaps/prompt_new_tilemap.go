package tilemaps

import (
	finch "github.com/adm87/finch-application/application"
	"github.com/adm87/finch-core/ecs"
)

type PromptNewTilemap struct {
	app   *finch.Application
	world *ecs.World
}

func NewPromptNewTilemap(app *finch.Application, world *ecs.World) *PromptNewTilemap {
	return &PromptNewTilemap{
		app:   app,
		world: world,
	}
}

func (c *PromptNewTilemap) Execute() error {
	unsaved, err := check_unsaved_changes(c.world)
	if err != nil {
		return err
	}
	if unsaved && !confirm_discard_changes() {
		return nil
	}
	return c.app.CommandStack().ExecuteCommand(NewNewTilemapCommand(c.app, c.world, 10, 20))
}
