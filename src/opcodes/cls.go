package CLS

import (
	screen "emulator/src"
)

// CLS remove screen
func CLS(s *screen) {
	var r = screen.SCREEN{}

	s.RemoveScreen()
}
