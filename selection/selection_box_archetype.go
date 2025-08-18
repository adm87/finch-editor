package selection

import (
	"image/color"

	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/errors"
	"github.com/adm87/finch-rendering/renderers/vector"
	"github.com/adm87/finch-rendering/rendering"
)

func NewSelectionBox(world *ecs.World) (ecs.Entity, error) {
	box := vector.NewBoxRenderer()
	box.DrawBorder = true
	box.DrawFill = true
	box.SetBorder(color.RGBA{R: 235, G: 203, B: 139, A: 255})
	box.SetFill(color.RGBA{R: 235, G: 203, B: 139, A: 64})
	return world.NewEntityWithComponents(
		NewSelectionBoxComponent(),
		rendering.NewRenderComponent(box, 1000),
	)
}

func FindSelectionBoxEntity(world *ecs.World) (ecs.Entity, error) {
	entities := world.FilterEntitiesByComponents(SelectionBoxComponentType)
	if len(entities) == 0 {
		return ecs.NilEntity, errors.NewNotFoundError("selection box entity not found")
	}
	entity, ok := entities.First()
	if !ok {
		return ecs.NilEntity, errors.NewNotFoundError("selection box entity not found")
	}
	return entity, nil
}

func FindSelectionBoxComponent(world *ecs.World) (*SelectionBoxComponent, error) {
	entity, err := FindSelectionBoxEntity(world)
	if err != nil {
		return nil, err
	}
	component, _, err := ecs.GetComponent[*SelectionBoxComponent](world, entity, SelectionBoxComponentType)
	if err != nil {
		return nil, err
	}
	return component, nil
}
