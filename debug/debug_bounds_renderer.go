package debug

import (
	"image/color"

	"github.com/adm87/finch-core/components/bounds"
	"github.com/adm87/finch-core/components/transform"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-editor/camera"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var DebugBoundsRendererType = ecs.NewSystemType[*DebugBoundsRenderer]()

type DebugBoundsRenderer struct {
	img *ebiten.Image
}

func NewDebugBoundsRenderer() *DebugBoundsRenderer {
	img := ebiten.NewImage(1, 1)
	img.Fill(color.RGBA{R: 163, G: 190, B: 140, A: 255})
	return &DebugBoundsRenderer{img: img}
}

func (d *DebugBoundsRenderer) Type() ecs.SystemType {
	return DebugBoundsRendererType
}

func (d *DebugBoundsRenderer) Render(world *ecs.World, buffer *ebiten.Image) error {
	entities := world.FilterEntitiesByComponents(
		bounds.BoundsComponentType,
		transform.TransformComponentType,
	)

	if len(entities) == 0 {
		return nil
	}

	cameraComponent, err := camera.FindCameraComponent(world)
	if err != nil {
		return err
	}

	view := cameraComponent.WorldMatrix()
	view.Invert()

	path := vector.Path{}

	for entity := range entities {
		boundsComponent, _, _ := ecs.GetComponent[*bounds.BoundsComponent](world, entity, bounds.BoundsComponentType)
		transformComponent, _, _ := ecs.GetComponent[*transform.TransformComponent](world, entity, transform.TransformComponentType)

		rect := boundsComponent.AABB(transformComponent.Position())

		x, y := rect.X, rect.Y
		w, h := rect.Width, rect.Height

		x0, y0 := view.Apply(x, y)
		x1, y1 := view.Apply(x+w, y)
		x2, y2 := view.Apply(x+w, y+h)
		x3, y3 := view.Apply(x, y+h)

		pts := [4][2]float64{
			{x0, y0},
			{x1, y1},
			{x2, y2},
			{x3, y3},
		}

		path.MoveTo(float32(pts[0][0]), float32(pts[0][1]))
		for i := 1; i < 4; i++ {
			path.LineTo(float32(pts[i][0]), float32(pts[i][1]))
		}
		path.LineTo(float32(pts[0][0]), float32(pts[0][1]))
	}

	vertices, indices := path.AppendVerticesAndIndicesForStroke(nil, nil, &vector.StrokeOptions{
		Width: 2,
	})

	buffer.DrawTriangles(vertices, indices, d.img, &ebiten.DrawTrianglesOptions{})

	return nil

}
