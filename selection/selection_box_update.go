package selection

import (
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/geometry"
	"github.com/adm87/finch-editor/camera"
	"github.com/adm87/finch-rendering/rendering"
	"github.com/adm87/finch-rendering/vector"
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
	boxComp, _, _ := ecs.GetComponent[*vector.BoxRenderComponent](world, selectionBoxEntity, vector.BoxRenderComponentType)
	visibleComp, _, _ := ecs.GetComponent[*rendering.VisibilityComponent](world, selectionBoxEntity, rendering.VisibilityComponentType)

	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		selectionBox.SelectionStartPoint.Invalidate()
		selectionBox.SelectionEndPoint.Invalidate()
		visibleComp.IsVisible = false
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

	boxComp.Min = selectionBox.SelectionStartPoint.Value()
	boxComp.Max = selectionBox.SelectionEndPoint.Value()

	visibleComp.IsVisible = true

	return nil
}
