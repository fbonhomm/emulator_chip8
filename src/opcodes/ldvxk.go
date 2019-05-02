package LDVXK

import "emulator/src/screen"

// LDVXK Fx0A - Wait for a key press, store the value of the key in Vx.
func LDVXK(c *cpu, opcode uint16, s *screen.SCREEN) {
	var x uint8 = (opcode & 0x0F00) >> 8

	s.waitPressKey(x)
}
