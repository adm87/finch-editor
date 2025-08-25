package commands

import (
	finch "github.com/adm87/finch-application/application"
)

type CloseApplication struct {
	App *finch.Application
}

func NewCloseApplication(app *finch.Application) *CloseApplication {
	return &CloseApplication{
		App: app,
	}
}

func (c *CloseApplication) Execute() error {
	c.App.Close()
	return nil
}
