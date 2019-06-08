package cpu

import "github.com/fbonhomm/emulator_chip8/src/screen"

// CLS 00E0 - Clear the display.
func (c *CPU) CLS(s *screen.SCREEN) {
	for y := 0; y < len(s.Pixels); y++ {
		for x := 0; x < len(s.Pixels[y]); x++ {
			s.Pixels[y][x] = 0
		}
	}
	s.ToDraw = true
}
