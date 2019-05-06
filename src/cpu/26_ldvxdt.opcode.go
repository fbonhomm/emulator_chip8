package cpu

// LDVXDT Fx07 - Set Vx = delay timer value.
func (c *CPU) LDVXDT(opcode uint16) {
	x := uint8((opcode & 0x0F00) >> 8)

	c.V[x] = c.DT
}
