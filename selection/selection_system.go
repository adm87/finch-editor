package selection

import "github.com/adm87/finch-core/ecs"

var SelectionSystemType = ecs.NewSystemType[*SelectionSystem]()

type SelectionSystem struct {
	enabled bool
}

func NewSelectionSystem() *SelectionSystem {
	return &SelectionSystem{
		enabled: true,
	}
}

func (s *SelectionSystem) Type() ecs.SystemType {
	return SelectionSystemType
}

func (s *SelectionSystem) Enable() {
	s.enabled = true
}

func (s *SelectionSystem) Disable() {
	s.enabled = false
}

func (s *SelectionSystem) IsEnabled() bool {
	return s.enabled
}

func (s *SelectionSystem) EarlyUpdate(world *ecs.World, deltaSeconds float64) error {
	return nil
}
