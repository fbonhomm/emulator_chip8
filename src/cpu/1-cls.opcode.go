package cpu

import "emulator/src/screen"

// CLS 00E0 - Clear the display.
func (c *CPU) CLS(s *screen.SCREEN) {
	s.RemoveScreen()
}
