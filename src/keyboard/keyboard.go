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
	pressed uint8
	code    sdl.Keycode
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
					if t.Keysym.Sym == keybords[i].code {
						keybords[i].pressed = 0
					}
				}
			} else if t.Type == sdl.KEYUP {
				for i := 0; i < 15; i++ {
					if t.Keysym.Sym == keybords[i].code {
						keybords[i].pressed = 1
					}
				}
			}
			break
		}
	}
	sdl.Delay(delay)
	return true
}

// CheckKey method return 1 if key down
func CheckKey(key uint8) uint8 {
	return keybords[key].pressed
}

// WaitPressKey method return 1 if key down
func WaitPressKey(key uint8) {
	for true {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {

			case *sdl.KeyboardEvent:
				if t.Type == sdl.KEYUP {
					if t.Keysym.Sym == keybords[key].code {
						keybords[key].pressed = 1
						return
					}
				}
				break
			}
		}
	}
}
