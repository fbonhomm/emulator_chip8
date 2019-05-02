package LDVXB

// LDVXB 6xkk - Set Vx = kk.
func LDVXB(c *cpu, opcode uint16) {
	var x uint8 = (opcode & 0x0F00) >> 8

	c.register[x] == uint8((opcode & 0x00FF))
}
