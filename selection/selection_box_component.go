package selection

import (
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/geometry"
	"github.com/adm87/finch-core/types"
)

var SelectionBoxComponentType = ecs.NewComponentType[*SelectionBoxComponent]()

type SelectionBoxComponent struct {
	SelectionStartPoint types.Optional[geometry.Point64]
	SelectionEndPoint   types.Optional[geometry.Point64]
}

func NewSelectionBoxComponent() *SelectionBoxComponent {
	return &SelectionBoxComponent{
		SelectionStartPoint: types.NewEmptyOptional[geometry.Point64](),
		SelectionEndPoint:   types.NewEmptyOptional[geometry.Point64](),
	}
}

func (c *SelectionBoxComponent) Type() ecs.ComponentType {
	return SelectionBoxComponentType
}

func (c *SelectionBoxComponent) Dispose() {
}
