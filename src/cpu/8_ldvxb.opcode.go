package cpu

// LDVXB 6xkk - Set Vx = kk.
func (c *CPU) LDVXB(opcode uint16) {
	x := uint8((opcode & 0x0F00) >> 8)

	c.V[x] = uint8((opcode & 0x00FF))
}
