# Mandelbrot

This repo contains code to render Mandelbrot fractal using raylib library using golang.
Basicall Mandelbrot fractal is a collection of points C(a+ib) in complex space, for
which it's subsequent functional values remains finite i.e it doesn't tend to infinity.

The subsequent points is calculated by
Z(n+1) = (Z(n))^2 + C

## How to use the Application

- Install the raylib go bindings (https://github.com/gen2brain/raylib-go)
- run `go run main.go`

## Controllers

- To zoom in/out the Mandelbrot, use Mouse scroll wheel.
- To drag the fractal, use Mouse right button.
- To re-center the fractal, use left Mouse button to mark it as center.

## Note

If you are not interested in installing the go raylib bindings, I will also puslish
the executable for linux which you can try.
