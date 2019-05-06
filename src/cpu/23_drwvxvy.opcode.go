package cpu

import (
	"emulator/src/screen"
)

// DRWVXVY Dxyn - Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision.
func (c *CPU) DRWVXVY(s *screen.SCREEN, opcode uint16) {
	x := uint16((opcode & 0x0F00) >> 8)
	y := uint16((opcode & 0x00F0) >> 4)
	n := uint16(opcode & 0x000F)
	// var pixel uint8

	c.V[0xF] = 0
	for i := uint16(0); i < n; i++ {
		pixel := c.Memory[c.I+i]

		for j := uint16(0); j < 8; j++ {
			if (pixel & (0x80 >> j)) != 0 {
				if s.Pixels[y+i][x+j] == 1 {
					s.Pixels[y+i][x+j] = 0
					c.V[0xF] = 1
				} else {
					s.Pixels[y+i][x+j] = 1
				}
			}
		}
	}
	s.ToDraw = true
}
