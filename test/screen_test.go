package test

import (
	"emulator/src/screen"
	"testing"

	"github.com/stretchr/testify/assert"
)

var s = screen.SCREEN{}

func TestInitialize(t *testing.T) {
	s.Initialize()

	assert.Equal(t, len(s.Pixels), 64)
	assert.Equal(t, len(s.Pixels[0]), 32)
	assert.NotNil(t, s.Window)
	assert.NotNil(t, s.Surface)
}

func TestDestroy(t *testing.T) {
	s.Destroy()
	assert.Nil(t, s.Window)
	assert.Nil(t, s.Surface)
}

func TestRemoveScreen(t *testing.T) {
	s.Pixels[0][0] = 1
	s.RemoveScreen()

	assert.Equal(t, s.Pixels[0][0], uint8(0))
}
