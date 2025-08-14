package components

import (
	"github.com/adm87/finch-core/components/transform"
	"github.com/adm87/finch-core/ecs"
)

var CameraComponentType = ecs.NewComponentType[*CameraComponent]()

type CameraComponent struct {
	*transform.TransformComponent

	Zoom       float64
	ZoomFactor float64

	DragStartThreshold float64
	IsDragging         bool
}

func (c *CameraComponent) Type() ecs.ComponentType {
	return CameraComponentType
}

func (c *CameraComponent) Dispose() {
	c.TransformComponent = nil
}

func NewCameraComponent() *CameraComponent {
	return &CameraComponent{
		TransformComponent: transform.NewTransformComponent(),
		Zoom:               1.0,
		ZoomFactor:         1.0,
		DragStartThreshold: 5.0,
		IsDragging:         false,
	}
}
