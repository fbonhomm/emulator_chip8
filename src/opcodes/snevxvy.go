package SNEVXVY

// SNEVXVY 9xy0 - Skip next instruction if Vx != Vy.
func SNEVXVY(c *cpu, opcode uint16) {
	var x uint8 = (opcode & 0x0F00) >> 8
	var y uint8 = (opcode & 0x00F0) >> 4

	if c.register[x] != c.register[y] {
		c.pc++
	}
}
