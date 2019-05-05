package cpu

// ORVXVY 8xy1 - Set Vx = Vx OR Vy.
func (c *CPU) ORVXVY(opcode uint16) {
	x := uint8((opcode & 0x0F00) >> 8)
	y := uint8((opcode & 0x00F0) >> 4)

	c.V[x] |= c.V[y]
}
