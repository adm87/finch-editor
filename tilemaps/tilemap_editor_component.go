package tilemaps

import (
	"image/color"

	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/geometry"
)

var TilemapEditorComponentType = ecs.NewComponentType[*TilemapEditorComponent]()

type TilemapEditorComponent struct {
	Border        geometry.Rectangle64
	BorderColor   color.RGBA
	BorderWidth   float32
	BorderVisible bool

	Cursor        geometry.Rectangle64
	CursorColor   color.RGBA
	CursorWidth   float32
	CursorVisible bool

	LoadedTilemapID string
}

func NewTilemapEditorComponent() *TilemapEditorComponent {
	return &TilemapEditorComponent{
		BorderColor:     color.RGBA{208, 135, 112, 255},
		BorderWidth:     2,
		BorderVisible:   true,
		CursorColor:     color.RGBA{235, 203, 139, 255},
		CursorWidth:     2,
		LoadedTilemapID: "",
	}
}

func (c *TilemapEditorComponent) Type() ecs.ComponentType {
	return TilemapEditorComponentType
}
