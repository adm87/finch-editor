package commands

import (
	"path/filepath"

	finch "github.com/adm87/finch-application/application"
	"github.com/sqweek/dialog"
)

type SaveTilemap struct {
	app      *finch.Application
	lastOpen string
}

func NewSaveTilemap(app *finch.Application) *SaveTilemap {
	return &SaveTilemap{
		app: app,
	}
}

func (c *SaveTilemap) Execute() error {
	startDir := c.app.Config().Resources.Path
	if c.lastOpen != "" {
		startDir = c.lastOpen
	}

	path, err := dialog.
		File().
		SetStartDir(startDir).
		Title("Save Tilemap").
		Filter("Tilemap Files", "tilemap").
		Save()
	if err != nil {
		if err.Error() == "Cancelled" {
			return nil
		}
		return err
	}

	c.lastOpen = filepath.Dir(path)
	return nil
}
