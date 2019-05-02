package SNEVXB

// SNEVXB 4xkk - Skip next instruction if Vx != kk.
func SNEVXB(c *cpu, opcode uint16) {
	var x uint8 = (opcode & 0x0F00) >> 8
	var kk uint8 = (opcode & 0x00FF)

	if c.register[x] != kk {
		c.pc++
	}
}
