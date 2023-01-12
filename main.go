package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	width, height = 840, 600
	title         = "Convex hull in 2D"
)

var window *glfw.Window
var points, hull PointList

func main() {
	// Initialize GLFW
	glfw.Init()
	defer glfw.Terminate()

	// Create the window
	window, _ = glfw.CreateWindow(width, height, title, nil, nil)
	window.MakeContextCurrent()

	// Set the mouse button callback
	window.SetMouseButtonCallback(mouseButtonCallback)
	window.SetKeyCallback(keyCallback)

	// Initialize OpenGL
	if err := gl.Init(); err != nil {
		panic(err)
	}

	// Set the viewport size and projection matrix
	width, height := window.GetFramebufferSize()
	gl.Viewport(0, 0, int32(width), int32(height))
	aspectRatio := float32(width) / float32(height)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(float64(-aspectRatio), float64(aspectRatio), -1, 1, -1, 1)
	gl.MatrixMode(gl.MODELVIEW)

	// Set the clear color
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	points = make(PointList, 0)

	for !window.ShouldClose() {
		drawScene()
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func keyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	switch key {
	// clear the canvas
	case 'C':
		points, hull = nil, nil
		points = make(PointList, 0)
		hull = make(PointList, 0)
	}
}

func mouseButtonCallback(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if button == glfw.MouseButton1 && action == glfw.Press {
		x, y := w.GetCursorPos()
		wx, wy := w.GetSize()
		x = x/float64(wx)*2 - 1
		y = -y/float64(wy)*2 + 1
		points = append(points, makePoint(x, y))
		hull, _ = points.Compute()
	}
}

func drawScene() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	points.DrawPoints()
	points.DrawLowestPoint()
	hull.DrawLines()
	drawCartesianCoordinates()
}

func drawCartesianCoordinates() {
	gl.LineWidth(3)
	gl.Color3f(0.5, 0.5, 0.5)

	gl.Begin(gl.LINES)
	// Draw the x-axis
	gl.Vertex2f(-1, 0)
	gl.Vertex2f(1, 0)

	// Draw the y-axis
	gl.Vertex2f(0, -1)
	gl.Vertex2f(0, 1)
	gl.End()

	gl.End()
}
