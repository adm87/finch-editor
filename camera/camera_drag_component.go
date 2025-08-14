package camera

import (
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/geometry"
	"github.com/adm87/finch-core/types"
)

var CameraDragComponentType = ecs.NewComponentType[*CameraDragComponent]()

type CameraDragComponent struct {
	DragVector         types.Optional[geometry.Point64]
	DragStartThreshold float64
	IsDragging         bool
}

func NewCameraDragComponent() *CameraDragComponent {
	return &CameraDragComponent{
		DragVector:         types.NewEmptyOption[geometry.Point64](),
		DragStartThreshold: 0.1,
		IsDragging:         false,
	}
}

func (c *CameraDragComponent) Type() ecs.ComponentType {
	return CameraDragComponentType
}

func (c *CameraDragComponent) Dispose() {
	c.DragVector.Invalidate()
	c.IsDragging = false
}
