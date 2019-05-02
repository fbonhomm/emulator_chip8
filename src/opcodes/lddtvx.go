package LDDTVX

import "emulator/src/screen"

// LDDTVX Fx15 - Set delay timer = Vx.
func LDDTVX(c *cpu, opcode uint16, s *screen.SCREEN) {
	var x uint8 = (opcode & 0x0F00) >> 8

	c.systemTimer = c.register[x]
}
