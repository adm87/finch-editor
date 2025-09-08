package tilemaps

import (
	finch "github.com/adm87/finch-application/application"
	"github.com/adm87/finch-core/ecs"

	tm "github.com/adm87/finch-tilemap/tilemaps"
	ts "github.com/adm87/finch-tilemap/tilesets"
)

type NewTilemapCommand struct {
	app     *finch.Application
	world   *ecs.World
	Rows    int
	Columns int
}

func NewNewTilemapCommand(app *finch.Application, world *ecs.World, rows, columns int) *NewTilemapCommand {
	return &NewTilemapCommand{
		app:     app,
		world:   world,
		Rows:    rows,
		Columns: columns,
	}
}

func (c *NewTilemapCommand) Execute() error {
	tilemapComp, err := FindTilemapComponent(c.world)
	if err != nil {
		return err
	}

	tm.Storage().Put(DefaultTilemapID, tm.NewTilemap(
		c.Rows,
		c.Columns,
		ts.Storage().DefaultKey(),
	))
	tilemapComp.TilemapID = DefaultTilemapID

	editorComp, err := FindTilemapEditorComponent(c.world)
	if err != nil {
		return err
	}
	editorComp.IsDirty = false // The editor system should handle this, but because the id hasn't changed, it won't.

	c.app.CommandStack().Clear()
	return nil
}
