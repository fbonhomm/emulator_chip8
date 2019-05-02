package screen

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

// SCREEN struct screen
type SCREEN struct {
	Pixels  [width][height]uint8
	Drawn   [width][height]sdl.Rect
	Window  *sdl.Window
	Surface *sdl.Surface
}

func (s *SCREEN) initializeWindow() {
	window, err := sdl.CreateWindow(
		"Emulator Chip8",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		int32(windowWidth), int32(windowHeight),
		sdl.WINDOW_SHOWN,
	)

	if err != nil {
		panic(err)
	}

	s.Window = window
}

func (s *SCREEN) initializeSurface() {
	if s.Window == nil {
		panic("You should window initialized")
	}

	surface, err := s.Window.GetSurface()

	if err != nil {
		panic(err)
	}

	s.Surface = surface
	s.Surface.FillRect(nil, black)
	s.Window.UpdateSurface()
}

func (s *SCREEN) initializeDrawn() {
	for x := uint32(0); x < width; x++ {
		for y := uint32(0); y < height; y++ {
			s.Drawn[x][y] = sdl.Rect{
				X: int32(x),
				Y: int32(y),
				W: int32(pixelSize),
				H: int32(pixelSize),
			}
		}
	}
}

// Initialize method
func (s *SCREEN) Initialize() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	s.initializeWindow()
	s.initializeSurface()
	s.initializeDrawn()
}

// Destroy window destroy
func (s *SCREEN) Destroy() {
	s.Surface.Free()
	s.Window.Destroy()
	s.Surface = nil
	s.Window = nil
	sdl.Quit()
}

// Apply method
func (s *SCREEN) Apply() {
	for x := uint32(0); x < width; x++ {
		for y := uint32(0); y < height; y++ {
			if s.Pixels[x][y] == 1 {
				s.Surface.FillRect(&s.Drawn[x][y], white)
			} else {
				s.Surface.FillRect(&s.Drawn[x][y], black)
			}
		}
	}
	s.Window.UpdateSurface()
}

// RemoveScreen method
func (s *SCREEN) RemoveScreen() {
	for x := uint32(0); x < width; x++ {
		for y := uint32(0); y < height; y++ {
			s.Pixels[x][y] = 0
		}
	}
}
