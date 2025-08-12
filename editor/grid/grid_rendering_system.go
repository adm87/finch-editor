package grid

import (
	"image/color"
	"math"

	"github.com/adm87/finch-application/config"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/errors"
	"github.com/adm87/finch-editor/editor/camera"
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
	img.Fill(color.White)
	return &GridRenderingSystem{
		img:    img,
		window: window,
	}
}

func (s *GridRenderingSystem) Type() ecs.SystemType {
	return GridRenderingSystemType
}

func (s *GridRenderingSystem) Render(world *ecs.ECSWorld, buffer *ebiten.Image, view ebiten.GeoM) error {
	entity, err := internal_get_rendering_entity(world)
	if err != nil {
		return err
	}

	gridComponent, _, _ := ecs.GetComponent[*GridComponent](world, entity, GridComponentType)
	cameraComponent, _, _ := ecs.GetComponent[*camera.CameraComponent](world, entity, camera.CameraComponentType)

	paths := s.CalculateGridPaths(gridComponent, cameraComponent, view)
	for path, opacity := range paths {
		vertices, indices := path.AppendVerticesAndIndicesForStroke(nil, nil, &vector.StrokeOptions{
			Width: 1,
		})

		if len(vertices) == 0 || len(indices) == 0 {
			continue
		}

		for i := range vertices {
			vertices[i].ColorR = gridComponent.LineColor[0]
			vertices[i].ColorG = gridComponent.LineColor[1]
			vertices[i].ColorB = gridComponent.LineColor[2]
			vertices[i].ColorA = gridComponent.LineColor[3] * opacity
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
		if rangeSize == 0 {
			continue
		}

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

	topLeftX, topLeftY := invView.Apply(0, 0)
	topRightX, topRightY := invView.Apply(float64(width), 0)
	bottomLeftX, bottomLeftY := invView.Apply(0, float64(height))
	bottomRightX, bottomRightY := invView.Apply(float64(width), float64(height))

	left := float32(math.Min(math.Min(topLeftX, topRightX), math.Min(bottomLeftX, bottomRightX)))
	right := float32(math.Max(math.Max(topLeftX, topRightX), math.Max(bottomLeftX, bottomRightX)))
	top := float32(math.Min(math.Min(topLeftY, topRightY), math.Min(bottomLeftY, bottomRightY)))
	bottom := float32(math.Max(math.Max(topLeftY, topRightY), math.Max(bottomLeftY, bottomRightY)))

	startX := float32(math.Floor(float64(left/spacing)) * float64(spacing))
	endX := float32(math.Ceil(float64(right/spacing)) * float64(spacing))
	startY := float32(math.Floor(float64(top/spacing)) * float64(spacing))
	endY := float32(math.Ceil(float64(bottom/spacing)) * float64(spacing))

	return startX, endX, startY, endY
}

func internal_get_rendering_entity(world *ecs.ECSWorld) (ecs.Entity, error) {
	set := world.FilterEntitiesByComponents(
		camera.CameraComponentType,
		GridComponentType,
	)

	count := len(set)

	if count == 0 {
		return ecs.NilEntity, ErrGridRenderingEntityNotFound
	}

	if count > 1 {
		return ecs.NilEntity, ErrAmbiguousGridRenderingEntities
	}

	if entity, ok := set.First(); ok {
		return entity, nil
	}

	return ecs.NilEntity, nil
}
