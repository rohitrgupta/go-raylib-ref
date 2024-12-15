package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"golang.org/x/exp/rand"
)

type Game struct {
	grid      *Grid
	block     *Block
	nextBlock *Block
	blocks    []func() *Block
	GameOver  bool
	Score     int
}

func NewGame() *Game {
	g := &Game{
		grid:     NewGrid(),
		GameOver: false,
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
	rl.DrawText("Score", 350, 15, 32, rl.White)
	rl.DrawRectangleRounded(rl.NewRectangle(320, 55, 170, 60), 0.5, 5, rl.Blue)
	rl.DrawText(fmt.Sprintf("%7d", g.Score), 350, 70, 32, rl.White)
	rl.DrawText("Next", 360, 170, 32, rl.White)
	rl.DrawRectangleRounded(rl.NewRectangle(320, 215, 170, 170), 0.5, 5, rl.Blue)
	if g.GameOver {
		rl.DrawText("Game Over", 320, 450, 32, rl.White)
	}
	g.grid.Draw()
	g.block.Draw(DrawOffset, DrawOffset)
	switch g.nextBlock.colorId {
	case 3:
		g.nextBlock.Draw(255, 280)
	case 4:
		g.nextBlock.Draw(255, 265)
	default:
		g.nextBlock.Draw(270, 265)
	}
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
	if g.GameOver && key == rl.KeySpace {
		g.ResetGame()
		return
	}
	if key == rl.KeyLeft {
		g.MoveBlockLeft()
	} else if key == rl.KeyRight {
		g.MoveBlockRight()
	} else if key == rl.KeyUp {
		g.RotateBlock()
	} else if key == rl.KeyDown {
		if g.MoveBlockDown() {
			g.UpdateScore(0, 1)
		}
	}
}

func (g *Game) ResetGame() {
	g.GameOver = false
	g.Score = 0
	g.grid.Clear()
	g.block = nil
	g.nextBlock = nil
	g.SpawnBlock()
}

func (g *Game) UpdateScore(cleared, lines int) {
	g.Score += lines
	switch cleared {
	case 1:
		g.Score += 100
	case 2:
		g.Score += 300
	case 3:
		g.Score += 600
	case 4:
		g.Score += 1000
	}
}

func (g *Game) doesBlockFits() bool {
	blockFits := true
	for _, c := range g.block.GetCellPositions() {
		if g.grid.IsCellOutside(c) || g.grid.IsCellFIlled(c) {
			blockFits = false
			break
		}
	}
	return blockFits
}

func (g *Game) MoveBlockLeft() {
	g.block.MoveLeft()
	if !g.doesBlockFits() {
		g.block.MoveRight()
	}
}

func (g *Game) MoveBlockRight() {
	g.block.MoveRight()
	if !g.doesBlockFits() {
		g.block.MoveLeft()
	}
}

func (g *Game) MoveBlockDown() bool {
	if g.GameOver {
		return false
	}
	g.block.MoveDown()
	if !g.doesBlockFits() {
		g.block.MoveUp()
		g.LockBlock()
		cleared := g.grid.ClearCompleted()
		g.UpdateScore(cleared, 0)
		g.SpawnBlock()
		if !g.doesBlockFits() {
			g.GameOver = true
		}
		return false
	}
	return true
}

func (g *Game) RotateBlock() {
	g.block.Rotate(true)
	if !g.doesBlockFits() {
		g.block.Rotate(false)
	}
}

func (g *Game) LockBlock() {
	for _, c := range g.block.GetCellPositions() {
		g.grid.grid[c[0]][c[1]] = g.block.colorId
	}
}
