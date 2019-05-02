package SHRVXVY

// SHRVXVY 8xy6 - Set Vx = Vx SHR 1.
func SHRVXVY(c *cpu, opcode uint16) {
	var x uint8 = (opcode & 0x0F00) >> 8

	c.register[0xF] = (x & 0x01)

	// division par 2 egal a c.register[x] /= 2
	c.register[x] >>= 1
}
