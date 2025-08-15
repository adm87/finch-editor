package debug

import "github.com/adm87/finch-core/ecs"

type DebugBoundsToggle struct {
	world *ecs.World
}

func NewDebugBoundsToggle(world *ecs.World) *DebugBoundsToggle {
	return &DebugBoundsToggle{
		world: world,
	}
}

func (t *DebugBoundsToggle) Execute() error {
	system, exists := t.world.GetSystem(DebugBoundsRendererType)

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
