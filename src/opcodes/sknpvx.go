package SKNPVX

// SKNPVX ExA1 - Skip next instruction if key with the value of Vx is not pressed.
func SKNPVX(c *cpu, opcode uint16, s *screen) {
	var x uint8 = uint8((opcode & 0x0F00) >> 8)

	if s.checkKey(x) == 0 {
		c.cp++
	}
}
