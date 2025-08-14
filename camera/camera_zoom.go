package camera

import (
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/geometry"
	"github.com/adm87/finch-core/math"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	MinZoom = 0.01
	MaxZoom = 10.0
)

var CameraZoomType = ecs.NewSystemType[*CameraZoom]()

type CameraZoom struct {
}

func NewCameraZoom() *CameraZoom {
	return &CameraZoom{}
}

func (s *CameraZoom) Type() ecs.SystemType {
	return CameraZoomType
}

func (s *CameraZoom) EarlyUpdate(world *ecs.ECSWorld, deltaSeconds float64) error {
	cameraComponent, err := FindCameraComponent(world)
	if err != nil {
		return err
	}

	dragComponent, err := FindCameraDragComponent(world)
	if err != nil {
		return err
	}

	if dragComponent.IsDragging {
		return nil
	}

	matrix := cameraComponent.WorldMatrix()
	sx, sy := ebiten.CursorPosition()
	wx, wy := matrix.Apply(float64(sx), float64(sy))

	if _, y := ebiten.Wheel(); y != 0 {
		cameraComponent.Zoom *= 1 + y*cameraComponent.ZoomFactor
	}

	cameraComponent.Zoom = math.Clamp(cameraComponent.Zoom, MinZoom, MaxZoom)
	cameraComponent.SetScale(geometry.Point64{
		X: cameraComponent.Zoom,
		Y: cameraComponent.Zoom,
	})

	matrix = cameraComponent.WorldMatrix()
	sx, sy = ebiten.CursorPosition()
	nx, ny := matrix.Apply(float64(sx), float64(sy))

	dx, dy := nx-wx, ny-wy

	position := cameraComponent.Position()
	cameraComponent.SetPosition(geometry.Point64{
		X: position.X - dx,
		Y: position.Y - dy,
	})

	return nil
}
