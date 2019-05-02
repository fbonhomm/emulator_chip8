package SUBNVXVY

// SUBNVXVY 8xy7 - Set Vx = Vy - Vx, set VF = NOT borrow.
func SUBNVXVY(c *cpu, opcode uint16) {
	var x uint8 = (opcode & 0x0F00) >> 8
	var y uint8 = (opcode & 0x00F0) >> 4

	if c.register[y] > c.register[x] {
		c.register[0xF] = 1
	} else {
		c.register[0xF] = 0
	}

	// division par 2 egal a c.register[x] /= 2
	c.register[x] -= c.register[y]
}
