package cpu

// SHRVX 8xy6 - Set Vx = Vx SHR 1.
func (c *CPU) SHRVX(opcode uint16) {
	x := uint8((opcode & 0x0F00) >> 8)

	c.V[0xF] = (c.V[x] & 0x01)

	// division par 2 egal a c.V[x] /= 2
	c.V[x] >>= 1
}
