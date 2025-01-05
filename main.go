package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	width, height = 840, 600
	title         = "Convex hull in 2D"
)

var points, hull PointList

func main() {
	// Initialize window
	rl.InitWindow(width, height, title)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	points = make(PointList, 0)

	for !rl.WindowShouldClose() {
		// Update
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			mousePos := rl.GetMousePosition()
			// Convert to normalized coordinates (-1 to 1)
			x := (float64(mousePos.X)/float64(width)*2 - 1)
			y := -(float64(mousePos.Y)/float64(height)*2 - 1)
			points = append(points, makePoint(x, y))
			hull, _ = points.Compute()
		}

		if rl.IsKeyPressed(rl.KeyC) {
			points, hull = nil, nil
			points = make(PointList, 0)
			hull = make(PointList, 0)
		}

		// Draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)

		drawScene()

		rl.EndDrawing()
	}
}

func drawScene() {
	// Draw cartesian coordinates
	rl.DrawLineEx(
		rl.Vector2{X: 0, Y: float32(height / 2)},
		rl.Vector2{X: float32(width), Y: float32(height / 2)},
		3,
		rl.Gray,
	)
	rl.DrawLineEx(
		rl.Vector2{X: float32(width / 2), Y: 0},
		rl.Vector2{X: float32(width / 2), Y: float32(height)},
		3,
		rl.Gray,
	)

	// Your existing point drawing methods will need to be updated to use raylib drawing functions
	points.DrawPoints()
	points.DrawLowestPoint()
	hull.DrawLines()
}
