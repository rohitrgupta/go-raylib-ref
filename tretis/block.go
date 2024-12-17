package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Block struct {
	cells    [][][]int
	row      int
	col      int
	rotation int
	colorId  int
}

func (b *Block) MoveLeft() {
	b.col--
}
func (b *Block) MoveRight() {
	b.col++
}
func (b *Block) MoveUp() {
	b.row--
}
func (b *Block) MoveDown() {
	b.row++
}

// Rotate rotates the block clockwise or counter-clockwise.
// If clockwise is true, the block is rotated clockwise, otherwise it is rotated counter-clockwise.
func (b *Block) Rotate(clockwise bool) {
	if clockwise {
		b.rotation = (b.rotation + 1) % len(b.cells)
	} else {
		b.rotation = (b.rotation + len(b.cells) - 1) % len(b.cells)
	}
}

// GetCellPositions returns the positions of all cells in the current rotation of the block.
// The positions are relative to the block's row and column.
func (b *Block) GetCellPositions() [][]int {
	cellPositions := [][]int{}
	for _, cell := range b.cells[b.rotation] {
		cellPositions = append(cellPositions, []int{cell[0] + b.row, cell[1] + b.col})
	}
	return cellPositions
}

// Draw draws the block on the screen at the given offset.
func (b *Block) Draw(offsetX, offsetY int) {
	colors := AllColors()
	for _, cell := range b.cells[b.rotation] {
		cellRow := cell[0]
		cellCol := cell[1]
		rl.DrawRectangle(
			int32((cellCol+b.col)*CellSize+1+offsetX),
			int32((cellRow+b.row)*CellSize+1+offsetY),
			int32(CellSize-1),
			int32(CellSize-1),
			colors[b.colorId])
	}
}

// NewIBlock creates a new I-shaped block.
func NewLBlock() *Block {
	return &Block{
		col:     3,
		colorId: 1,
		cells: [][][]int{
			{{0, 2}, {1, 0}, {1, 1}, {1, 2}},
			{{0, 1}, {1, 1}, {2, 1}, {2, 2}},
			{{1, 0}, {1, 1}, {1, 2}, {2, 0}},
			{{0, 0}, {0, 1}, {1, 1}, {2, 1}},
		},
	}
}

// NewLBlock creates a new J-shaped block.
func NewJBlock() *Block {
	return &Block{
		col:     3,
		colorId: 2,
		cells: [][][]int{
			{{0, 0}, {1, 0}, {1, 1}, {1, 2}},
			{{0, 1}, {0, 2}, {1, 1}, {2, 1}},
			{{1, 0}, {1, 1}, {1, 2}, {2, 2}},
			{{0, 1}, {1, 1}, {2, 0}, {2, 1}},
		},
	}
}

// NewLBlock creates a new L-shaped block.
func NewIBlock() *Block {
	return &Block{
		col:     3,
		row:     -1,
		colorId: 3,
		cells: [][][]int{
			{{1, 0}, {1, 1}, {1, 2}, {1, 3}},
			{{0, 2}, {1, 2}, {2, 2}, {3, 2}},
			{{2, 0}, {2, 1}, {2, 2}, {2, 3}},
			{{0, 1}, {1, 1}, {2, 1}, {3, 1}},
		},
	}
}

// NewLBlock creates a new O-shaped block.
func NewOBlock() *Block {
	return &Block{
		col:     4,
		colorId: 4,
		cells: [][][]int{
			{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
			{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
			{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
			{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		},
	}
}

// NewLBlock creates a new S-shaped block.
func NewSBlock() *Block {
	return &Block{
		col:     3,
		colorId: 5,
		cells: [][][]int{
			{{0, 1}, {0, 2}, {1, 0}, {1, 1}},
			{{0, 1}, {1, 1}, {1, 2}, {2, 2}},
			{{1, 1}, {1, 2}, {2, 0}, {2, 1}},
			{{0, 0}, {1, 0}, {1, 1}, {2, 1}},
		},
	}
}

// NewLBlock creates a new T-shaped block.
func NewTBlock() *Block {
	return &Block{
		col:     3,
		colorId: 6,
		cells: [][][]int{
			{{0, 1}, {1, 0}, {1, 1}, {1, 2}},
			{{0, 1}, {1, 1}, {1, 2}, {2, 1}},
			{{1, 0}, {1, 1}, {1, 2}, {2, 1}},
			{{0, 1}, {1, 0}, {1, 1}, {2, 1}},
		},
	}
}

// NewLBlock creates a new Z-shaped block.
func NewZBlock() *Block {
	return &Block{
		col:     3,
		colorId: 7,
		cells: [][][]int{
			{{0, 0}, {0, 1}, {1, 1}, {1, 2}},
			{{0, 2}, {1, 1}, {1, 2}, {2, 1}},
			{{1, 0}, {1, 1}, {2, 1}, {2, 2}},
			{{0, 1}, {1, 0}, {1, 1}, {2, 0}},
		},
	}
}
