package systems

import (
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
	cameraComponent, err := find_camera_component(world)
	if err != nil {
		return err
	}

	matrix := cameraComponent.WorldMatrix()

	sx, sy := ebiten.CursorPosition()
	wx, wy := matrix.Apply(float64(sx), float64(sy))

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !s.downPosition.IsValid() {
			s.downPosition.SetValue(geometry.Point64{X: wx, Y: wy})
		}

		cameraPosition := cameraComponent.Position()
		downPosition := s.downPosition.Value()

		if cameraPosition.DistanceTo(downPosition) > cameraComponent.DragStartThreshold*cameraComponent.Zoom {
			cameraComponent.IsDragging = true
		}
	} else {
		if s.downPosition.IsValid() {
			s.downPosition.Invalidate()
		}
		cameraComponent.IsDragging = false
	}

	if cameraComponent.IsDragging {
		cameraPosition := cameraComponent.Position()
		downPosition := s.downPosition.Value()

		deltaX := wx - downPosition.X
		deltaY := wy - downPosition.Y

		cameraComponent.SetPosition(geometry.Point64{
			X: cameraPosition.X - deltaX,
			Y: cameraPosition.Y - deltaY,
		})
	}

	return nil
}
