package editor

import (
	finapp "github.com/adm87/finch-application/application"
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
	if _, err := grid.NewEditorGridEntity(world); err != nil {
		return err
	}

	tile0000Img, err := app.Cache().Images().Get("tile0000")
	if err != nil {
		return err
	}

	spriteRenderer := sprites.NewSpriteRenderer(
		tile0000Img, geometry.Point64{
			X: 0.5,
			Y: 0.5,
		},
	)

	if _, err := world.NewEntityWithComponents(
		rendering.NewRenderComponent(spriteRenderer, 0),
		transform.NewTransformComponent(),
	); err != nil {
		return err
	}

	return nil
}
