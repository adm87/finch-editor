package camera

import (
	fmsg "github.com/adm87/finch-application/messages"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/geometry"
)

type CameraResizeHandler struct {
	world *ecs.World
}

func NewCameraResizeHandler(world *ecs.World) *CameraResizeHandler {
	return &CameraResizeHandler{
		world: world,
	}
}

func (handler *CameraResizeHandler) HandleMessage(msg fmsg.ApplicationResizeMessage) error {
	cameraComponent, err := FindCameraComponent(handler.world)
	if err != nil {
		return err
	}
	cameraComponent.SetOrigin(geometry.Point64{
		X: float64(msg.To.X) / 2,
		Y: float64(msg.To.Y) / 2,
	})
	return nil
}
