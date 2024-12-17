package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	CellSize   = 90
	DrawOffset = 10
)

func main() {
	rl.InitWindow(512, 512, "raylib 2048")
	defer rl.CloseWindow()
	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	rl.SetTargetFPS(60)
	game := NewGame()
	for !rl.WindowShouldClose() {
		game.HandleInput()
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		game.Draw()
		rl.EndDrawing()

	}
}
