package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const pixelSize uint32 = 16
const height uint32 = 32
const width uint32 = 64
const windowHeight uint32 = height * pixelSize
const windowWidth uint32 = width * pixelSize
const white uint32 = 0xffffffff
const black uint32 = 0x00000000
const opBysec uint32 = 4
const fps uint32 = 60

var delay uint32 = uint32(math.Round(float64(fps / opBysec)))

type elementaryBlock struct {
	color uint32
}

type screen struct {
	pixels  [width][height]elementaryBlock
	window  *sdl.Window
	surface *sdl.Surface
}

func (s *screen) initialize() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow(
		"Emulator Chip8",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		int32(windowWidth), int32(windowHeight),
		sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	s.window = window
	s.surface = surface
	s.surface.FillRect(nil, black)
	s.window.UpdateSurface()
}

func (s *screen) update() {
	running := true

	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			case *sdl.KeyboardEvent:
				if t.Keysym.Sym == sdl.K_ESCAPE {
					running = false
					break
				}
			}
		}
		sdl.Delay(delay)
	}
	s.window.Destroy()
	sdl.Quit()
}

func (s *screen) removePixel() {
	for x := 0; uint32(x) < width; x++ {
		for y := 0; uint32(y) < height; y++ {
			s.pixels[x][y].color = black
		}
	}
}

func (s *screen) printPixel(x, y, color uint32) {
	s.surface.FillRect(nil, 0)

	rect := sdl.Rect{int32(x * pixelSize), int32(y * pixelSize), int32(pixelSize), int32(pixelSize)}
	s.surface.FillRect(&rect, color)
	s.window.UpdateSurface()
}

func main() {
	var win = screen{}

	win.removePixel()
	win.initialize()
	win.update()
}
