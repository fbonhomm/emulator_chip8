package cpu

// RET 00EE - Return from a subroutine.
func (c *CPU) RET() {
	c.Sp--
	c.Pc = c.Stack[c.Sp]

	// different de la doc
}
