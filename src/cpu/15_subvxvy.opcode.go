package cpu

// SUBVXVY 8xy5 - Set Vx = Vx - Vy, set VF = NOT borrow.
func (c *CPU) SUBVXVY(opcode uint16) {
	x := uint8((opcode & 0x0F00) >> 8)
	y := uint8((opcode & 0x00F0) >> 4)

	if c.V[x] > c.V[y] {
		c.V[0xF] = 1
	} else {
		c.V[0xF] = 0
	}

	c.V[x] -= c.V[y]
}
