package cpu

// RET 00EE - Return from a subroutine.
func (c *CPU) RET() {
	c.Pc = c.Stack[c.Sp]

	c.Sp--
}
