package LDIA

// LDIA Annn - Set I = nnn.
func LDIA(c *cpu, opcode uint16) {
	c.I = (opcode & 0x0FFF)
}
