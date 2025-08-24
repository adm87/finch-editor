package editor

import (
	"math/rand"

	finapp "github.com/adm87/finch-application/application"
	"github.com/adm87/finch-core/components/bounds"
	"github.com/adm87/finch-core/components/transform"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/geometry"
	"github.com/adm87/finch-editor/camera"
	"github.com/adm87/finch-editor/grid"
	"github.com/adm87/finch-editor/selection"
	"github.com/adm87/finch-rendering/rendering"
	"github.com/adm87/finch-rendering/sprites"
	"github.com/adm87/finch-resources/storage"
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

	if err := storage.Load("tile_0000"); err != nil {
		return err
	}

	min := -300
	max := 300

	for i := 0; i < 100; i++ {
		x := rand.Intn(max-min) + min
		y := rand.Intn(max-min) + min
		if _, err := NewTestEntity(world, float64(x), float64(y), i); err != nil {
			return err
		}
	}

	return nil
}

func NewTestEntity(world *ecs.World, x, y float64, z int) (ecs.Entity, error) {
	t := transform.NewTransformComponent()
	t.SetPosition(geometry.Point64{X: x, Y: y})

	sprite := sprites.NewSpriteRenderComponent("tile_0000", z)
	sprite.Anchor = geometry.Point64{X: 0.5, Y: 0.5}

	img, _ := storage.ImageHandle(sprite.ImageID).Get()
	origin := geometry.Point64{
		X: float64(img.Bounds().Dx()) * sprite.Anchor.X,
		Y: float64(img.Bounds().Dy()) * sprite.Anchor.Y,
	}

	visibility := rendering.NewVisibilityComponent()
	visibility.VisibleArea.SetValue(geometry.Rectangle64{
		X:      -origin.X,
		Y:      -origin.Y,
		Width:  float64(img.Bounds().Dx()),
		Height: float64(img.Bounds().Dy()),
	})

	bounds := bounds.NewBoundsComponent(
		geometry.Point64{
			X: float64(img.Bounds().Dx()),
			Y: float64(img.Bounds().Dy()),
		},
		sprite.Anchor,
	)

	return world.NewEntityWithComponents(t, sprite, visibility, bounds)
}
