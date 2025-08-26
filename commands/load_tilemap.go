package commands

import (
	"path/filepath"

	finch "github.com/adm87/finch-application/application"
	"github.com/sqweek/dialog"
)

type LoadTilemap struct {
	app      *finch.Application
	lastOpen string
}

func NewLoadTilemap(app *finch.Application) *LoadTilemap {
	return &LoadTilemap{
		app: app,
	}
}

func (c *LoadTilemap) Execute() error {
	startDir := c.app.Config().Resources.Path
	if c.lastOpen != "" {
		startDir = c.lastOpen
	}

	path, err := dialog.
		File().
		SetStartDir(startDir).
		Title("Load Tilemap").
		Filter("Tilemap Files", "tilemap").
		Load()
	if err != nil {
		if err.Error() == "Cancelled" {
			return nil
		}
		return err
	}

	c.lastOpen = filepath.Dir(path)
	return nil
}
