package cpu

// ANDVXVY 8xy2 - Set Vx = Vx AND Vy.
func (c *CPU) ANDVXVY(opcode uint16) {
	x := uint8((opcode & 0x0F00) >> 8)
	y := uint8((opcode & 0x00F0) >> 4)

	c.V[x] &= c.V[y]
}
