package main

import (
	"image/color"
	"log"
	"math/cmplx"

	"github.com/hajimehoshi/ebiten/v2"
)

const ITERATIONS uint8 = 100

type Mandelbrot struct{}

func IsInSet(c complex128) uint8 {
	z := 0 + 0i
	for i := uint8(0); i < ITERATIONS; i++ {
		z = cmplx.Pow(z, 2) + c
		if cmplx.Abs(z)*cmplx.Abs(z) > 10 {
			return i
		}
	}
	return 0
}

// Linear interpolation function
func Lerp(a, b, f float64) float64 {
	return (a*(1.0-f) + (b * f))
}

func (mb *Mandelbrot) Update() error {
	return nil
}

func (mb *Mandelbrot) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{255, 255, 255, 255})

	for x := 0.0; x < 1.0; x += 0.001 {
		for y := 0.0; y < 1.0; y += 0.001 {
			px := Lerp(-2.0, 2.0, x)
			py := Lerp(-2.0, 2.0, y)
			iter := IsInSet(complex(px, py))
			if iter == 0 {
				screen.Set(
					int(x*1000),
					int(y*1000),
					color.RGBA{0, 0, 0, 255},
				)
			} else {
				screen.Set(
					int(x*100),
					int(y*100),
					color.RGBA{255 - iter, 255 - iter, 255 - iter, 255},
				)
			}
		}
	}
}

func (mb *Mandelbrot) Layout(_, _ int) (screenWidth, screenHeight int) {
	return 100, 100
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Mandelbrot Set")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(&Mandelbrot{}); err != nil {
		log.Fatal(err)
	}
}
