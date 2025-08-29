package tilemaps

import (
	"github.com/adm87/finch-core/ecs"
)

var TilemapEditorSystemType = ecs.NewSystemType[*TilemapEditor]()

type TilemapEditor struct {
	enabled bool
}

func NewTilemapEditor() *TilemapEditor {
	return &TilemapEditor{
		enabled: true,
	}
}

func (t *TilemapEditor) Type() ecs.SystemType {
	return TilemapEditorSystemType
}

func (t *TilemapEditor) IsEnabled() bool {
	return t.enabled
}

func (t *TilemapEditor) Enable() {
	t.enabled = true
}

func (t *TilemapEditor) Disable() {
	t.enabled = false
}

func (t *TilemapEditor) EarlyUpdate(world *ecs.World, deltaSeconds float64) error {
	return nil
}
