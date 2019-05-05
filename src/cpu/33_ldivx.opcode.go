package cpu

// LDIVX Fx55 - Store registers V0 through Vx in memory starting at location I.
func (c *CPU) LDIVX(opcode uint16) {
	x := uint16((opcode & 0x0F00) >> 8)

	for i := uint16(0); i <= x; i++ {
		c.Memory[c.I+i] = c.V[i]
	}
}
