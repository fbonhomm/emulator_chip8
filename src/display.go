package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const pixelSize uint8 = 16
const height uint8 = 32
const width uint8 = 64
const windowHeight uint16 = uint16(height) * uint16(pixelSize)
const windowWidth uint16 = uint16(width) * uint16(pixelSize)

type elementaryBlock struct {
	opts  *ebiten.DrawImageOptions
	color color.Color
}

type screen struct {
	display [width][height]elementaryBlock
	x       int
	y       int
}

func (s *screen) update(window *ebiten.Image) error {
	// window.Fill(color.Black)
	window.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff})

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	ebitenutil.DebugPrint(screen, "Our first game in Ebiten!")
	s.x++
	s.y++
	s.drawPixel(x, y, window)
	return nil
}

func (s *screen) initialize() {
	for ; x < (x + pixelSize); x++ {
		for ; y < (y + pixelSize); y++ {
			s.display[x][y].opts.GeoM.Translate(x*pixelSize, y*pixelSize)
			s.display[x][y].color = color.Black
		}
	}

	err := ebiten.Run(s.update, windowWidth, windowHeight, 1, "Emulator Chip8")
	if err != nil {
		panic(err)
	}
}

func (s *screen) drawPixel(x, y int, window *ebiten.Image) {
	var square *ebiten.Image

	square, _ = ebiten.NewImage(pixelSize, pixelSize, ebiten.FilterNearest)

	square.fill(color.White)
	window.DrawImage(square, s.display[x][y].opts)
}

func main() {
	var win = screen{}

	win.initialize()
}
