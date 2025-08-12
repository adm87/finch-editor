package camera

import (
	"github.com/adm87/finch-application/config"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/errors"
	"github.com/adm87/finch-core/geometry"
)

var CameraFinalizationUpdateType = ecs.NewSystemType[*CameraFinalizationUpdate]()

var (
	ErrAmbiguousCameras     = errors.NewAmbiguousError("multiple cameras found")
	ErrCameraEntityNotFound = errors.NewNotFoundError("camera entity not found")
)

type CameraFinalizationUpdate struct {
	window *config.Window
}

func (s *CameraFinalizationUpdate) Type() ecs.SystemType {
	return CameraFinalizationUpdateType
}

func NewCameraFinalizationUpdate(window *config.Window) *CameraFinalizationUpdate {
	return &CameraFinalizationUpdate{
		window: window,
	}
}

func (s *CameraFinalizationUpdate) LateUpdate(world *ecs.ECSWorld, deltaSeconds float64) error {
	entity, err := internal_get_camera_entity(world)
	if err != nil {
		return err
	}

	cameraComponent, _, _ := ecs.GetComponent[*CameraComponent](world, entity, CameraComponentType)

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

func internal_get_camera_entity(world *ecs.ECSWorld) (ecs.Entity, error) {
	set := world.FilterEntitiesByComponents(
		CameraComponentType,
	)

	count := len(set)

	if count == 0 {
		return ecs.NilEntity, ErrCameraEntityNotFound
	}

	if count > 1 {
		return ecs.NilEntity, ErrAmbiguousCameras
	}

	if entity, ok := set.First(); ok {
		return entity, nil
	}

	return ecs.NilEntity, nil
}
