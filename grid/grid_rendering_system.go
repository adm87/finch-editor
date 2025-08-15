package grid

import (
	"image/color"
	"math"

	"github.com/adm87/finch-application/config"
	"github.com/adm87/finch-core/components/camera"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var GridRenderingSystemType = ecs.NewSystemType[*GridRenderingSystem]()

var (
	ErrAmbiguousGridRenderingEntities = errors.NewAmbiguousError("grid rendering entities")
	ErrGridRenderingEntityNotFound    = errors.NewNotFoundError("grid rendering entity")
)

type GridRenderingSystem struct {
	img    *ebiten.Image
	window *config.Window
}

func NewGridRenderingSystem(window *config.Window) *GridRenderingSystem {
	img := ebiten.NewImage(1, 1)
	img.Fill(color.RGBA{R: 216, G: 222, B: 233, A: 255})
	return &GridRenderingSystem{
		img:    img,
		window: window,
	}
}

func (s *GridRenderingSystem) Type() ecs.SystemType {
	return GridRenderingSystemType
}

func (s *GridRenderingSystem) Render(world *ecs.World, buffer *ebiten.Image) error {
	gridComponent, err := FindGridComponent(world)
	if err != nil {
		return err
	}

	cameraComponent, err := camera.FindCameraComponent(world)
	if err != nil {
		return err
	}

	view := cameraComponent.WorldMatrix()
	view.Invert()

	paths := s.CalculateGridPaths(gridComponent, cameraComponent, view)
	for path, opacity := range paths {
		vertices, indices := path.AppendVerticesAndIndicesForStroke(nil, nil, &vector.StrokeOptions{
			Width: 1,
		})

		if len(vertices) == 0 || len(indices) == 0 {
			continue
		}

		for i := range vertices {
			vertices[i].ColorR = 1
			vertices[i].ColorG = 1
			vertices[i].ColorB = 1
			vertices[i].ColorA = 0.5 * opacity
		}

		buffer.DrawTriangles(vertices, indices, s.img, &ebiten.DrawTrianglesOptions{})
	}

	return nil
}

func (s *GridRenderingSystem) CalculateGridPaths(gridComponent *GridComponent, cameraComponent *camera.CameraComponent, view ebiten.GeoM) map[*vector.Path]float32 {
	paths := make(map[*vector.Path]float32)

	for _, grid := range gridComponent.GridStates {
		spacing := gridComponent.CellSize * grid.Scale
		grid.Size = spacing / float32(cameraComponent.Zoom)

		i := grid.Size / gridComponent.CellSize
		rangeSize := gridComponent.MaxGridSize - gridComponent.MinGridSize

		// Calculate opacity based on how close i is to the min/max
		dMin := (i - gridComponent.MinGridSize) / rangeSize
		dMax := (gridComponent.MaxGridSize - i) / rangeSize
		opacity := dMin * dMax * 4

		if opacity <= 0.01 {
			continue
		}
		if opacity > 1 {
			opacity = 1
		} else if opacity < 0 {
			opacity = 0
		}

		path := s.GetGridPath(spacing, view, cameraComponent.WorldMatrix())
		paths[&path] = opacity
	}

	return paths
}

func (s *GridRenderingSystem) GetGridPath(spacing float32, view ebiten.GeoM, invView ebiten.GeoM) vector.Path {
	left, right, top, bottom := s.GetCorners(spacing, view, invView)

	path := vector.Path{}
	for x := left; x <= right; x += spacing {
		x1, y1 := view.Apply(float64(x), float64(top))
		x2, y2 := view.Apply(float64(x), float64(bottom))

		path.MoveTo(float32(x1), float32(y1))
		path.LineTo(float32(x2), float32(y2))
	}
	for y := top; y <= bottom; y += spacing {
		x1, y1 := view.Apply(float64(left), float64(y))
		x2, y2 := view.Apply(float64(right), float64(y))

		path.MoveTo(float32(x1), float32(y1))
		path.LineTo(float32(x2), float32(y2))
	}
	return path
}

func (s *GridRenderingSystem) GetCorners(spacing float32, view ebiten.GeoM, invView ebiten.GeoM) (float32, float32, float32, float32) {
	width := float32(s.window.ScreenWidth)
	height := float32(s.window.ScreenHeight)
	corners := [4][2]float64{
		{0, 0},
		{float64(width), 0},
		{0, float64(height)},
		{float64(width), float64(height)},
	}
	minX, maxX := corners[0][0], corners[0][0]
	minY, maxY := corners[0][1], corners[0][1]

	for _, c := range corners {
		x, y := invView.Apply(c[0], c[1])
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}

	startX := float32(math.Floor(minX/float64(spacing)) * float64(spacing))
	endX := float32(math.Ceil(maxX/float64(spacing)) * float64(spacing))
	startY := float32(math.Floor(minY/float64(spacing)) * float64(spacing))
	endY := float32(math.Ceil(maxY/float64(spacing)) * float64(spacing))

	return startX, endX, startY, endY
}
