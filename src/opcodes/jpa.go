package JPA

// JPA 1nnn - Jump to location nnn.
func JPA(c *cpu, opcode uint16) {
	// recuperer juste l'address qui est defini par NNN
	c.pc = (opcode & 0x0FFF)

	// moins 1 parce que c.pc++ va etre effectuer a la fin du switch
	c.pc--
}
