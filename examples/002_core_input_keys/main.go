package main

import rl "github.com/gen2brain/raylib-go/raylib"

const screenWidth = 800
const screenHeight = 450

func main() {
	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - keyboard input")
	defer rl.CloseWindow()

	ballPosition := rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2))
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("move the ball with arrow keys", 10, 10, 20, rl.DarkGray)
		rl.DrawCircleV(ballPosition, 50, rl.Maroon)
		rl.EndDrawing()
		if rl.IsKeyDown(rl.KeyRight) {
			ballPosition.X += 2
		}
		if rl.IsKeyDown(rl.KeyLeft) {
			ballPosition.X -= 2
		}
		if rl.IsKeyDown(rl.KeyDown) {
			ballPosition.Y += 2
		}
		if rl.IsKeyDown(rl.KeyUp) {
			ballPosition.Y -= 2
		}
	}
}
