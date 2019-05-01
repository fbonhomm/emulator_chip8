package ANDVXVY

// ANDVXVY 8xy2 - Set Vx = Vx AND Vy.
func ANDVXVY(c *cpu, opcode uint16) {
	var x uint8 = uint8((opcode & 0x0F00) >> 8)
	var y uint8 = uint8((opcode & 0x00F0) >> 8)

	c.register[x] &= c.register[y]
	// c.register[x] = c.register[x] & c.register[y]
}
