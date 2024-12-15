package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const CellSize = 30

func main() {
	rl.InitWindow(301, 601, "raylib Tetris")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	game := NewGame()
	ticker := time.NewTicker(500 * time.Millisecond)
	// go func() {
	// 	for range ticker.C {
	// 		game.MoveBlockDown()
	// 	}
	// }()
	for !rl.WindowShouldClose() {
		select {
		case _ = <-ticker.C:
			game.MoveBlockDown()
		default:
			game.HandleInput()
			rl.BeginDrawing()
			rl.ClearBackground(rl.Black)
			game.Draw()
			rl.EndDrawing()

		}
	}
}
