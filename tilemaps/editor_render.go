package tilemaps

import (
	"image/color"

	"github.com/adm87/finch-core/components/camera"
	"github.com/adm87/finch-core/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var TilemapEditorRenderSystemType = ecs.NewSystemType[*TilemapEditorRender]()

type TilemapEditorRender struct {
	enabled bool

	drawImg     *ebiten.Image
	drawOptions *ebiten.DrawTrianglesOptions
}

func NewTilemapEditorRender() *TilemapEditorRender {
	img := ebiten.NewImage(1, 1)
	img.Fill(color.White)
	return &TilemapEditorRender{
		enabled:     true,
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
	tilemapComp, err := FindTilemapComponent(world)
	if err != nil {
		return err
	}

	if tilemapComp.TilemapID == "" {
		return nil // No loaded tilemap, nothing to render.
	}

	editorComp, err := FindTilemapEditorComponent(world)
	if err != nil {
		return err
	}

	cameraComp, err := camera.FindCameraComponent(world)
	if err != nil {
		return err
	}

	minX, minY := camera.WorldToScreen(cameraComp, editorComp.Border.X, editorComp.Border.Y)
	maxX, maxY := camera.WorldToScreen(cameraComp, editorComp.Border.X+editorComp.Border.Width, editorComp.Border.Y+editorComp.Border.Height)
	t.draw_tilemap_rect(minX, minY, maxX, maxY, editorComp.BorderWidth, editorComp.BorderColor, buffer)

	if editorComp.CursorVisible {
		minX, minY = camera.WorldToScreen(cameraComp, editorComp.Cursor.X, editorComp.Cursor.Y)
		maxX, maxY = camera.WorldToScreen(cameraComp, editorComp.Cursor.X+editorComp.Cursor.Width, editorComp.Cursor.Y+editorComp.Cursor.Height)
		t.draw_tilemap_rect(minX, minY, maxX, maxY, editorComp.CursorWidth, editorComp.CursorColor, buffer)
	}

	return nil
}

func (t *TilemapEditorRender) draw_tilemap_rect(minX, minY, maxX, maxY float64, w float32, c color.RGBA, buffer *ebiten.Image) {
	path := vector.Path{}
	path.MoveTo(float32(minX), float32(minY))
	path.LineTo(float32(maxX), float32(minY))
	path.LineTo(float32(maxX), float32(maxY))
	path.LineTo(float32(minX), float32(maxY))
	path.Close()
	vs, is := path.AppendVerticesAndIndicesForStroke(nil, nil, &vector.StrokeOptions{
		Width: w,
	})
	if len(vs) == 0 || len(is) == 0 {
		return
	}
	for i := 0; i < len(vs); i++ {
		vs[i].ColorR = float32(c.R) / 255.0
		vs[i].ColorG = float32(c.G) / 255.0
		vs[i].ColorB = float32(c.B) / 255.0
		vs[i].ColorA = float32(c.A) / 255.0
	}
	buffer.DrawTriangles(vs, is, t.drawImg, t.drawOptions)
}
