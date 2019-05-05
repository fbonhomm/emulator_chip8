package cpu

// ADDIVX Fx1E - Set I = I + Vx.
func (c *CPU) ADDIVX(opcode uint16) {
	x := uint8((opcode & 0x0F00) >> 8)

	c.I += uint16(c.V[x])
}
