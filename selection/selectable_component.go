package selection

import "github.com/adm87/finch-core/ecs"

var SelectableComponentType = ecs.NewComponentType[*SelectableComponent]()

type SelectableComponent struct {
	IsSelected bool
}

func NewSelectableComponent() *SelectableComponent {
	return &SelectableComponent{
		IsSelected: false,
	}
}

func (s *SelectableComponent) Type() ecs.ComponentType {
	return SelectableComponentType
}

func (s *SelectableComponent) Dispose() {
}
