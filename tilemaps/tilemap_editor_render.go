package tilemaps

import (
	"image/color"

	"github.com/adm87/finch-core/components/transform"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-editor/camera"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	tm "github.com/adm87/finch-tilemap/tilemaps"
	ts "github.com/adm87/finch-tilemap/tilesets"
)

var TilemapEditorRenderSystemType = ecs.NewSystemType[*TilemapEditorRender]()

type TilemapEditorRender struct {
	enabled bool

	borderColor color.RGBA
	borderWidth float32

	drawImg     *ebiten.Image
	drawOptions *ebiten.DrawTrianglesOptions
}

func NewTilemapEditorRender() *TilemapEditorRender {
	img := ebiten.NewImage(1, 1)
	img.Fill(color.White)
	return &TilemapEditorRender{
		enabled:     true,
		borderColor: color.RGBA{136, 192, 208, 255},
		borderWidth: 2,
		drawImg:     img,
		drawOptions: &ebiten.DrawTrianglesOptions{},
	}
}

func (t *TilemapEditorRender) IsEnabled() bool {
	return t.enabled
}

func (t *TilemapEditorRender) Enable() {
	t.enabled = true
}

func (t *TilemapEditorRender) Disable() {
	t.enabled = false
}

func (t *TilemapEditorRender) Type() ecs.SystemType {
	return TilemapEditorRenderSystemType
}

func (t *TilemapEditorRender) Render(world *ecs.World, buffer *ebiten.Image) error {
	entity, ok := world.FilterEntitiesByComponents(tm.TilemapComponentType).First()
	if !ok {
		return nil
	}

	tilemapComp, _, _ := ecs.GetComponent[*tm.TilemapComponent](world, entity, tm.TilemapComponentType)
	transformComp, _, _ := ecs.GetComponent[*transform.TransformComponent](world, entity, transform.TransformComponentType)

	cameraComp, err := camera.FindCameraComponent(world)
	if err != nil {
		return err
	}

	tilemap, err := tm.Cache().Get(tilemapComp.TilemapID)
	if err != nil {
		return err
	}

	tileset, err := ts.Cache().Get(tilemap.TilesetID)
	if err != nil {
		return err
	}

	position := transformComp.Position()

	view := cameraComp.WorldMatrix()
	view.Invert()

	matrix := transformComp.WorldMatrix()
	matrix.Concat(view)

	x, y := matrix.Apply(position.X, position.Y)
	width := float64(tilemap.Columns*tileset.TileSize) / cameraComp.Zoom
	height := float64(tilemap.Rows*tileset.TileSize) / cameraComp.Zoom

	path := vector.Path{}
	path.MoveTo(float32(x), float32(y))
	path.LineTo(float32(x+width), float32(y))
	path.LineTo(float32(x+width), float32(y+height))
	path.LineTo(float32(x), float32(y+height))
	path.Close()

	vs, is := path.AppendVerticesAndIndicesForStroke(nil, nil, &vector.StrokeOptions{
		Width: t.borderWidth,
	})

	if len(vs) == 0 || len(is) == 0 {
		return nil
	}

	for i := 0; i < len(vs); i++ {
		vs[i].ColorR = float32(t.borderColor.R) / 255.0
		vs[i].ColorG = float32(t.borderColor.G) / 255.0
		vs[i].ColorB = float32(t.borderColor.B) / 255.0
		vs[i].ColorA = float32(t.borderColor.A) / 255.0
	}

	buffer.DrawTriangles(vs, is, t.drawImg, t.drawOptions)
	return nil
}
