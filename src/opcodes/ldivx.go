package LDIVX

import "emulator/src/screen"

// LDIVX Fx55 - Store registers V0 through Vx in memory starting at location I.
func LDIVX(c *cpu, opcode uint16, s *screen.SCREEN) {
	var x uint8 = (opcode & 0x0F00) >> 8

	for i := 0; i <= x; i++ {
		c.memory[cpu.I+i] = cpu.register[i]
	}
}
