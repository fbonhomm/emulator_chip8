package cpu

// ADDVXVY 8xy4 - Set Vx = Vx + Vy, set VF = carry.
func (c *CPU) ADDVXVY(opcode uint16) {
	x := uint8((opcode & 0x0F00) >> 8)
	y := uint8((opcode & 0x00F0) >> 4)
	sum := uint16(c.V[x]) + uint16(c.V[y])

	if sum > 0xFF {
		c.V[0xF] = 1
	} else {
		c.V[0xF] = 0
	}

	c.V[x] = uint8(sum)
}
