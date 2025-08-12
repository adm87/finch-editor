package camera

import (
	"github.com/adm87/finch-core/components/transform"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/errors"
	"github.com/adm87/finch-core/geometry"
)

var CameraComponentType = ecs.NewComponentType[*CameraComponent]()

type CameraComponent struct {
	*transform.TransformComponent
	Zoom float64
}

func (c *CameraComponent) Type() ecs.ComponentType {
	return CameraComponentType
}

func (c *CameraComponent) Dispose() {
	c.TransformComponent = nil
}

func (c *CameraComponent) SetScale(scale geometry.Point64) {
	panic(errors.NewConflictError("do not set camera scale directly. Use Zoom instead"))
}

func NewCameraComponent() *CameraComponent {
	return &CameraComponent{
		TransformComponent: transform.NewTransformComponent(),
		Zoom:               1.0,
	}
}
