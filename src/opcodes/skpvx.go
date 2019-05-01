package SKPVX

// SKPVX Ex9E - Skip next instruction if key with the value of Vx is pressed.
func SKPVX(c *cpu, opcode uint16, s *screen) {
	var x uint8 = uint8((opcode & 0x0F00) >> 8)

	if s.checkKey(x) == 1 {
		c.cp++
	}
}
