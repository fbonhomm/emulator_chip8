package LDBVX

import "math"

// LDBVX Fx33 - Store BCD representation of Vx in memory locations I, I+1, and I+2.
func LDBVX(c *cpu, opcode uint16, s *screen) {
	var x uint8 = uint8((opcode & 0x0F00) >> 8)

	c.memory[c.I], _ = math.Modf(c.register[x] / 100)
	c.memory[c.I+1], _ = math.Modf((c.register[x] / 10) % 10)
	c.memory[c.I+2] = c.register[x] % 10
}
