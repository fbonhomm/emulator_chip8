package cpu

// JPV0 Bnnn - Jump to location nnn + V0.
func (c *CPU) JPV0(opcode uint16) {
	c.Pc = (opcode & 0x0FFF) + uint16(c.V[0x0])

	// moins 2 parce que plus 2 va etre appliquer a la fin du switch
	c.Pc -= 2
}
