package LDSTVX

// LDSTVX Fx18 - Set sound timer = Vx.
func LDSTVX(c *cpu, opcode uint16, s *screen) {
	var x uint8 = uint8((opcode & 0x0F00) >> 8)

	c.soundTimer = c.register[x]
}
