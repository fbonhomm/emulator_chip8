package SUBVXVY

// SUBVXVY 8xy5 - Set Vx = Vx - Vy, set VF = NOT borrow.
func SUBVXVY(c *cpu, opcode uint16) {
	var x uint8 = uint8((opcode & 0x0F00) >> 8)
	var y uint8 = uint8((opcode & 0x00F0) >> 8)

	if c.register[x] > c.register[y] {
		c.register[0xF] = 1
	} else {
		c.register[0xF] = 0
	}

	c.register[x] -= c.register[y]
}
