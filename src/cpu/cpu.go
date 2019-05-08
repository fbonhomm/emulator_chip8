package cpu

import (
	"emulator/src/screen"
)

const sizeMemory uint16 = 4096
const startMemory uint16 = 512
const nbrRegister uint8 = 16
const nbrLevelStack uint8 = 16
const nbrOpcode uint8 = 35

var opcodeMask = []uint16{
	0x0000, 0xFFFF, 0xFFFF, 0xF000, 0xF000, 0xF000, 0xF000, 0xF00F, 0xF000,
	0xF000, 0xF00F, 0xF00F, 0xF00F, 0xF00F, 0xF00F, 0xF00F, 0xF00F, 0xF00F,
	0xF00F, 0xF00F, 0xF000, 0xF000, 0xF000, 0xF000, 0xF0FF, 0xF0FF, 0xF0FF,
	0xF0FF, 0xF0FF, 0xF0FF, 0xF0FF, 0xF0FF, 0xF0FF, 0xF0FF, 0xF0FF,
}
var opcodeID = []uint16{
	0x0FFF, 0x00E0, 0x00EE, 0x1000, 0x2000, 0x3000, 0x4000, 0x5000, 0x6000,
	0x7000, 0x8000, 0x8001, 0x8002, 0x8003, 0x8004, 0x8005, 0x8006, 0x8007,
	0x800E, 0x9000, 0xA000, 0xB000, 0xC000, 0xD000, 0xE09E, 0xE0A1, 0xF007,
	0xF00A, 0xF015, 0xF018, 0xF01E, 0xF029, 0xF033, 0xF055, 0xF065,
}

var fonts = []uint8{
	0x0F, 0x90, 0x90, 0x90, 0xF0,
	0x20, 0x60, 0x20, 0x20, 0x70,
	0xF0, 0x10, 0xF0, 0x80, 0xF0,
	0xF0, 0x10, 0xF0, 0x10, 0xF0,
	0x90, 0x90, 0xF0, 0x10, 0x10,
	0xF0, 0x80, 0xF0, 0x10, 0xF0,
	0xF0, 0x80, 0xF0, 0x90, 0xF0,
	0xF0, 0x10, 0x20, 0x40, 0x40,
	0xF0, 0x90, 0xF0, 0x90, 0xF0,
	0xF0, 0x90, 0xF0, 0x10, 0xF0,
	0xF0, 0x90, 0xF0, 0x90, 0x90,
	0xE0, 0x90, 0xE0, 0x90, 0xE0,
	0xF0, 0x80, 0x80, 0x80, 0xF0,
	0xE0, 0x90, 0x90, 0x90, 0xE0,
	0xF0, 0x80, 0xF0, 0x80, 0xF0,
	0xF0, 0x80, 0xF0, 0x80, 0x80,
}

// CPU - cpu struct
type CPU struct {
	Stack  [nbrLevelStack]uint16
	Memory [sizeMemory]uint8
	V      [nbrRegister]uint8
	Pc     uint16
	I      uint16
	Sp     uint8
	DT     uint8
	ST     uint8
}

// Initialize - initialize cpu
func (c *CPU) Initialize(rom []uint8) {
	copy(c.Memory[0:], fonts)
	copy(c.Memory[0x200:], rom)
	c.Pc = 0x200
}

// IdentifyOpcode - return index opcode or 0 if not found
func (c *CPU) IdentifyOpcode(opcode uint16) uint8 {
	for idx := uint8(0); idx < nbrOpcode; idx++ {
		if (opcode & opcodeMask[idx]) == opcodeID[idx] {
			return idx
		}
	}
	return 0
}

// DecreaseTimer - decrease delay timer and sound timer
func (c *CPU) DecreaseTimer() {
	if c.DT > 0 {
		c.DT--
	}
	if c.ST > 0 {
		c.ST--
	}
}

// InterpreterOpcode - interprete opcode
func (c *CPU) InterpreterOpcode(display *screen.SCREEN) {
	opcode := ((uint16(c.Memory[c.Pc]) << 8) + uint16(c.Memory[c.Pc+1]))
	idx := c.IdentifyOpcode(opcode)

	switch idx {
	case 0:
		break
	case 1:
		c.CLS(display)
	case 2:
		c.RET()
	case 3:
		c.JPA(opcode)
	case 4:
		c.CA(opcode)
	case 5:
		c.SEVXB(opcode)
	case 6:
		c.SNEVXB(opcode)
	case 7:
		c.SEVXVY(opcode)
	case 8:
		c.LDVXB(opcode)
	case 9:
		c.ADDVXB(opcode)
	case 10:
		c.LDVXVY(opcode)
	case 11:
		c.ORVXVY(opcode)
	case 12:
		c.ANDVXVY(opcode)
	case 13:
		c.XORVXVY(opcode)
	case 14:
		c.ADDVXVY(opcode)
	case 15:
		c.SUBVXVY(opcode)
	case 16:
		c.SHRVX(opcode)
	case 17:
		c.SUBNVXVY(opcode)
	case 18:
		c.SHLVX(opcode)
	case 19:
		c.SNEVXVY(opcode)
	case 20:
		c.LDIA(opcode)
	case 21:
		c.JPV0(opcode)
	case 22:
		c.RNDVXB(opcode)
	case 23:
		c.DRWVXVY(display, opcode)
	case 24:
		c.SKPVX(opcode)
	case 25:
		c.SKNPVX(opcode)
	case 26:
		c.LDVXDT(opcode)
	case 27:
		c.LDVXK(opcode)
	case 28:
		c.LDDTVX(opcode)
	case 29:
		c.LDSTVX(opcode)
	case 30:
		c.ADDIVX(opcode)
	case 31:
		c.LDBVX(opcode)
	case 32:
		c.LDIVX(opcode)
	case 33:
		c.LDVX(opcode)
	}
	c.Pc += 2
}
