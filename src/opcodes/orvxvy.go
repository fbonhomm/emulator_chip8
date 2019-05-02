package ORVXVY

// ORVXVY 8xy1 - Set Vx = Vx OR Vy.
func ORVXVY(c *cpu, opcode uint16) {
	var x uint8 = (opcode & 0x0F00) >> 8
	var y uint8 = (opcode & 0x00F0) >> 4

	c.register[x] |= c.register[y]
	// c.register[x] = c.register[x] | c.register[y]
}
