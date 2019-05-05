package cpu

// SUBNVXVY 8xy7 - Set Vx = Vy - Vx, set VF = NOT borrow.
func (c *CPU) SUBNVXVY(opcode uint16) {
	x := uint8((opcode & 0x0F00) >> 8)
	y := uint8((opcode & 0x00F0) >> 4)

	if c.V[y] > c.V[x] {
		c.V[0xF] = 1
	} else {
		c.V[0xF] = 0
	}

	// division par 2 egal a c.V[x] /= 2
	c.V[x] -= c.V[y]
}
