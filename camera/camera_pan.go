package camera

import (
	stdmath "math"

	fcam "github.com/adm87/finch-core/components/camera"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/geometry"
	"github.com/hajimehoshi/ebiten/v2"
)

var CameraPanSystemType = ecs.NewSystemType[*CameraPan]()

type CameraPan struct {
	enabled bool
}

func NewCameraPan() *CameraPan {
	return &CameraPan{
		enabled: true,
	}
}

func (s *CameraPan) Enable() {
	s.enabled = true
}

func (s *CameraPan) Disable() {
	s.enabled = false
}

func (s *CameraPan) IsEnabled() bool {
	return s.enabled
}

func (s *CameraPan) Type() ecs.SystemType {
	return CameraPanSystemType
}

func (s *CameraPan) EarlyUpdate(world *ecs.World, deltaSeconds float64) error {
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
	if !panComponent.IsPanning {
		return nil
	}

	_, rad := cameraComponent.Rotation()

	direction := geometry.Point64{
		X: dx*stdmath.Cos(rad) - dy*stdmath.Sin(rad),
		Y: dx*stdmath.Sin(rad) + dy*stdmath.Cos(rad),
	}.Normalize()

	speed := panComponent.PanSpeed * cameraComponent.Zoom
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		speed = panComponent.QuickPanSpeed * cameraComponent.Zoom
	}

	position := cameraComponent.Position()
	position.X += direction.X * speed * deltaSeconds
	position.Y += direction.Y * speed * deltaSeconds
	cameraComponent.SetPosition(position)

	return nil
}
