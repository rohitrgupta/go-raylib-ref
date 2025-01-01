package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	windowWidth  = 640
	windowHeight = 480
	scrollSpeed  = 10
)

func main() {
	rl.InitWindow(windowWidth, windowHeight, "raylib [core] example - input mouse wheel")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	boxPositionY := int32(windowHeight / 2)

	for !rl.WindowShouldClose() {
		boxPositionY -= int32(rl.GetMouseWheelMove() * scrollSpeed)
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawRectangle(int32(windowWidth/2-40), boxPositionY, 80, 80, rl.Maroon)

		rl.DrawText("Use mouse wheel to move the cube up and down!", 10, 10, 20, rl.Gray)
		rl.DrawText(fmt.Sprintf("Box position Y: %03d", boxPositionY), 10, 40, 20, rl.LightGray)

		rl.EndDrawing()
	}
}
