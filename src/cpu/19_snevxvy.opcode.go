package cpu

// SNEVXVY 9xy0 - Skip next instruction if Vx != Vy.
func (c *CPU) SNEVXVY(opcode uint16) {
	x := uint8((opcode & 0x0F00) >> 8)
	y := uint8((opcode & 0x00F0) >> 4)

	if c.V[x] != c.V[y] {
		c.Pc += 2
	}
}
