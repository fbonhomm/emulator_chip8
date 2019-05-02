package ADDVXB

// ADDVXB 7xkk - Set Vx = Vx + kk.
func ADDVXB(c *cpu, opcode uint16) {
	var x uint8 = (opcode & 0x0F00) >> 8

	c.register[x] += uint8((opcode & 0x00FF))
}
