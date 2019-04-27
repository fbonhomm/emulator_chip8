package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const pixelSize uint32 = 16
const height uint32 = 32
const width uint32 = 64
const windowHeight uint32 = height * pixelSize
const windowWidth uint32 = width * pixelSize
const white uint32 = 0xffffffff
const black uint32 = 0x00000000

type elementaryBlock struct {
	color uint32
}

type screen struct {
	display [width][height]elementaryBlock
	window  *sdl.Window
	surface *sdl.Surface
}

func (s *screen) initialize() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(
		"Emulator Chip8",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		int32(windowWidth), int32(windowHeight),
		sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	s.window = window
	s.surface = surface
	s.surface.FillRect(nil, black)
	s.window.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}
	}
}

func (s *screen) initDisplay() {
	for x := 0; uint32(x) < width; x++ {
		for y := 0; uint32(y) < height; y++ {
			s.display[x][y].color = black
		}
	}
}

func (s *screen) printPixel(x, y uint32) {
	s.surface.FillRect(nil, 0)

	rect := sdl.Rect{int32(x), int32(y), int32(pixelSize), int32(pixelSize)}
	s.surface.FillRect(&rect, 0xffff0000)
	s.window.UpdateSurface()
}

func main() {
	var win = screen{}

	// win.initDisplay()
	win.initialize()
}
