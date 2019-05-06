package cpu

import "emulator/src/keyboard"

// LDVXK Fx0A - Wait for a key press, store the value of the key in Vx.
func (c *CPU) LDVXK(opcode uint16) {
	x := uint8((opcode & 0x0F00) >> 8)

	c.V[x] = keyboard.WaitPressKey()
}
