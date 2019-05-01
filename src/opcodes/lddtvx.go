package LDDTVX

// LDDTVX Fx15 - Set delay timer = Vx.
func LDDTVX(c *cpu, opcode uint16, s *screen) {
	var x uint8 = uint8((opcode & 0x0F00) >> 8)

	c.systemTimer = c.register[x]
}
