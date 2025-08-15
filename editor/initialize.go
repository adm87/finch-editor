package editor

import (
	finapp "github.com/adm87/finch-application/application"
	"github.com/adm87/finch-core/components/bounds"
	"github.com/adm87/finch-core/components/transform"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/geometry"
	"github.com/adm87/finch-editor/camera"
	"github.com/adm87/finch-editor/grid"
	"github.com/adm87/finch-rendering/renderers/sprites"
	"github.com/adm87/finch-rendering/rendering"
)

func Initialize(app *finapp.Application, world *ecs.World) error {
	if _, err := camera.NewCameraEntity(world); err != nil {
		return err
	}
	if _, err := grid.NewGridLineEntity(world); err != nil {
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

	spriteA, err := world.NewEntityWithComponents(
		rendering.NewRenderComponent(spriteRenderer, 0),
		transform.NewTransformComponent(),
		bounds.NewBoundsComponent(
			spriteRenderer.Size(),
			anchor,
		),
	)
	if err != nil {
		return err
	}
	tA, _, _ := ecs.GetComponent[*transform.TransformComponent](world, spriteA, transform.TransformComponentType)
	tA.SetPosition(geometry.Point64{X: -20, Y: 0})

	spriteB, err := world.NewEntityWithComponents(
		rendering.NewRenderComponent(spriteRenderer, 0),
		transform.NewTransformComponent(),
		bounds.NewBoundsComponent(
			spriteRenderer.Size(),
			anchor,
		),
	)
	if err != nil {
		return err
	}
	tB, _, _ := ecs.GetComponent[*transform.TransformComponent](world, spriteB, transform.TransformComponentType)
	tB.SetPosition(geometry.Point64{X: 20, Y: 0})

	return nil
}
