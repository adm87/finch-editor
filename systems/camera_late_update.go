package systems

import (
	"github.com/adm87/finch-application/application"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/errors"
	"github.com/adm87/finch-core/geometry"
	"github.com/adm87/finch-core/hash"
	"github.com/adm87/finch-editor/components"
)

var (
	CameraLateUpdateSystemType   = ecs.NewSystemType[*CameraLateUpdateSystem]()
	CameraLateUpdateSystemFilter = []ecs.ComponentType{
		components.CameraComponentType,
	}
)

var (
	ErrAmbiguousCameras = errors.NewAmbiguousError("multiple cameras found")
	ErrCameraNotFound   = errors.NewNotFoundError("camera not found")
)

type CameraLateUpdateSystem struct {
	app *application.Application
}

func NewCameraLateUpdateSystem(app *application.Application) *CameraLateUpdateSystem {
	return &CameraLateUpdateSystem{
		app: app,
	}
}

func (s *CameraLateUpdateSystem) Type() ecs.SystemType {
	return CameraLateUpdateSystemType
}

func (s *CameraLateUpdateSystem) Filter() []ecs.ComponentType {
	return CameraLateUpdateSystemFilter
}

func (s *CameraLateUpdateSystem) LateUpdate(entities hash.HashSet[ecs.Entity], deltaSeconds float64) error {
	if len(entities) == 0 {
		return ErrCameraNotFound
	}

	if len(entities) > 1 {
		return ErrAmbiguousCameras
	}

	entity, _ := entities.First()
	camera, _, err := ecs.GetComponent[*components.CameraComponent](entity, components.CameraComponentType)
	if err != nil {
		return err
	}

	screenWidth := s.app.Config().ScreenWidth
	screenHeight := s.app.Config().ScreenHeight

	camera.SetOrigin(geometry.Point64{
		X: float64(screenWidth) * 0.5,
		Y: float64(screenHeight) * 0.5,
	})

	viewMatrix := camera.WorldMatrix()
	viewMatrix.Invert()

	s.app.SetRenderMatrix(viewMatrix)
	return nil
}
