package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const pixelSize uint8 = 16
const height uint8 = 32
const width uint8 = 64
const windowHeight uint16 = uint16(height) * uint16(pixelSize)
const windowWidth uint16 = uint16(width) * uint16(pixelSize)

type screen struct {
	display *pixel.PictureData
	window  *pixelgl.Window
}

func (s *screen) run() {
	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:     "Emulator Chip8",
		Bounds:    pixel.R(0, 0, float64(windowWidth), float64(windowHeight)),
		Resizable: false,
		VSync:     true,
	})

	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Black)

	s.window = win
	var x uint8 = 0
	var y uint8 = 0
	for !win.Closed() {
		x++
		y++
		if x >= width {
			x = 0
		}
		if y >= height {
			y = 0
		}

		fmt.Printf("x: %v, y: %v\n", x, y)

		s.drawPixel(x, y)
		win.Update()
		time.Sleep(1 * time.Second)
	}

	win.Destroy()
}

func (s *screen) initialize() {
	s.display = &pixel.PictureData{
		Pix:    make([]color.RGBA, uint32(windowWidth)*uint32(windowHeight)),
		Stride: int(windowWidth),
		Rect:   pixel.R(0, 0, float64(windowWidth), float64(windowHeight)),
	}

	pixelgl.Run(s.run)
}

func (s *screen) drawPixel(x uint8, y uint8) {
	for ; x < (x + pixelSize); x++ {
		for ; y < (y + pixelSize); y++ {
			s.display.Pix[x*y] = colornames.White
		}
	}

	spr := pixel.NewSprite(
		pixel.Picture(s.display),
		pixel.R(0, 0, float64(windowWidth), float64(windowHeight)),
	)
	spr.Draw(s.window, pixel.IM)
}

// func (s *screen) updateWindow() {
// 	for x := uint8(0); x < width; x++ {
// 		for y := uint8(0); y < height; y++ {
//
// 		}
// 	}
// }

func main() {
	var win = screen{}

	win.initialize()
}
