package cpu

// ADDVXB 7xkk - Set Vx = Vx + kk.
func (c *CPU) ADDVXB(opcode uint16) {
	x := uint8((opcode & 0x0F00) >> 8)

	c.V[x] += uint8((opcode & 0x00FF))
}
