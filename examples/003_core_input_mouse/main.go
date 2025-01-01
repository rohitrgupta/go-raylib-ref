package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth  = 800
	screenHeight = 450
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - mouse input")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	ballPosition := rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2))
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("move the ball with mouse and click mouse button", 10, 10, 20, rl.DarkGray)
		rl.DrawCircleV(ballPosition, 50, rl.Maroon)
		rl.EndDrawing()
		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			ballPosition = rl.GetMousePosition()
		}
	}
}
