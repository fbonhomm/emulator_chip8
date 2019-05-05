package cpu

// LDDTVX Fx15 - Set delay timer = Vx.
func (c *CPU) LDDTVX(opcode uint16) {
	x := uint8((opcode & 0x0F00) >> 8)

	c.SystemTimer = c.V[x]
}
