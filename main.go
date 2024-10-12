package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Mandelbrot struct{}

func (mb *Mandelbrot) Update() error {
	return nil
}

func (mb *Mandelbrot) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	ebitenutil.DebugPrint(screen, "Hello, World")
}

func (mb *Mandelbrot) Layout(_, _ int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize()
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Mandelbrot Set")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(&Mandelbrot{}); err != nil {
		log.Fatal(err)
	}
}
