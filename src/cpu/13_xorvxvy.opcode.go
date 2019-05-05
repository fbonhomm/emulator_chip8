package cpu

// XORVXVY 8xy3 - Set Vx = Vx XOR Vy.
func (c *CPU) XORVXVY(opcode uint16) {
	x := uint8((opcode & 0x0F00) >> 8)
	y := uint8((opcode & 0x00F0) >> 4)

	c.V[x] ^= c.V[y]
}
