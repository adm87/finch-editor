package camera

import (
	fcam "github.com/adm87/finch-core/components/camera"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/geometry"
	"github.com/hajimehoshi/ebiten/v2"
)

var CameraDragType = ecs.NewSystemType[*CameraDrag]()

type CameraDrag struct {
}

func NewCameraDrag() *CameraDrag {
	return &CameraDrag{}
}

func (s *CameraDrag) Type() ecs.SystemType {
	return CameraDragType
}

func (s *CameraDrag) EarlyUpdate(world *ecs.ECSWorld, deltaSeconds float64) error {
	cameraEntity, err := FindCameraEntity(world)
	if err != nil {
		return err
	}

	cameraComponent, _, _ := ecs.GetComponent[*fcam.CameraComponent](world, cameraEntity, fcam.CameraComponentType)
	dragComponent, _, _ := ecs.GetComponent[*CameraDragComponent](world, cameraEntity, CameraDragComponentType)
	panComponent, _, _ := ecs.GetComponent[*CameraPanComponent](world, cameraEntity, CameraPanComponentType)

	if panComponent.IsPanning {
		dragComponent.IsDragging = false
		return nil
	}

	matrix := cameraComponent.WorldMatrix()

	sx, sy := ebiten.CursorPosition()
	wx, wy := matrix.Apply(float64(sx), float64(sy))

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !dragComponent.DownPosition.IsValid() {
			dragComponent.DownPosition.SetValue(geometry.Point64{X: wx, Y: wy})
		}

		cameraPosition := cameraComponent.Position()
		downPosition := dragComponent.DownPosition.Value()

		if cameraPosition.DistanceTo(downPosition) > dragComponent.DragStartThreshold*cameraComponent.Zoom {
			dragComponent.IsDragging = true
		}
	} else {
		if dragComponent.DownPosition.IsValid() {
			dragComponent.DownPosition.Invalidate()
		}
		dragComponent.IsDragging = false
	}

	if dragComponent.IsDragging {
		cameraPosition := cameraComponent.Position()
		downPosition := dragComponent.DownPosition.Value()

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
