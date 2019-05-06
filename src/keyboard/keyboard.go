package keyboard

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const opBysec uint32 = 4
const fps uint32 = 60

var delay uint32 = uint32(math.Round(float64(fps / opBysec)))

// KEYBOARD struct screen
type KEYBOARD struct {
	Pressed uint8
	Code    sdl.Keycode
}

// keybords touch keyboard
var keybords = []KEYBOARD{
	KEYBOARD{0, sdl.K_0}, KEYBOARD{0, sdl.K_1}, KEYBOARD{0, sdl.K_2},
	KEYBOARD{0, sdl.K_3}, KEYBOARD{0, sdl.K_4}, KEYBOARD{0, sdl.K_5},
	KEYBOARD{0, sdl.K_6}, KEYBOARD{0, sdl.K_7}, KEYBOARD{0, sdl.K_8},
	KEYBOARD{0, sdl.K_9}, KEYBOARD{0, sdl.K_a}, KEYBOARD{0, sdl.K_b},
	KEYBOARD{0, sdl.K_c}, KEYBOARD{0, sdl.K_d}, KEYBOARD{0, sdl.K_e},
	KEYBOARD{0, sdl.K_f},
}

// GetOpBySec return operation number by second
func GetOpBySec() uint32 {
	return opBysec
}

// Event method
func Event() bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			return false
		case *sdl.KeyboardEvent:
			if t.Keysym.Sym == sdl.K_ESCAPE {
				return false
			}
			if t.Type == sdl.KEYDOWN {
				for i := 0; i < 15; i++ {
					if t.Keysym.Sym == keybords[i].Code {
						keybords[i].Pressed = 0
					}
				}
			} else if t.Type == sdl.KEYUP {
				for i := 0; i < 15; i++ {
					if t.Keysym.Sym == keybords[i].Code {
						keybords[i].Pressed = 1
					}
				}
			}
			break
		}
	}
	return true
}

// Delay method return 1 if key down
func Delay() {
	sdl.Delay(delay)
}

// CheckKey method return 1 if key down
func CheckKey(key uint8) uint8 {
	return keybords[key].Pressed
}

// GetKey method return 1 if key down
func GetKey(key uint8) *KEYBOARD {
	return &keybords[key]
}

// WaitPressKey method return 1 if key down
func WaitPressKey() uint8 {
	for true {
		for idx := 0; idx < len(keybords); idx++ {
			if keybords[idx].Pressed == 1 {
				return uint8(idx)
			}
		}
	}
	return 0
}
