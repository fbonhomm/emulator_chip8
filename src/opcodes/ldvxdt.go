package LDVXDT

// LDVXDT Fx07 - Set Vx = delay timer value.
func LDVXDT(c *cpu, opcode uint16) {
	var x uint8 = (opcode & 0x0F00) >> 8

	c.register[x] = c.systemTimer
}
