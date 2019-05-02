package XORVXVY

// XORVXVY 8xy3 - Set Vx = Vx XOR Vy.
func XORVXVY(c *cpu, opcode uint16) {
	var x uint8 = (opcode & 0x0F00) >> 8
	var y uint8 = (opcode & 0x00F0) >> 4

	c.register[x] ^= c.register[y]
}
