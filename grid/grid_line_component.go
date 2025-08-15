package grid

import "github.com/adm87/finch-core/ecs"

var GridLineComponentType = ecs.NewComponentType[*GridLineComponent]()

type GridLineState struct {
	Size  float32
	Scale float32
}

type GridLineComponent struct {
	CellSize    float32
	MaxGridSize float32
	MinGridSize float32
	GridStates  []GridLineState
}

func (c *GridLineComponent) Type() ecs.ComponentType {
	return GridLineComponentType
}

func (c *GridLineComponent) Dispose() {

}

func NewGridLineComponent() *GridLineComponent {
	return &GridLineComponent{
		CellSize:    32.0,
		MaxGridSize: 16.0,
		MinGridSize: 0.25,
		GridStates: []GridLineState{
			{Scale: 0.01},
			{Scale: 0.1},
			{Scale: 1.0},
			{Scale: 10.0},
			{Scale: 100.0},
		},
	}
}
