package opcodes

import "emulator/src/screen"

// CLS 00E0 - Clear the display.
func CLS(s *screen.SCREEN) {
	s.RemoveScreen()
}
