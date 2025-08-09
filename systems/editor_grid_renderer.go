package systems

import (
	"image/color"
	"math"

	"github.com/adm87/finch-application/config"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/hash"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	CellSize    float32 = 32.0
	MaxGridSize float32 = 16.0
	MinGridSize float32 = 0.25
)

var (
	EditorGridColor          = []float32{1.0, 1.0, 1.0, 0.5}
	EditorGridRendererType   = ecs.SystemType(hash.GetHashFromType[EditorGridRenderer]())
	EditorGridRendererFilter = []ecs.ComponentType{}
)

type ScaleGrid struct {
	Scale float32
	Size  float32
}

type EditorGridRenderer struct {
	img    *ebiten.Image
	world  *ecs.World
	window *config.Window
	grids  []ScaleGrid
}

func NewEditorGridRenderer(world *ecs.World, window *config.Window) *EditorGridRenderer {
	img := ebiten.NewImage(1, 1)
	img.Fill(color.White)
	return &EditorGridRenderer{
		img:    img,
		world:  world,
		window: window,
		grids: []ScaleGrid{
			{Scale: 0.01},
			{Scale: 0.1},
			{Scale: 1.0},
			{Scale: 10.0},
			{Scale: 100.0},
		},
	}
}

func (s *EditorGridRenderer) Filter() []ecs.ComponentType {
	return EditorGridRendererFilter
}

func (s *EditorGridRenderer) Type() ecs.SystemType {
	return EditorGridRendererType
}

func (s *EditorGridRenderer) Render(entities []*ecs.Entity, buffer *ebiten.Image, view ebiten.GeoM, interpolation float64) error {
	zoom := 1.0

	invView := view
	invView.Invert()

	paths := s.CalculateGridPaths(zoom, view, invView)
	for path, opacity := range paths {
		vertices, indices := path.AppendVerticesAndIndicesForStroke(nil, nil, &vector.StrokeOptions{
			Width: 1,
		})

		if len(vertices) == 0 || len(indices) == 0 {
			continue
		}

		for i := range vertices {
			vertices[i].ColorR = EditorGridColor[0]
			vertices[i].ColorG = EditorGridColor[1]
			vertices[i].ColorB = EditorGridColor[2]
			vertices[i].ColorA = EditorGridColor[3] * opacity
		}

		buffer.DrawTriangles(vertices, indices, s.img, &ebiten.DrawTrianglesOptions{})
	}

	return nil
}

func (s *EditorGridRenderer) CalculateGridPaths(zoom float64, view, invView ebiten.GeoM) map[*vector.Path]float32 {
	paths := make(map[*vector.Path]float32)

	for _, grid := range s.grids {
		spacing := CellSize * grid.Scale
		grid.Size = spacing / float32(zoom)

		i := grid.Size / CellSize
		rangeSize := MaxGridSize - MinGridSize
		if rangeSize == 0 {
			continue
		}

		// Calculate opacity based on how close i is to the min/max
		dMin := (i - MinGridSize) / rangeSize
		dMax := (MaxGridSize - i) / rangeSize
		opacity := dMin * dMax * 4

		if opacity <= 0.01 {
			continue
		}
		if opacity > 1 {
			opacity = 1
		} else if opacity < 0 {
			opacity = 0
		}

		path := s.GetGridPath(spacing, view, invView)
		paths[&path] = opacity
	}

	return paths
}

func (s *EditorGridRenderer) GetGridPath(spacing float32, view ebiten.GeoM, invView ebiten.GeoM) vector.Path {
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

func (s *EditorGridRenderer) GetCorners(spacing float32, view ebiten.GeoM, invView ebiten.GeoM) (float32, float32, float32, float32) {
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
