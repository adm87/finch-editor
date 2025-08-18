package editor

import (
	finapp "github.com/adm87/finch-application/application"
	"github.com/adm87/finch-core/components/bounds"
	"github.com/adm87/finch-core/components/transform"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/geometry"
	"github.com/adm87/finch-editor/camera"
	"github.com/adm87/finch-editor/grid"
	"github.com/adm87/finch-editor/selection"
	"github.com/adm87/finch-rendering/renderers/sprites"
	"github.com/adm87/finch-rendering/rendering"
)

func Initialize(app *finapp.Application, world *ecs.World) error {
	if _, err := grid.NewGridLines(world); err != nil {
		return err
	}
	if _, err := selection.NewSelectionBox(world); err != nil {
		return err
	}
	if _, err := camera.NewCamera(world); err != nil {
		return err
	}

	if err := app.Cache().Load("tile_0000"); err != nil {
		return err
	}

	tile0000Img, err := app.Cache().Images().Get("tile_0000")
	if err != nil {
		return err
	}

	anchor := geometry.Point64{
		X: 0.5,
		Y: 1.0,
	}

	spriteRenderer := sprites.NewSpriteRenderer(
		tile0000Img, anchor,
	)

	if _, err := NewTestEntity(world, -20, 0, spriteRenderer); err != nil {
		return err
	}
	if _, err := NewTestEntity(world, 20, 0, spriteRenderer); err != nil {
		return err
	}

	return nil
}

func NewTestEntity(world *ecs.World, x, y float64, spriteRenderer *sprites.SpriteRenderer) (ecs.Entity, error) {
	t := transform.NewTransformComponent()
	t.SetPosition(geometry.Point64{X: x, Y: y})

	return world.NewEntityWithComponents(
		rendering.NewRenderComponent(spriteRenderer, 0),
		t,
		bounds.NewBoundsComponent(
			spriteRenderer.Size(),
			geometry.Point64{X: 0.5, Y: 1.0},
		),
		selection.NewSelectableComponent(),
	)
}
