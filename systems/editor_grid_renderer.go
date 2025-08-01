package systems

import (
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-core/hash"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	EditorGridRendererType   = ecs.SystemType(hash.GetHashFromType[EditorGridRenderer]())
	EditorGridRendererFilter = []ecs.ComponentType{}
)

type EditorGridRenderer struct {
}

func NewEditorGridRenderer() *EditorGridRenderer {
	return &EditorGridRenderer{}
}

func (s *EditorGridRenderer) Filter() []ecs.ComponentType {
	return EditorGridRendererFilter
}

func (s *EditorGridRenderer) Type() ecs.SystemType {
	return EditorGridRendererType
}

func (s *EditorGridRenderer) Render(entities []*ecs.Entity, buffer *ebiten.Image, view ebiten.GeoM) error {
	return nil
}
