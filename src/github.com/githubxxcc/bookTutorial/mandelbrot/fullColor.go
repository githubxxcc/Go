package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		height, width          = 1024, 1024
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
	)

	img := image.NewRGBA((image.Rect(0, 0, height, width)))

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)

			img.Set(px, py, mandelbrot(z))
		}
	}

	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const (
		iteration = 200
		contrast  = 15
	)

	var v complex128

	for n := uint8(0); n < iteration; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{200, 255 - n*contrast, n * contrast / 255, 255}
		}
	}

	return color.RGBA{100, 100, 50, 200}
}
