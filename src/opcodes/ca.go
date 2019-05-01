package CA

// CA 2nnn - Call subroutine at nnn.
func CA(c *cpu, opcode uint16) {
	c.sp++
	c.stack[c.sp] = c.pc

	// recuperer juste l'address qui est defini par NNN
	c.pc = (opcode & 0x0FFF)

	// moins 1 parce que c.pc++ va etre effectuer a la fin du switch
	c.pc--
}
