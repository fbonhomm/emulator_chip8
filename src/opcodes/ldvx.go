package LDVX

// LDVX Fx65 - Read registers V0 through Vx from memory starting at location I.
func LDVX(c *cpu, opcode uint16, s *screen) {
	var x uint8 = uint8((opcode & 0x0F00) >> 8)

	for i := 0; i <= x; i++ {
		cpu.register[i] = c.memory[cpu.I+i]
	}
}
