package systems

import (
	"github.com/adm87/finch-application/config"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/errors"
	"github.com/adm87/finch-core/geometry"
	"github.com/adm87/finch-editor/editor/components"
)

var (
	ErrAmbiguousCameras     = errors.NewAmbiguousError("multiple cameras found")
	ErrCameraEntityNotFound = errors.NewNotFoundError("camera entity not found")
)

var CameraFinalizationType = ecs.NewSystemType[*CameraFinalization]()

type CameraFinalization struct {
	window *config.Window
}

func (s *CameraFinalization) Type() ecs.SystemType {
	return CameraFinalizationType
}

func NewCameraFinalization(window *config.Window) *CameraFinalization {
	return &CameraFinalization{
		window: window,
	}
}

func (s *CameraFinalization) LateUpdate(world *ecs.ECSWorld, deltaSeconds float64) error {
	cameraComponent, err := find_camera_component(world)
	if err != nil {
		return err
	}

	screenWidth := s.window.ScreenWidth
	screenHeight := s.window.ScreenHeight

	cameraComponent.SetOrigin(geometry.Point64{
		X: float64(screenWidth) * 0.5,
		Y: float64(screenHeight) * 0.5,
	})

	viewMatrix := cameraComponent.WorldMatrix()
	viewMatrix.Invert()

	world.SetRenderMatrix(viewMatrix)

	return nil
}

func find_camera_component(world *ecs.ECSWorld) (*components.CameraComponent, error) {
	set := world.FilterEntitiesByComponents(
		components.CameraComponentType,
	)

	count := len(set)

	if count == 0 {
		return nil, ErrCameraEntityNotFound
	}

	if count > 1 {
		return nil, ErrAmbiguousCameras
	}

	entity, ok := set.First()
	if !ok {
		return nil, ErrCameraEntityNotFound
	}

	cameraComponent, _, err := ecs.GetComponent[*components.CameraComponent](world, entity, components.CameraComponentType)
	return cameraComponent, err
}
