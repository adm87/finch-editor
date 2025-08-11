package components

import (
	"github.com/adm87/finch-core/components/transform"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/geometry"
)

var CameraComponentType = ecs.NewComponentType[*CameraComponent]()

type CameraComponent struct {
	*transform.TransformComponent
	zoom float64
}

func NewCameraComponent() *CameraComponent {
	return &CameraComponent{
		TransformComponent: transform.NewTransformComponent(),
		zoom:               1.0,
	}
}

func (c *CameraComponent) Type() ecs.ComponentType {
	return CameraComponentType
}

func (c *CameraComponent) SetScale(scale geometry.Point64) {
	panic("do not use SetScale on CameraComponent. Use SetZoom instead.")
}

func (c *CameraComponent) Zoom() float64 {
	return c.zoom
}

func (c *CameraComponent) SetZoom(zoom float64) {
	if zoom <= 0 {
		panic("zoom must be greater than zero")
	}
	c.zoom = zoom
	c.TransformComponent.SetScale(geometry.Point64{X: zoom, Y: zoom})
}
