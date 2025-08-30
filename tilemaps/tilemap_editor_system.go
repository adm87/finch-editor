package tilemaps

import (
	stdmath "math"

	"github.com/adm87/finch-core/components/camera"
	"github.com/adm87/finch-core/components/transform"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/geometry"
	"github.com/adm87/finch-core/math"
	"github.com/hajimehoshi/ebiten/v2"

	finch "github.com/adm87/finch-application/application"
	tm "github.com/adm87/finch-tilemap/tilemaps"
	ts "github.com/adm87/finch-tilemap/tilesets"
)

var TilemapEditorSystemType = ecs.NewSystemType[*TilemapEditorSystem]()

type TilemapEditorSystem struct {
	app *finch.Application

	enabled bool
}

func NewTilemapEditorSystem(app *finch.Application) *TilemapEditorSystem {
	return &TilemapEditorSystem{
		app:     app,
		enabled: true,
	}
}

func (t *TilemapEditorSystem) Type() ecs.SystemType {
	return TilemapEditorSystemType
}

func (t *TilemapEditorSystem) IsEnabled() bool {
	return t.enabled
}

func (t *TilemapEditorSystem) Enable() {
	t.enabled = true
}

func (t *TilemapEditorSystem) Disable() {
	t.enabled = false
}

func (t *TilemapEditorSystem) EarlyUpdate(world *ecs.World, deltaSeconds float64) error {
	editorComp, err := FindTilemapEditorComponent(world)
	if err != nil {
		return err
	}

	tilemapComp, err := FindTilemapComponent(world)
	if err != nil {
		return err
	}

	transformComp, err := FindTilemapEditorTransform(world)
	if err != nil {
		return err
	}

	if editorComp.LoadedTilemap == "" {
		return nil
	}

	updateBorder := tilemapComp.TilemapID != editorComp.LoadedTilemap
	if updateBorder {
		if err := t.load_tilemap(tilemapComp, editorComp); err != nil {
			return err
		}
	}

	tilemap, tileset, err := get_tilemap_info(tilemapComp.TilemapID)
	if err != nil {
		return err
	}

	if updateBorder {
		if err := t.update_editor_border(editorComp, transformComp, tilemap, tileset); err != nil {
			return err
		}
	}

	if err := t.update_editor_cursor(world, editorComp, transformComp, tilemap, tileset); err != nil {
		return err
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		tx := int((editorComp.Cursor.X - editorComp.Border.X) / float64(tileset.TileSize))
		ty := int((editorComp.Cursor.Y - editorComp.Border.Y) / float64(tileset.TileSize))

		tx = math.Clamp(tx, 0, tilemap.Columns-1)
		ty = math.Clamp(ty, 0, tilemap.Rows-1)

		tile := 0
		if ebiten.IsKeyPressed(ebiten.KeyControl) {
			tile = -1
		}

		if err := t.place_tile(tx, ty, tile, tilemap); err != nil {
			return err
		}
	}

	return nil
}

func (t *TilemapEditorSystem) load_tilemap(tilemapComp *tm.TilemapComponent, editorComp *TilemapEditorComponent) error {
	tilemapComp.TilemapID = editorComp.LoadedTilemap
	// TODO - make sure all dependencies are loaded for the tilemap

	t.app.SetTitleContext(tilemapComp.TilemapID)
	return nil
}

func (t *TilemapEditorSystem) update_editor_border(editorComp *TilemapEditorComponent, transformComp *transform.TransformComponent, tilemap *tm.Tilemap, tileset *ts.Tileset) error {
	borderWidth := float64(tilemap.Columns * tileset.TileSize)
	borderHeight := float64(tilemap.Rows * tileset.TileSize)
	borderX := -borderWidth / 2
	borderY := -borderHeight / 2

	transformComp.SetPosition(geometry.Point64{X: borderX, Y: borderY})

	editorComp.Border.X = borderX
	editorComp.Border.Y = borderY
	editorComp.Border.Width = borderWidth
	editorComp.Border.Height = borderHeight
	return nil
}

func (t *TilemapEditorSystem) update_editor_cursor(world *ecs.World, editorComp *TilemapEditorComponent, transformComp *transform.TransformComponent, tilemap *tm.Tilemap, tileset *ts.Tileset) error {
	cameraComp, err := camera.FindCameraComponent(world)
	if err != nil {
		return err
	}

	size := float64(tileset.TileSize)
	editorComp.Cursor.Width = size
	editorComp.Cursor.Height = size

	sx, sy := ebiten.CursorPosition()
	wx, wy := camera.ScreenToWorld(cameraComp, float64(sx), float64(sy))

	wx -= size * 0.5
	wy -= size * 0.5

	wx = stdmath.Round(wx/size) * size
	wy = stdmath.Round(wy/size) * size

	editorComp.Cursor.X = wx
	editorComp.Cursor.Y = wy

	editorComp.CursorVisible = editorComp.Border.Intersects(editorComp.Cursor)
	return nil
}

func (t *TilemapEditorSystem) place_tile(x, y int, tile int, tilemap *tm.Tilemap) error {
	i := y*tilemap.Columns + x

	if tilemap.Data[i] == tile {
		return nil
	}

	tilemap.Data[i] = tile
	tilemap.IsDirty = true
	return nil
}
