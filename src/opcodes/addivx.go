package ADDIVX

// ADDIVX Fx1E - Set I = I + Vx.
func ADDIVX(c *cpu, opcode uint16, s *screen) {
	var x uint8 = uint8((opcode & 0x0F00) >> 8)

	c.I += c.register[x]
}
