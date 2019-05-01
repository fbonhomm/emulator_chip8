package RET

// RET 00EE - Return from a subroutine.
func RET(c *cpu, opcode uint16) {
	c.pc = c.stack[c.sp]

	c.sp--
	// moins 1 parce que c.pc++ va etre effectuer a la fin du switch
	c.pc--
}
