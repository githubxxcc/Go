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

			img.Set(px, py, newton(z))
		}
	}

	png.Encode(os.Stdout, img)
}

func newton(z complex128) color.Color {
	const (
		iteration = 200
		contrast  = 15
	)

	for n := uint8(0); n < iteration; n++ {
		z = z - (z*z*z*z-1)/(z*z*z*4)

		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*n}
		}
	}

	return color.RGBA{100, 100, 50, 200}
}
