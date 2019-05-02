package ADDIVX

// ADDIVX Fx1E - Set I = I + Vx.
func ADDIVX(c *cpu, opcode uint16) {
	var x uint8 = (opcode & 0x0F00) >> 8

	c.I += c.register[x]
}
