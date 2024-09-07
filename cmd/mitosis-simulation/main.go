package main

import (
	"fmt"
	"mitosis-simulation/internals/models"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	backgroundColor               = rl.NewColor(51, 51, 51, 255)
	cells           []models.Cell = []models.Cell{}
)

func restart() {
	cells = []models.Cell{}
	cells = append(cells, *models.NewCell())
}

func main() {
	rl.InitWindow(700, 700, "mitosis-simulation")
	rl.SetTargetFPS(60)
	restart()
	camera := rl.Camera2D{
		Offset:   rl.NewVector2(float32(rl.GetScreenWidth())/2, float32(rl.GetScreenHeight())/2),
		Target:   rl.NewVector2(float32(rl.GetScreenWidth())/2, float32(rl.GetScreenHeight())/2),
		Rotation: 0,
		Zoom:     1,
	}
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.BeginMode2D(camera)
		if rl.IsKeyPressed(rl.KeyR) {
			restart()
		}
		if rl.IsKeyPressed(rl.KeyW) {
			camera.Zoom += 0.1
			// How to go to the cells direction
			camera.Target = cells[0].Vector2
		} else if rl.IsKeyPressed(rl.KeyS) {
			camera.Zoom -= 0.1
			camera.Target = cells[0].Vector2
		}
		rl.ClearBackground(backgroundColor)

		for i := len(cells) - 1; i >= 0; i-- {
			cell := &cells[i]
			cell.Move()
			cell.Draw()
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				mousePos := rl.GetScreenToWorld2D(rl.GetMousePosition(), camera)
				if cell.Clicked(mousePos) {

					cells = append(cells, cell.Mitosis()...)
					fmt.Println("Mitosis!")
					cells = append(cells[:i], cells[i+1:]...)
				}
			}
		}
		rl.EndMode2D()
		rl.EndDrawing()
	}
	rl.CloseWindow()
}

