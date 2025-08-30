package commands

import (
	"path/filepath"

	finch "github.com/adm87/finch-application/application"
	"github.com/sqweek/dialog"
)

type PromptSaveTilemap struct {
	app      *finch.Application
	lastOpen string
}

func NewPromptSaveTilemap(app *finch.Application) *PromptSaveTilemap {
	return &PromptSaveTilemap{
		app: app,
	}
}

func (c *PromptSaveTilemap) Execute() error {
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
