package SEVXVY

// SEVXVY 5xy0 - Skip next instruction if Vx = Vy.
func SEVXVY(c *cpu, opcode uint16) {
	var x uint8 = uint8((opcode & 0x0F00) >> 8)
	var y uint8 = uint8((opcode & 0x00F0) >> 4)

	if c.register[x] == c.register[y] {
		c.pc++
	}
}
