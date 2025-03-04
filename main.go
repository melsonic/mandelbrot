package main

import rl "github.com/gen2brain/raylib-go/raylib"
import "image/color"
import "fmt"

const (
	HW = 800
	Iter = 255
)

func render(xc float32, yc float32) {
	rl.ClearBackground(rl.RayWhite)
	var i int32
	var j int32
	for i = 0; i<=HW; i++ {
		for j = 0; j<=HW; j++ {
			x0 := 4*(float32(i) - xc)/float32(HW)
			y0 := 4*(float32(j) - yc)/float32(HW)
			var iter uint8 = 0
			var x float32
			var y float32
			for iter < Iter && x*x + y*y <= 4 {
				x, y = x0 + x*x - y*y, y0 + 2*x*y
				iter = iter + 1
			}
			if iter == Iter {
				rl.DrawPixel(i, j, color.RGBA{0, 0, 0, iter})
			} else {
				rl.DrawPixel(i, j, color.RGBA{37, 7, 89, 255-iter})
			}
		}
	}
}

func main() {
	rl.InitWindow(HW, HW, "Mandelbrot Fractal!")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	var x float32 = HW/2
	var y float32 = HW/2
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			x = float32(rl.GetMouseX())
			y = float32(rl.GetMouseY())
		}
		render(x, y)
		rl.EndDrawing()
	}
}
