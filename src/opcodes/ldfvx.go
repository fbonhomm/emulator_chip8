package LDFVX

import "emulator/src/screen"

// LDFVX Fx29 - Set I = location of sprite for digit Vx.
func LDFVX(c *cpu, opcode uint16, s *screen.SCREEN) {
	var x uint8 = (opcode & 0x0F00) >> 8

	c.I = c.register[x] * 5
}
