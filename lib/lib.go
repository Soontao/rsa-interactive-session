package lib

import (
	"crypto/sha256"
	"math/big"
	"math/rand"
	"time"
)

type KeyPair struct {
	// prime number 1
	P *big.Int
	// prime number 2
	Q *big.Int
	// lcm(p-1, q-1) = (p-1) * (q-1)/ gcd(p-1, q-1)
	L *big.Int
	// 1 < E < L, gcd(E, L) = 1
	E *big.Int
	// 1 < D < L, E * D mod L  = 1
	D *big.Int
	// p*q number
	N *big.Int
}

// Encrypt message, by public key, encrypt message by the Encryption number
func (k *KeyPair) Encrypt(msg []byte) (enctypedParts []*big.Int) {
	for _, msgPart := range msg {
		enctypedParts = append(enctypedParts, big.NewInt(1).Exp(big.NewInt(int64(msgPart)), k.E, k.N))
	}
	return
}

// Verify signature, by public key, decrypt signature by the Encryption number
func (k *KeyPair) Verify(msg []byte, signature []*big.Int) bool {

	// compute hash local
	expectedHash := BytesToBigInt(k.HashMsg(msg))

	if len(expectedHash) != len(signature) {
		return false
	}

	decrypedHash := []*big.Int{}

	for idx, sigPart := range signature {
		decrtypedSigPart := big.NewInt(1).Exp(sigPart, k.E, k.N)
		decrypedHash = append(decrypedHash, decrtypedSigPart)
		if decrtypedSigPart.Cmp(expectedHash[idx]) != 0 {
			return false
		}
	}

	return true
}

// Decrypt message, by private key,
func (k *KeyPair) Decrypt(encrtypedMsg []*big.Int) (bytes []byte) {

	for _, encryptedByte := range encrtypedMsg {
		bytes = append(bytes, uint8(big.NewInt(1).Exp(encryptedByte, k.D, k.N).Int64()))
	}

	return
}

// Sign signature for message, by private key, use the decryption number (to encrypt)
//
// it will looks like Encrypt
func (k *KeyPair) Sign(msg []byte) (signature []*big.Int) {
	for _, msgHashPart := range k.HashMsg(msg) {
		signature = append(signature, big.NewInt(1).Exp(big.NewInt(int64(msgHashPart)), k.D, k.N))
	}
	return
}

func (k *KeyPair) HashMsg(msg []byte) []byte {
	return sha256.New().Sum(msg)
}

// NewPrime with range, if n <= 1, will create a real random prime
func NewPrime(n int64) *big.Int {
	// set seed for each value
	rand.Seed(time.Now().UnixNano())
	var newRandValue func() int64
	if n <= 1 {
		newRandValue = func() int64 {
			return int64(rand.Int())
		}
	} else {
		newRandValue = func() int64 {
			return int64(rand.Intn(int(n)))
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

func BigIntToBytes(bigBytes []*big.Int) []byte {
	bytes := []byte{}

	for _, bigByte := range bigBytes {
		bytes = append(bytes, byte(bigByte.Int64()))
	}

	return bytes
}

func BytesToBigInt(bytes []byte) []*big.Int {
	bigBytes := []*big.Int{}
	for _, byte := range bytes {
		bigBytes = append(bigBytes, big.NewInt(int64(byte)))
	}
	return bigBytes
}

// NewKeyPair with parameter, put '0' as automatically long int key
func NewKeyPair(n int64) (p *KeyPair) {
	p = &KeyPair{}
	p.P = NewPrime(n)
	for {
		// not same
		p.Q = NewPrime(n)
		if p.Q.Cmp(p.P) != 0 {
			break
		}
	}
	pqValue := big.NewInt(1).Mul(
		big.NewInt(0).Sub(p.P, big.NewInt(1)),
		big.NewInt(0).Sub(p.Q, big.NewInt(1)),
	)
	pqGCD := big.NewInt(0).GCD(nil, nil, big.NewInt(0).Sub(p.P, big.NewInt(1)), big.NewInt(0).Sub(p.Q, big.NewInt(1)))

	p.L = pqValue.Div(pqValue, pqGCD)

	for {
		p.E = NewPrime(p.L.Int64())
		if big.NewInt(0).GCD(nil, nil, p.L, p.E).Cmp(big.NewInt(1)) == 0 {
			break
		}
	}

	p.D = big.NewInt(int64(1))

	for {
		p.D = p.D.Add(p.D, big.NewInt(1))
		if p.D.ModInverse(p.E, p.L) != nil || p.D.Cmp(p.L) > 0 {
			break
		}
	}

	p.N = big.NewInt(1).Mul(p.P, p.Q)
	return
}
