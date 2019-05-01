package SEVXB

// SEVXB 3xkk - Skip next instruction if Vx = kk.
func SEVXB(c *cpu, opcode uint16) {
	var x uint8 = uint8((opcode & 0x0F00) >> 8)

	if c.register[x] == uint8((opcode & 0x00FF)) {
		c.pc++
	}
}
