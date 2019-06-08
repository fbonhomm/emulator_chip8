package cpu

import (
	"github.com/fbonhomm/emulator_chip8/src/screen"
)

// DRWVXVY Dxyn - Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision.
func (c *CPU) DRWVXVY(s *screen.SCREEN, opcode uint16) {
	x := uint8((opcode & 0x0F00) >> 8)
	y := uint8((opcode & 0x00F0) >> 4)
	n := uint8(opcode & 0x000F)
	var pixel uint8
	var vy, vx uint8

	c.V[0xF] = 0
	for i := uint8(0); i < n; i++ {
		pixel = c.Memory[c.I+uint16(i)]

		for j := uint8(0); j < 8; j++ {
			if (pixel & (0x80 >> j)) != 0 {
				if (c.V[y] + i) >= 32 {
					vy = (c.V[y] + i) - 32
				} else {
					vy = c.V[y] + i
				}

				if (c.V[x] + j) >= 64 {
					vx = (c.V[x] + j) - 64
				} else {
					vx = c.V[x] + j
				}

				if s.Pixels[vy][vx] == 1 {
					c.V[0xF] = 1
				}
				s.Pixels[vy][vx] ^= 0x01
			}
		}
	}
	s.ToDraw = true
}
