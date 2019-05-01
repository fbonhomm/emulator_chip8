package SHLVXVY

// SHLVXVY 8xyE - Set Vx = Vx SHL 1.
func SHLVXVY(c *cpu, opcode uint16) {
	var x uint8 = uint8((opcode & 0x0F00) >> 8)

	c.register[0xF] = (x >> 7)

	// multiplication par 2 egal a c.register[x] *= 2
	c.register[x] <<= 1
}
