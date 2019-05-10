package screen

import (
	"github.com/veandco/go-sdl2/sdl"
)

const pixelSize uint32 = 16
const height uint32 = 32
const width uint32 = 64

// SCREEN - struct screen
type SCREEN struct {
	ToDraw        bool
	Pixels        [height][width]uint8
	HistoryPixels [height][width]uint8
	Drawn         [height][width]sdl.Rect
	Color         [2]*sdl.Surface
	Window        *sdl.Window
	Surface       *sdl.Surface
}

func (s *SCREEN) initializeWindow() {
	window, err := sdl.CreateWindow(
		"Emulator Chip8",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		int32(width*pixelSize), int32(height*pixelSize),
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
	s.Surface.FillRect(nil, 0x00000000)
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

func (s *SCREEN) initializeColor() {
	black, _ := sdl.CreateRGBSurface(sdl.SWSURFACE, int32(pixelSize), int32(pixelSize), 32, 0, 0, 0, 0)
	white, _ := sdl.CreateRGBSurface(sdl.SWSURFACE, int32(pixelSize), int32(pixelSize), 32, 0, 0, 0, 0)

	black.FillRect(nil, 0x00000000)
	white.FillRect(nil, 0xffffffff)

	s.Color[0] = black
	s.Color[1] = white
}

// Initialize - screen initialize
func (s *SCREEN) Initialize() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	s.initializeWindow()
	s.initializeSurface()
	s.initializeDrawn()
	s.initializeColor()
}

// Destroy - destroy screen
func (s *SCREEN) Destroy() {
	s.Surface.Free()
	s.Window.Destroy()
	s.Surface = nil
	s.Window = nil
	sdl.Quit()
}

// Apply - changement apply on screen
func (s *SCREEN) Apply() {
	if s.ToDraw == false {
		return
	}

	for y := uint32(0); y < height; y++ {
		for x := uint32(0); x < width; x++ {
			if s.Pixels[y][x] != s.HistoryPixels[y][x] {
				s.Color[s.Pixels[y][x]].Blit(nil, s.Surface, &s.Drawn[y][x])
			}
		}
	}
	s.Window.UpdateSurface()
	s.ToDraw = false
	s.HistoryPixels = s.Pixels
}
