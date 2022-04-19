package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPrime(t *testing.T) {
	assert := assert.New(t)
	p := NewPrime(1000)
	assert.NotNil(p)
}

func TestKeyPair(t *testing.T) {
	assert := assert.New(t)
	p := NewKeyPair(1000)

	assert.NotNil(p)

	assert.Equal(p.N.Int64(), p.E.Int64()*p.D.Int64())
}
