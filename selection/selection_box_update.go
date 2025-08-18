package selection

import (
	"github.com/adm87/finch-core/components/camera"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/geometry"
	"github.com/adm87/finch-rendering/renderers/vector"
	"github.com/adm87/finch-rendering/rendering"
	"github.com/hajimehoshi/ebiten/v2"
)

var SelectionBoxUpdateType = ecs.NewSystemType[*SelectionBoxUpdate]()

type SelectionBoxUpdate struct {
	enabled bool
}

func NewSelectionBoxUpdate() *SelectionBoxUpdate {
	return &SelectionBoxUpdate{
		enabled: true,
	}
}

func (s *SelectionBoxUpdate) Type() ecs.SystemType {
	return SelectionBoxUpdateType
}

func (s *SelectionBoxUpdate) Enable() {
	s.enabled = true
}

func (s *SelectionBoxUpdate) Disable() {
	s.enabled = false
}

func (s *SelectionBoxUpdate) IsEnabled() bool {
	return s.enabled
}

func (s *SelectionBoxUpdate) EarlyUpdate(world *ecs.World, deltaSeconds float64) error {
	selectionBoxEntity, err := FindSelectionBoxEntity(world)
	if err != nil {
		return err
	}

	selectionBox, _, _ := ecs.GetComponent[*SelectionBoxComponent](world, selectionBoxEntity, SelectionBoxComponentType)
	renderComponent, _, _ := ecs.GetComponent[*rendering.RenderComponent](world, selectionBoxEntity, rendering.RenderComponentType)

	boxRenderer, ok := renderComponent.Renderer.(*vector.BoxRenderer)

	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		selectionBox.SelectionStartPoint.Invalidate()
		selectionBox.SelectionEndPoint.Invalidate()
		renderComponent.IsVisible = false
		return nil
	}

	cameraComponent, err := camera.FindCameraComponent(world)
	if err != nil {
		return err
	}
	if cameraComponent == nil {
		return nil
	}

	view := cameraComponent.WorldMatrix()
	mx, my := ebiten.CursorPosition()
	wx, wy := view.Apply(float64(mx), float64(my))

	if !selectionBox.SelectionStartPoint.IsValid() {
		selectionBox.SelectionStartPoint.SetValue(geometry.Point64{X: wx, Y: wy})
	}
	selectionBox.SelectionEndPoint.SetValue(geometry.Point64{X: wx, Y: wy})

	if !ok {
		return nil
	}

	start := selectionBox.SelectionStartPoint.Value()
	end := selectionBox.SelectionEndPoint.Value()

	boxRenderer.SetArea(start, end)
	renderComponent.IsVisible = true

	return nil
}
