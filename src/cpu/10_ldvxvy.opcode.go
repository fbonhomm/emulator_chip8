package cpu

// LDVXVY 8xy0 - Set Vx = Vy.
func (c *CPU) LDVXVY(opcode uint16) {
	x := uint8((opcode & 0x0F00) >> 8)
	y := uint8((opcode & 0x00F0) >> 4)

	c.V[x] = c.V[y]
}
