package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"golang.org/x/exp/rand"
)

type Game struct {
	grid      *Grid
	block     *Block
	nextBlock *Block
	blocks    []func() *Block
}

func NewGame() *Game {
	g := &Game{
		grid: NewGrid(),
		blocks: []func() *Block{
			NewIBlock,
			NewJBlock,
			NewLBlock,
			NewOBlock,
			NewSBlock,
			NewTBlock,
			NewZBlock,
		},
	}
	g.SpawnBlock()
	return g
}

func (g *Game) Draw() {
	g.grid.Draw()
	g.block.Draw()
}

func (g *Game) SpawnBlock() {
	if g.nextBlock != nil {
		g.block = g.nextBlock
	} else {
		g.block = g.blocks[rand.Intn(len(g.blocks))]()
	}
	g.nextBlock = g.blocks[rand.Intn(len(g.blocks))]()
}

func (g *Game) HandleInput() {
	key := rl.GetKeyPressed()
	if key == rl.KeyLeft {
		g.MoveBlockLeft()
	} else if key == rl.KeyRight {
		g.MoveBlockRight()
	} else if key == rl.KeyUp {
		g.RotateBlock()
	} else if key == rl.KeyDown {
		g.MoveBlockDown()
	}
}

func (g *Game) MoveBlockLeft() {
	g.block.MoveLeft()
	for _, c := range g.block.GetCellPositions() {
		if g.grid.IsCellOutside(c) || g.grid.IsCellFIlled(c) {
			g.block.MoveRight()
			break
		}
	}
}

func (g *Game) MoveBlockRight() {
	g.block.MoveRight()
	for _, c := range g.block.GetCellPositions() {
		if g.grid.IsCellOutside(c) || g.grid.IsCellFIlled(c) {
			g.block.MoveLeft()
			break
		}
	}
}

func (g *Game) MoveBlockDown() {
	g.block.MoveDown()
	for _, c := range g.block.GetCellPositions() {
		if g.grid.IsCellOutside(c) || g.grid.IsCellFIlled(c) {
			g.block.MoveUp()
			g.LockBlock()
			g.grid.ClearCompleted()
			g.SpawnBlock()
			break
		}
	}
}

func (g *Game) RotateBlock() {
	g.block.Rotate(true)
	for _, c := range g.block.GetCellPositions() {
		if g.grid.IsCellOutside(c) || g.grid.IsCellFIlled(c) {
			g.block.Rotate(false)
			break
		}
	}
}

func (g *Game) LockBlock() {
	for _, c := range g.block.GetCellPositions() {
		g.grid.grid[c[0]][c[1]] = g.block.colorId
	}
}
