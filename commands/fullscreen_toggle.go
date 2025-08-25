package commands

import (
	"github.com/adm87/finch-application/config"
	"github.com/hajimehoshi/ebiten/v2"
)

type FullscreenToggle struct {
	window *config.Window
}

func NewFullscreenToggle(window *config.Window) *FullscreenToggle {
	return &FullscreenToggle{
		window: window,
	}
}

func (t *FullscreenToggle) Execute() error {
	t.window.Fullscreen = !t.window.Fullscreen
	ebiten.SetFullscreen(t.window.Fullscreen)
	return nil
}
