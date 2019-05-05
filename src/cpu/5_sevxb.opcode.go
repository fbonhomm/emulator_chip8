package cpu

// SEVXB 3xkk - Skip next instruction if Vx = kk.
func (c *CPU) SEVXB(opcode uint16) {
	x := uint8((opcode & 0x0F00) >> 8)
	kk := uint8((opcode & 0x00FF))

	if c.V[x] == kk {
		c.Pc += 2
	}
}
