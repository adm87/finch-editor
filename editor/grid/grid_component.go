package grid

import "github.com/adm87/finch-core/ecs"

var GridComponentType = ecs.NewComponentType[*GridComponent]()

type GridState struct {
	Size  float32
	Scale float32
}

type GridComponent struct {
	CellSize    float32
	MaxGridSize float32
	MinGridSize float32
	LineColor   []float32
	GridStates  []GridState
}

func (c *GridComponent) Type() ecs.ComponentType {
	return GridComponentType
}

func (c *GridComponent) Dispose() {

}

func NewGridComponent() *GridComponent {
	return &GridComponent{
		CellSize:    32.0,
		MaxGridSize: 16.0,
		MinGridSize: 0.25,
		LineColor: []float32{
			1.0,
			1.0,
			1.0,
			0.5,
		},
		GridStates: []GridState{
			{Scale: 0.01},
			{Scale: 0.1},
			{Scale: 1.0},
			{Scale: 10.0},
			{Scale: 100.0},
		},
	}
}
