package keygen

import (
	"encoding/base64"
	"encoding/hex"
)

type StringKeyGenerator interface {
	GenerateKey() string
}

type base64StringKeyGenerator struct {
	keyGenerator BytesKeyGenerator
	encoder      *base64.Encoding
}

func (b base64StringKeyGenerator) GenerateKey() string {
	key := b.keyGenerator.GenerateKey()
	return b.encoder.EncodeToString(key)
}

type Base64StringKeyGeneratorOption func(generator *base64StringKeyGenerator)

func Base64StringKeyGenerator(opts ...Base64StringKeyGeneratorOption) StringKeyGenerator {
	keyGenerator := &base64StringKeyGenerator{
		keyGenerator: RandomBytesKeyGenerator(RandomWithKeyLength(32)),
		encoder:      base64.StdEncoding,
	}
	for _, opt := range opts {
		opt(keyGenerator)
	}
	return keyGenerator
}

func Base64WithEncoder(encoder *base64.Encoding) Base64StringKeyGeneratorOption {
	return func(generator *base64StringKeyGenerator) {
		generator.encoder = encoder
	}
}

func Base64WithStdEncoder() Base64StringKeyGeneratorOption {
	return Base64WithEncoder(base64.StdEncoding)
}

func Base64WithURLEncoder() Base64StringKeyGeneratorOption {
	return Base64WithEncoder(base64.URLEncoding)
}

func Base64WithKeyLength(length int) Base64StringKeyGeneratorOption {
	return func(generator *base64StringKeyGenerator) {
		generator.keyGenerator.setKeyLength(length)
	}
}

type hexEncodingStringKeyGenerator struct {
	keyGenerator BytesKeyGenerator
}

func (h hexEncodingStringKeyGenerator) GenerateKey() string {
	key := h.keyGenerator.GenerateKey()
	return hex.EncodeToString(key)
}

func HexEncodingStringKeyGenerator(keyGenerator BytesKeyGenerator) StringKeyGenerator {
	return &hexEncodingStringKeyGenerator{
		keyGenerator: keyGenerator,
	}
}
