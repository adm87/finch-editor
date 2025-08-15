package camera

import (
	finmsg "github.com/adm87/finch-application/messages"
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

func (handler *CameraResizeHandler) HandleMessage(msg finmsg.ApplicationResizeMessage) error {
	cameraComponent, err := FindCameraComponent(handler.world)
	if err != nil {
		return err
	}

	cameraComponent.ViewWidth = float64(msg.To.X)
	cameraComponent.ViewHeight = float64(msg.To.Y)

	cameraComponent.SetOrigin(geometry.Point64{
		X: cameraComponent.ViewWidth / 2,
		Y: cameraComponent.ViewHeight / 2,
	})
	return nil
}
