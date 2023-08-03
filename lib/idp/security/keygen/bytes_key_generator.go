package keygen

import (
	"crypto/rand"
)

// BytesKeyGenerator A generator for unique byte array-based keys.
type BytesKeyGenerator interface {
	// GetKeyLength Get the length, in bytes, of keys created by this generator.
	// Unique keys are at least 8 bytes in length.
	GetKeyLength() int

	setKeyLength(length int)

	// GenerateKey Generate a new key.
	GenerateKey() []byte
}

type secureRandomBytesKeyGenerator struct {
	keyLength   int
	visibleCode bool
}

func (g *secureRandomBytesKeyGenerator) GetKeyLength() int {
	return g.keyLength
}

func (g *secureRandomBytesKeyGenerator) setKeyLength(length int) {
	g.keyLength = length
	if length <= 0 {
		g.keyLength = 8
	}
}

func (g *secureRandomBytesKeyGenerator) GenerateKey() []byte {
	randomBytes := make([]byte, g.keyLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	if g.visibleCode {
		for i, randomByte := range randomBytes {
			randomBytes[i] = randomByte%95 + 32
		}
	}
	return randomBytes
}

type BytesKeyGeneratorOption func(g *secureRandomBytesKeyGenerator)

func RandomBytesKeyGenerator(opts ...BytesKeyGeneratorOption) BytesKeyGenerator {
	generator := &secureRandomBytesKeyGenerator{
		keyLength:   8,
		visibleCode: false,
	}
	for _, opt := range opts {
		opt(generator)
	}

	return generator
}

func RandomWithKeyLength(length int) BytesKeyGeneratorOption {
	return func(g *secureRandomBytesKeyGenerator) {
		g.setKeyLength(length)
	}
}

func RandomWithVisibleCode() BytesKeyGeneratorOption {
	return func(g *secureRandomBytesKeyGenerator) {
		g.visibleCode = true
	}
}
