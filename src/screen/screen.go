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
	ToDraw  bool
	Pixels  [height][width]uint8
	Drawn   [height][width]sdl.Rect
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
	for y := uint32(0); y < height; y++ {
		for x := uint32(0); x < width; x++ {
			s.Drawn[y][x] = sdl.Rect{
				Y: int32(y * pixelSize),
				X: int32(x * pixelSize),
				H: int32(pixelSize),
				W: int32(pixelSize),
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
	if s.ToDraw == false {
		return
	}

	for y := uint32(0); y < height; y++ {
		for x := uint32(0); x < width; x++ {
			if s.Pixels[y][x] == 1 {
				s.Surface.FillRect(&s.Drawn[y][x], white)
			} else {
				s.Surface.FillRect(&s.Drawn[y][x], black)
			}
		}
	}
	s.Window.UpdateSurface()
	s.ToDraw = false
}
