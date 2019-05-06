package cpu

// ADDIVX Fx1E - Set I = I + Vx.
func (c *CPU) ADDIVX(opcode uint16) {
	x := uint8((opcode & 0x0F00) >> 8)

	if (c.I + uint16(c.V[x])) > 0xFFF {
		c.V[0xF] = 1
	} else {
		c.V[0xF] = 0
	}

	c.I += uint16(c.V[x])

	// diffrent de la doc
}
