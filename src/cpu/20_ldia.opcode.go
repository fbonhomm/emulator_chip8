package cpu

// LDIA Annn - Set I = nnn.
func (c *CPU) LDIA(opcode uint16) {
	c.I = (opcode & 0x0FFF)
}
