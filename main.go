package main

import rl "github.com/gen2brain/raylib-go/raylib"
import "image/color"

const (
	HW = 800
	Iter = 255
)

func main() {
	rl.InitWindow(HW, HW, "Mandelbrot Fractal!")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		var i int32
		var j int32
		for i = 0; i<=HW; i++ {
			for j = 0; j<=HW; j++ {
				x0 := (float32(4*i) - 2.5*float32(HW))/HW
				y0 := (float32(4*j) - 2*float32(HW))/HW
				var iter uint8 = 0
				var x float32
				var y float32
				for iter < Iter && x*x + y*y <= 4 {
					x, y = x0 + x*x - y*y, y0 + 2*x*y
					iter = iter + 1
				}
				if iter == Iter {
					// finished
					rl.DrawPixel(i, j, color.RGBA{0, 0, 0, 255})
				} else {
					rl.DrawPixel(i, j, color.RGBA{37, 7, 89, 255-iter})
				}
			}
		}
		rl.EndDrawing()
	}
}
