package grid

import "github.com/adm87/finch-core/ecs"

type GridLineToggle struct {
	world *ecs.World
}

func NewGridLineToggle(world *ecs.World) *GridLineToggle {
	return &GridLineToggle{
		world: world,
	}
}

func (t *GridLineToggle) Execute() error {
	system, exists := t.world.GetSystem(GridLineRendererType)

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
