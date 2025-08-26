package commands

import (
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-editor/grid"
)

type ToggleGridLines struct {
	world *ecs.World
}

func NewToggleGridLines(world *ecs.World) *ToggleGridLines {
	return &ToggleGridLines{
		world: world,
	}
}

func (c *ToggleGridLines) Execute() error {
	system, exists := c.world.GetSystem(grid.GridLineRendererType)

	if !exists {
		return nil
	}

	if system.IsEnabled() {
		system.Disable()
	} else {
		system.Enable()
	}

	return nil
}
