package cpu

// SHLVX 8xyE - Set Vx = Vx SHL 1.
func (c *CPU) SHLVX(opcode uint16) {
	x := uint8((opcode & 0x0F00) >> 8)

	c.V[0xF] = (c.V[x] >> 7)

	// multiplication par 2 egal a c.V[x] *= 2
	c.V[x] <<= 1
}
