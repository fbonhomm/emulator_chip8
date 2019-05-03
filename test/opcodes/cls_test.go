package test

import (
	"emulator/src/opcodes"
	"emulator/src/screen"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCLS(t *testing.T) {
	var s = screen.SCREEN{}

	s.Pixels[1][1] = 1
	opcodes.CLS(&s)

	assert.Equal(t, uint8(0), s.Pixels[1][1])
}
