package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGCD(t *testing.T) {
	assert.Equal(t, 6, GCD(48, 18))
	assert.Equal(t, 6, GCD(18, 48))
}

func TestLCM(t *testing.T) {
	assert.Equal(t, 14, LCM(2, 7))
	assert.Equal(t, 14, LCM(7, 2))
	assert.Equal(t, 12, LCM(4, 6))
	assert.Equal(t, 12, LCM(6, 4))
}

func TestLCMM(t *testing.T) {
	assert.Equal(t, 420, LCMM([]int{2, 3, 4, 7, 5}))
}
