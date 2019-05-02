package LDSTVX

import "emulator/src/screen"

// LDSTVX Fx18 - Set sound timer = Vx.
func LDSTVX(c *cpu, opcode uint16, s *screen.SCREEN) {
	var x uint8 = (opcode & 0x0F00) >> 8

	c.soundTimer = c.register[x]
}
