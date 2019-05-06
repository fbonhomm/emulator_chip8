package cpu

// CA 2nnn - Call subroutine at nnn.
func (c *CPU) CA(opcode uint16) {
	// Different de la doc
	c.Stack[c.Sp] = c.Pc

	c.Sp++

	// recuperer juste l'address qui est defini par NNN
	c.Pc = (opcode & 0x0FFF)

	// moins 2 parce que plus 2 va etre appliquer a la fin du switch
	c.Pc -= 2
}
