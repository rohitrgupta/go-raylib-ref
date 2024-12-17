package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"golang.org/x/exp/rand"
)

const gridSize = 4

type Game struct {
	grid     [16]int
	colorMap map[int]rl.Color
}

func NewGame() *Game {
	colors := map[int]rl.Color{
		0:    rl.White,
		2:    rl.Beige,
		4:    rl.Pink,
		8:    rl.Yellow,
		16:   rl.Green,
		32:   rl.SkyBlue,
		64:   rl.Purple,
		128:  rl.DarkBrown,
		256:  rl.Brown,
		512:  rl.DarkGray,
		1024: rl.Blue,
		2048: rl.Lime,
		4096: rl.Red,
		8192: rl.Orange,
	}
	g := &Game{colorMap: colors}
	g.SpawnBlock()
	g.SpawnBlock()
	return g
}

func (g *Game) HandleInput() {
	if rl.IsKeyPressed(rl.KeyLeft) {
		if g.MoveBlockLeft() {
			g.SpawnBlock()
		}
	} else if rl.IsKeyPressed(rl.KeyRight) {
		if g.MoveBlockRight() {
			g.SpawnBlock()
		}
	} else if rl.IsKeyPressed(rl.KeyUp) {
		if g.MoveBlocksUp() {
			g.SpawnBlock()
		}
	} else if rl.IsKeyPressed(rl.KeyDown) {
		if g.MoveBlockDown() {
			g.SpawnBlock()
		}
	}
}
func (g *Game) SpawnBlock() {
	for {
		rand := rand.Intn(16)
		if g.grid[rand] == 0 {
			g.grid[rand] = 2
			break
		}
	}
}

func (g *Game) MoveBlocksUp() bool {
	moved := false
	for row := 0; row < gridSize; row++ {
		for col := 0; col < gridSize; col++ {
			moved = g.SlideBlockUp(row, col) || moved
		}
	}
	return moved
}
func (g *Game) SlideBlockUp(row, col int) bool {
	if row == 0 || g.grid[row*gridSize+col] == 0 {
		return false
	}
	if g.grid[(row-1)*gridSize+col] == 0 {
		// if the block above is empty, slide it up
		g.grid[(row-1)*gridSize+col] = g.grid[row*gridSize+col]
		g.grid[row*gridSize+col] = 0
		g.SlideBlockUp(row-1, col)
		return true
	} else if g.grid[(row-1)*gridSize+col] == g.grid[row*gridSize+col] {
		// if the block above is the same value, merge them
		g.grid[(row-1)*gridSize+col] *= 2
		g.grid[row*gridSize+col] = 0
		return true

	}
	return false
}

func (g *Game) MoveBlockDown() bool {
	moved := false
	for col := 0; col < gridSize; col++ {
		for row := 2; row >= 0; row-- {
			moved = g.SlideBlockDown(row, col) || moved
		}
	}
	return moved
}

func (g *Game) SlideBlockDown(row, col int) bool {
	if row == 3 || g.grid[row*gridSize+col] == 0 {
		return false
	}
	if g.grid[(row+1)*gridSize+col] == 0 {
		g.grid[(row+1)*gridSize+col] = g.grid[row*gridSize+col]
		g.grid[row*gridSize+col] = 0
		g.SlideBlockDown(row+1, col)
		return true
	} else if g.grid[(row+1)*gridSize+col] == g.grid[row*gridSize+col] {
		g.grid[(row+1)*gridSize+col] *= 2
		g.grid[row*gridSize+col] = 0
		return true
	}
	return false
}
func (g *Game) MoveBlockLeft() bool {
	moved := false
	for row := 0; row < gridSize; row++ {
		for col := 1; col < gridSize; col++ {
			if g.grid[row*gridSize+col] != 0 {
				moved = g.SlideBlockLeft(row, col) || moved
			}
		}
	}
	return moved
}

func (g *Game) SlideBlockLeft(row, col int) bool {
	if col == 0 || g.grid[row*gridSize+col] == 0 {
		return false
	}
	if g.grid[row*gridSize+col-1] == 0 {
		g.grid[row*gridSize+col-1] = g.grid[row*gridSize+col]
		g.grid[row*gridSize+col] = 0
		g.SlideBlockLeft(row, col-1)
		return true
	} else if g.grid[row*gridSize+col-1] == g.grid[row*gridSize+col] {
		g.grid[row*gridSize+col-1] *= 2
		g.grid[row*gridSize+col] = 0
		return true
	}
	return false
}

func (g *Game) MoveBlockRight() bool {
	moved := false
	for row := 0; row < gridSize; row++ {
		for col := 2; col >= 0; col-- {
			if g.grid[row*gridSize+col] != 0 {
				moved = g.SlideBlockRight(row, col) || moved
			}
		}
	}
	return moved
}
func (g *Game) SlideBlockRight(row, col int) bool {
	if col == 3 || g.grid[row*gridSize+col] == 0 {
		return false
	}
	if g.grid[row*gridSize+col+1] == 0 {
		g.grid[row*gridSize+col+1] = g.grid[row*gridSize+col]
		g.grid[row*gridSize+col] = 0
		g.SlideBlockRight(row, col+1)
		return true
	} else if g.grid[row*gridSize+col+1] == g.grid[row*gridSize+col] {
		g.grid[row*gridSize+col+1] *= 2
		g.grid[row*gridSize+col] = 0
		return true
	}
	return false
}

func (g *Game) Draw() {
	rl.DrawRectangle(0, 0, CellSize*gridSize+DrawOffset*2+5, CellSize*gridSize+DrawOffset*2+5, rl.Gray)
	for row := 0; row < gridSize; row++ {
		for col := 0; col < gridSize; col++ {
			cellColor := g.colorMap[g.grid[row*gridSize+col]]
			// cellColor := rl.LightGray
			rl.DrawRectangle(
				int32(col*CellSize+DrawOffset+5),
				int32(row*CellSize+DrawOffset+5),
				CellSize-5,
				CellSize-5,
				cellColor,
			)
			if g.grid[row*gridSize+col] != 0 {
				text := fmt.Sprintf("%d", g.grid[row*gridSize+col])
				length := rl.MeasureText(text, 30)
				rl.DrawText(
					text,
					int32(col*CellSize+DrawOffset+50-int(length/2)),
					int32(row*CellSize+DrawOffset+10+25),
					30,
					rl.Black,
				)
			}
		}
	}
}
