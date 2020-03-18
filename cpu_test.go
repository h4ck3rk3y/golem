package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCPU(t *testing.T) {
	c := NewCPU()
	assert.NotNil(t, c)
}
