package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// AllColors returns a list of all the colors used in the game.
func AllColors() []rl.Color {
	return []rl.Color{
		rl.DarkGray,
		rl.Red,
		rl.Green,
		rl.Yellow,
		rl.Purple,
		rl.Orange,
		rl.Pink,
		rl.Brown,
	}
}

type Grid struct {
	grid    [20][10]int
	numRows int
	numCols int
	colors  []rl.Color
}

func NewGrid() *Grid {
	return &Grid{
		grid:    [20][10]int{},
		numRows: 20,
		numCols: 10,
		colors:  AllColors(),
	}
}
func (g *Grid) Clear() {
	for row := 0; row < g.numRows; row++ {
		for col := 0; col < g.numCols; col++ {
			g.grid[row][col] = 0
		}
	}
}

func (g *Grid) GetCellColor(row, col int) rl.Color {
	color := g.grid[row][col]
	if color < len(g.colors) {
		return g.colors[color]
	}
	return rl.White
}

func (g *Grid) Draw() {
	for row := 0; row < g.numRows; row++ {
		for col := 0; col < g.numCols; col++ {
			cellColor := g.GetCellColor(row, col)
			rl.DrawRectangle(
				int32(col*CellSize+1+DrawOffset),
				int32(row*CellSize+1+DrawOffset),
				int32(CellSize-1),
				int32(CellSize-1),
				cellColor)
		}
	}
}

// IsCellOutside returns true if the given cell is outside the grid.
func (g *Grid) IsCellOutside(cell []int) bool {
	if cell[1] < 0 || cell[1] >= g.numCols {
		return true
	}
	if cell[0] < 0 || cell[0] >= g.numRows {
		return true
	}
	return false
}

// IsCellEmpty returns true if the given cell is empty.
func (g *Grid) IsCellFIlled(cell []int) bool {
	return g.grid[cell[0]][cell[1]] != 0
}

// ClearCompleted clears all completed rows and returns the number of rows cleared.
func (g *Grid) ClearCompleted() int {
	cleared := 0
	for row := 0; row < g.numRows; row++ {
		completed := g.isCompleted(row)
		if completed {
			cleared += 1
			g.clearRow(row)
			g.dropRowsDown(row)
		}
	}
	return cleared
}

// move upper rows down (called after a row is cleared)
func (g *Grid) dropRowsDown(row int) {
	for i := row; i > 0; i-- {
		for col := 0; col < g.numCols; col++ {
			g.grid[i][col] = g.grid[i-1][col]
		}
	}
}

// isCompleted returns true if the given row is completed.
func (g *Grid) isCompleted(row int) bool {
	completed := true
	for col := 0; col < g.numCols; col++ {
		if g.grid[row][col] == 0 {
			completed = false
			break
		}
	}
	return completed
}

func (g *Grid) clearRow(row int) {
	for col := 0; col < g.numCols; col++ {
		g.grid[row][col] = 0
	}
}
