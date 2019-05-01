package screen

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

// KEYBOARD struct screen
type KEYBOARD struct {
	press uint8
	code uint8
}

var keybords = []KEYBOARD{
	KEYBOARD{0, sdl.K_0}, KEYBOARD{0, sdl.K_1}, KEYBOARD{0, sdl.K_2},
	KEYBOARD{0, sdl.K_3}, KEYBOARD{0, sdl.K_4}, KEYBOARD{0, sdl.K_5},
	KEYBOARD{0, sdl.K_6}, KEYBOARD{0, sdl.K_7}, KEYBOARD{0, sdl.K_8},
	KEYBOARD{0, sdl.K_9}, KEYBOARD{0, sdl.K_a}, KEYBOARD{0, sdl.K_b},
	KEYBOARD{0, sdl.K_c}, KEYBOARD{0, sdl.K_d}, KEYBOARD{0, sdl.K_e},
	KEYBOARD{0, sdl.K_f}
}

// SCREEN struct screen
type SCREEN struct {
	pixels  [width][height]uint8
	window  *sdl.Window
	surface *sdl.Surface
}

// Initialize method
func (s *SCREEN) Initialize() {
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

// Event method
func (s *SCREEN) Event() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			return false
		case *sdl.KEYDOWN:
			for i := 0; i < 15; i++ {
				if t.Keysym.Sym == keybords[i].code {
					keybords[i].pressed = 0
				}
			}
			break
		case *sdl.KEYUP:
			for i := 0; i < 15; i++ {
				if t.Keysym.Sym == keybords[i].code {
					keybords[i].pressed = 1
				}
			}
			if t.Keysym.Sym == sdl.K_ESCAPE {
				return false
			}
			break
		}
	}
	sdl.Delay(delay)
	return true
}

// Destroy window destroy
func (s *screen) Destroy() {
	s.window.Destroy()
	sdl.Quit()
}

// RemoveScreen method
func (s *SCREEN) RemoveScreen() {
	for x := uint32(0); x < width; x++ {
		for y := uint32(0); y < height; y++ {
			s.pixels[x][y] = 0
		}
	}
	s.surface.FillRect(nil, black)
	s.window.UpdateSurface()
}

// checkKey method return 1 if key down
func (s *SCREEN) checkKey(key) uint8 {
	return keybords[key].pressed
}

// waitPressKey method return 1 if key down
func (s *SCREEN) waitPressKey(key) {
	for true {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.KEYUP:
				if t.Keysym.Sym == keybords[key].code {
					keybords[i].pressed = 1
					return
				}
			}
		}
	}
}
