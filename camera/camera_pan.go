package camera

import (
	fcam "github.com/adm87/finch-core/components/camera"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/geometry"
	"github.com/adm87/finch-core/types"
	"github.com/hajimehoshi/ebiten/v2"
)

var CameraPanType = ecs.NewSystemType[*CameraPan]()

type CameraPan struct {
	downPosition types.Optional[geometry.Point64]
}

func NewCameraPan() *CameraPan {
	return &CameraPan{
		downPosition: types.NewEmptyOption[geometry.Point64](),
	}
}

func (s *CameraPan) Type() ecs.SystemType {
	return CameraPanType
}

func (s *CameraPan) EarlyUpdate(world *ecs.ECSWorld, deltaSeconds float64) error {
	cameraEntity, err := FindCameraEntity(world)
	if err != nil {
		return err
	}

	cameraComponent, _, _ := ecs.GetComponent[*fcam.CameraComponent](world, cameraEntity, fcam.CameraComponentType)
	dragComponent, _, _ := ecs.GetComponent[*CameraDragComponent](world, cameraEntity, CameraDragComponentType)

	matrix := cameraComponent.WorldMatrix()

	sx, sy := ebiten.CursorPosition()
	wx, wy := matrix.Apply(float64(sx), float64(sy))

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !s.downPosition.IsValid() {
			s.downPosition.SetValue(geometry.Point64{X: wx, Y: wy})
		}

		cameraPosition := cameraComponent.Position()
		downPosition := s.downPosition.Value()

		if cameraPosition.DistanceTo(downPosition) > dragComponent.DragStartThreshold*cameraComponent.Zoom {
			dragComponent.IsDragging = true
		}
	} else {
		if s.downPosition.IsValid() {
			s.downPosition.Invalidate()
		}
		dragComponent.IsDragging = false
	}

	if dragComponent.IsDragging {
		cameraPosition := cameraComponent.Position()
		downPosition := s.downPosition.Value()

		deltaX := wx - downPosition.X
		deltaY := wy - downPosition.Y

		cameraComponent.SetPosition(geometry.Point64{
			X: cameraPosition.X - deltaX,
			Y: cameraPosition.Y - deltaY,
		})
		dragComponent.DragVector.SetValue(geometry.Point64{
			X: deltaX,
			Y: deltaY,
		})
	}

	return nil
}
