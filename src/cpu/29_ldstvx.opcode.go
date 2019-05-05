package cpu

// LDSTVX Fx18 - Set sound timer = Vx.
func (c *CPU) LDSTVX(opcode uint16) {
	x := uint8((opcode & 0x0F00) >> 8)

	c.SoundTimer = c.V[x]
}
