package camera

import (
	fcam "github.com/adm87/finch-core/components/camera"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/geometry"
	"github.com/hajimehoshi/ebiten/v2"
)

var CameraPanSystemType = ecs.NewSystemType[*CameraPan]()

type CameraPan struct {
}

func NewCameraPan() *CameraPan {
	return &CameraPan{}
}

func (s *CameraPan) Type() ecs.SystemType {
	return CameraPanSystemType
}

func (s *CameraPan) EarlyUpdate(world *ecs.ECSWorld, deltaSeconds float64) error {
	cameraEntity, err := FindCameraEntity(world)
	if err != nil {
		return err
	}

	cameraComponent, _, _ := ecs.GetComponent[*fcam.CameraComponent](world, cameraEntity, fcam.CameraComponentType)
	dragComponent, _, _ := ecs.GetComponent[*CameraDragComponent](world, cameraEntity, CameraDragComponentType)
	panComponent, _, _ := ecs.GetComponent[*CameraPanComponent](world, cameraEntity, CameraPanComponentType)

	if dragComponent.IsDragging {
		panComponent.IsPanning = false
		return nil
	}

	dx, dy := 0.0, 0.0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		dx -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		dx += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		dy -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		dy += 1
	}

	panComponent.IsPanning = dx != 0 || dy != 0

	delta := geometry.Point64{X: dx, Y: dy}.Normalize()
	speed := panComponent.PanSpeed * cameraComponent.Zoom

	position := cameraComponent.Position()
	position.X += delta.X * speed * deltaSeconds
	position.Y += delta.Y * speed * deltaSeconds
	cameraComponent.SetPosition(position)

	return nil
}
