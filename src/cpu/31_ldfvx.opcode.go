package cpu

// LDFVX Fx29 - Set I = location of sprite for digit Vx.
func (c *CPU) LDFVX(opcode uint16) {
	x := uint8((opcode & 0x0F00) >> 8)

	c.I = uint16(c.V[x]) * 5
}
