package cpu

import (
	"emulator/src/opcodes"
	"emulator/src/screen"
)

// 0x000-0x1FF - Chip 8 interpreter (contains font set in emu)
// 0x050-0x0A0 - Used for the built in 4x5 pixel font set (0-F)
// 0x200-0xFFF - Program ROM and work RAM

const sizeMemory uint16 = 4096 // taille de la memoire total
const startMemory uint16 = 512 // adresse du debut de la memoire
const nbrRegister uint8 = 16   // nombre total de registres
const nbrLevelStack uint8 = 16 // nombre de niveaux pour la stack
const nbrOpcode uint8 = 35     // nombre de niveaux pour la stack

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

// CPU primary struct
type CPU struct {
	Memory      [sizeMemory]uint8     // tableaux qui represente la memoire
	V           [nbrRegister]uint8    // tableaux qui represente les registres V{0..F}
	Stack       [nbrLevelStack]uint16 // tableaux qui represente la stack de sauvegarde
	Pc          uint16                // represente la tete de lecture de la memoire PC programCounter
	I           uint16                // stocke une adresse memoire
	Sp          uint8                 // represente la tete de lecture de la stack stackPointer
	SystemTimer uint8                 // mimuterie systeme
	SoundTimer  uint8                 // mimuterie sonore
}

// IdentifyOpcode return index opcode or -1 if not found
func (c *CPU) IdentifyOpcode(opcode uint16) uint8 {
	var idx uint8

	for idx = 0; idx < nbrOpcode; idx++ {
		if (opcode & opcodeMask[idx]) == opcodeID[idx] {
			break
		}
	}

	return idx
}

// Decrease decremente les timers
func (c *CPU) Decrease() {
	if c.SystemTimer > 0 {
		c.SystemTimer--
	}
	if c.SoundTimer > 0 {
		c.SoundTimer--
	}
}

// InterpreterOpcode return index opcode or -1 if not found
func (c *CPU) InterpreterOpcode(display *screen.SCREEN, opcode uint16) {
	idx := c.IdentifyOpcode(opcode)

	switch idx {
	case 0:
		break
	case 1:
		opcodes.CLS(display)
		// case 2:
		// 	opcodes.RET(&c)
		// case 3:
		// 	opcodes.JPA(&c, opcode)
		// 	case 4:
		// 	opcodes.CA(c, opcode)
		// case 5:
		// 	opcodes.SEVXB(c, opcode)
		// case 6:
		// 	opcodes.SNEVXB(c, opcode)
		// case 7:
		// 	opcodes.SEVXVY(c, opcode)
		// case 8:
		// 	opcodes.LDVXB(c, opcode)
		// case 9:
		// 	opcodes.ADDVXB(c, opcode)
		// case 10:
		// 	opcodes.LDVXVY(c, opcode)
		// case 11:
		// 	opcodes.ORVXVY(c, opcode)
		// case 12:
		// 	opcodes.ANDVXVY(c, opcode)
		// case 13:
		// 	opcodes.XORVXVY(c, opcode)
		// case 14:
		// 	opcodes.ADDVXVY(c, opcode)
		// case 15:
		// 	opcodes.SUBVXVY(c, opcode)
		// case 16:
		// 	opcodes.SHRVX(c, opcode)
		// case 17:
		// 	opcodes.SUBNVXVY(c, opcode)
		// case 18:
		// 	opcodes.SHLVX(c, opcode)
		// case 19:
		// 	opcodes.SNEVXVY(c, opcode)
		// case 21:
		// 	opcodes.LDIA(c, opcode)
		// case 22:
		// 	opcodes.JPV0(c, opcode)
		// case 23:
		// 	opcodes.RNDVXB(c, opcode)
		// case 24:
		// 	opcodes.DRWVXVY(c, display, opcode)
		// case 25:
		// 	opcodes.SKPVX(c, opcode)
		// case 26:
		// 	opcodes.SKNPVX(c, opcode)
		// case 27:
		// 	opcodes.LDVXDT(c, opcode)
		// case 28:
		// 	opcodes.LDVXK(opcode)
		// case 29:
		// 	opcodes.LDDTVX(c, opcode)
		// case 30:
		// 	opcodes.LDSTVX(c, opcode)
		// case 31:
		// 	opcodes.ADDIVX(c, opcode)
		// case 32:
		// 	opcodes.LDBVX(c, opcode)
		// case 33:
		// 	opcodes.LDIVX(c, opcode)
		// case 34:
		// 	opcodes.LDVX(c, opcode)
	}
	c.Pc += 2
}
