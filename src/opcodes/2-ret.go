package opcodes

import "emulator/src/cpu"

// RET 00EE - Return from a subroutine.
func RET(c *cpu.CPU) {
	c.Pc = c.Stack[c.Sp]

	c.Sp--
}
