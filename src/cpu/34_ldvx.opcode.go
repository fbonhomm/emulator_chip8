package cpu

// LDVX Fx65 - Read registers V0 through Vx from memory starting at location I.
func (c *CPU) LDVX(opcode uint16) {
	x := uint16((opcode & 0x0F00) >> 8)

	// different de la doc
	for i := uint16(0); i <= x; i++ {
		c.V[i] = c.Memory[c.I+i]
	}
	c.I = x + 1
}
