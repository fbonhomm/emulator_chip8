package test

import (
	"emulator/src/screen"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var s = screen.SCREEN{}

func TestInitialize(t *testing.T) {
	s.Initialize()

	assert.Equal(t, 32, len(s.Pixels))
	assert.Equal(t, 64, len(s.Pixels[0]))
	assert.NotNil(t, s.Window)
	assert.NotNil(t, s.Surface)
}

func TestDestroy(t *testing.T) {
	s.Destroy()
	assert.Nil(t, s.Window)
	assert.Nil(t, s.Surface)
}

func TestApply(t *testing.T) {
	s.Initialize()

	s.Pixels[0][0] = 1
	s.Pixels[31][63] = 1
	s.Apply()
	time.Sleep(5000 * time.Millisecond)
	s.Destroy()
}
