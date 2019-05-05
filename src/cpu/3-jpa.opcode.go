package cpu

// JPA 1nnn - Jump to location nnn.
func (c *CPU) JPA(opcode uint16) {
	// recuperer juste l'address qui est defini par NNN
	c.Pc = (opcode & 0x0FFF)

	// moins 2 parce que plus 2 va etre appliquer a la fin du switch
	c.Pc -= 2
}
