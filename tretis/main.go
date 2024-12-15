package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	CellSize   = 30
	DrawOffset = 10
)

func main() {
	rl.InitWindow(501, 621, "raylib Tetris")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	game := NewGame()
	ticker := time.NewTicker(500 * time.Millisecond)
	for !rl.WindowShouldClose() {
		select {
		case <-ticker.C:
			game.MoveBlockDown()
		default:
			game.HandleInput()
			rl.BeginDrawing()
			rl.ClearBackground(rl.DarkBlue)
			game.Draw()
			rl.EndDrawing()

		}
	}
}
