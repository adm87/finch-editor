package camera

import (
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/geometry"
	"github.com/adm87/finch-core/types"
)

var CameraDragComponentType = ecs.NewComponentType[*CameraDragComponent]()

type CameraDragComponent struct {
	DownPosition       types.Optional[geometry.Point64]
	DragVector         types.Optional[geometry.Point64]
	DragStartThreshold float64
	IsDragging         bool
}

func NewCameraDragComponent() *CameraDragComponent {
	return &CameraDragComponent{
		DownPosition:       types.NewEmptyOptional[geometry.Point64](),
		DragVector:         types.NewEmptyOptional[geometry.Point64](),
		DragStartThreshold: 0.1,
		IsDragging:         false,
	}
}

func (c *CameraDragComponent) Type() ecs.ComponentType {
	return CameraDragComponentType
}
