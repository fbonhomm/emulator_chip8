package JPV0

// JPV0 Bnnn - Jump to location nnn + V0.
func JPV0(c *cpu, opcode uint16) {
	c.cp = (opcode & 0x0FFF) + c.register[0x00]
}
