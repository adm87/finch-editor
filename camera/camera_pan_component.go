package camera

import (
	"github.com/adm87/finch-core/ecs"
)

var CameraPanComponentType = ecs.NewComponentType[*CameraPanComponent]()

type CameraPanComponent struct {
	PanSpeed      float64
	QuickPanSpeed float64

	IsPanning bool
}

func NewCameraPanComponent() *CameraPanComponent {
	return &CameraPanComponent{
		PanSpeed:      100.0,
		QuickPanSpeed: 400.0,
	}
}

func (c *CameraPanComponent) Type() ecs.ComponentType {
	return CameraPanComponentType
}
