package lib

import (
	"math/big"
	"math/rand"
	"time"
)

type KeyPair struct {
	// number
	N *big.Int
	// encryption number
	E *big.Int
	// decryption number
	D *big.Int
}

// NewPrime with range, if n <= 1, will create a real random prime
func NewPrime(n int) *big.Int {
	// set seed for each value
	rand.Seed(time.Now().UnixNano())
	var newRandValue func() int
	if n <= 1 {
		newRandValue = func() int {
			return rand.Int()
		}
	} else {
		newRandValue = func() int {
			return rand.Intn(n)
		}
	}
	for {
		v := big.NewInt(int64(newRandValue()))
		// ref: https://en.wikipedia.org/wiki/Miller%E2%80%93Rabin_primality_test
		if v.ProbablyPrime(10) {
			return v
		}
	}
}

func NewKeyPair(n int) (p *KeyPair) {
	p = &KeyPair{}
	p.E = NewPrime(n)
	p.D = NewPrime(n)
	p.N = big.NewInt(1).Mul(p.E, p.D)
	return
}
