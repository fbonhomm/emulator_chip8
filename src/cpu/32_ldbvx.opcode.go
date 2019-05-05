package cpu

// LDBVX Fx33 - Store BCD representation of Vx in memory locations I, I+1, and I+2.
func (c *CPU) LDBVX(opcode uint16) {
	x := uint8((opcode & 0x0F00) >> 8)

	c.Memory[c.I] = c.V[x] / 100
	c.Memory[c.I+1] = (c.V[x] / 10) % 10
	c.Memory[c.I+2] = (c.V[x] % 100) % 10
}
