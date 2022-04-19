package lib

import (
	"encoding/json"
	"math/big"
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
	p := NewKeyPair(100000)
	assert.NotNil(p)
	assert.Equal(p.N.Int64(), p.P.Int64()*p.Q.Int64())
	P := p.P.Int64()
	Q := p.Q.Int64()
	L := p.L.Int64()
	E := p.E.Int64()
	D := p.D.Int64()
	N := p.N.Int64()

	assert.EqualValues(P*Q, N)
	assert.EqualValues(0, L%(P-1))
	assert.EqualValues(0, L%(Q-1))

	assert.Greater(E, int64(1))
	assert.Less(E, L)

	assert.EqualValues((E*D)%L, 1)
}

func TestEncryptDecryptManual(t *testing.T) {
	assert := assert.New(t)
	p := &KeyPair{
		N: big.NewInt(323),
		E: big.NewInt(5),
		D: big.NewInt(29),
	}
	encrypted := p.Encrypt([]byte("hello rsa"))
	decrypted := p.Decrypt(encrypted)
	assert.EqualValues("hello rsa", string(decrypted))
}

func TestEncryptDecryptAuto(t *testing.T) {
	assert := assert.New(t)
	p := NewKeyPair(1000)
	encrypted := p.Encrypt([]byte("hello rsa"))
	decrypted := p.Decrypt(encrypted)
	assert.EqualValues("hello rsa", decrypted)
}

func TestEncryptDecryptMagic(t *testing.T) {
	assert := assert.New(t)
	p := NewKeyPair(1000)
	p.D, p.E = p.E, p.D // swap the encryption & decryption number
	encrypted := p.Encrypt([]byte("hello rsa"))
	decrypted := p.Decrypt(encrypted)
	assert.EqualValues("hello rsa", decrypted)
}

func TestSerializeAndDeserilization(t *testing.T) {
	assert := assert.New(t)
	p := NewKeyPair(0)
	bytes, err := json.Marshal(p)
	assert.Nil(err)
	cp := &KeyPair{}
	err = json.Unmarshal(bytes, cp)
	assert.Nil(err)
	encryped1 := p.Encrypt([]byte("whatever"))
	encryped2 := cp.Encrypt([]byte("whatever"))
	assert.EqualValues(encryped1, encryped2)
}

func TestDoSignature(t *testing.T) {
	assert := assert.New(t)
	p := NewKeyPair(0)
	msg := []byte("hello rsa")
	sig := p.Sign(msg)
	result := p.Verify(msg, sig)
	assert.True(result)
	assert.NotEqualValues(p.Sign([]byte("hello rsa")), p.Sign([]byte("hello rs2")))
}

func TestDoSignatureFailed(t *testing.T) {
	assert := assert.New(t)
	p := NewKeyPair(0)
	msg := []byte("hello rs2")
	sig := p.Sign(msg)
	result := p.Verify([]byte("hello rsa"), sig)
	assert.False(result)
}
