package keygen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type StringEncodingVector struct {
	keyLength    int
	stringLength int
}

var testHexEncodingVectors = []StringEncodingVector{
	{8, 16},
	{9, 18},
	{10, 20},
	{11, 22},
	{12, 24},
	{13, 26},
	{14, 28},
	{15, 30},
	{16, 32},
}

var testBase64EncodingVectors = []StringEncodingVector{
	{7, 12},
	{8, 12},
	{9, 12},
	{10, 16},
	{11, 16},
	{12, 16},
	{13, 20},
	{14, 20},
	{15, 20},
	{16, 24},
	{17, 24},
	{18, 24},
	{19, 28},
	{20, 28},
	{21, 28},
	{22, 32},
	{23, 32},
	{24, 32},
	{32, 44},
}

func testStringKeyGenerator(t *testing.T, stringKeyGenerator StringKeyGenerator, vectors []StringEncodingVector) {
	for _, v := range vectors {
		keyGenerator := RandomBytesKeyGenerator(RandomWithKeyLength(v.keyLength), RandomWithVisibleCode())
		var skg StringKeyGenerator

		if _, ok := stringKeyGenerator.(hexEncodingStringKeyGenerator); ok {
			skg = HexEncodingStringKeyGenerator(keyGenerator)
		}

		if _, ok := stringKeyGenerator.(base64StringKeyGenerator); ok {
			skg = Base64StringKeyGenerator(Base64WithKeyLength(v.keyLength), Base64WithStdEncoder())
		}

		s := skg.GenerateKey()
		n := v.keyLength / 3
		m := v.keyLength % 3
		if m > 0 {
			n += 1
		}
		assert.Equal(t, v.stringLength, n*4)
		assert.Equal(t, v.stringLength, len(s))
	}
}

func TestBase64StringKeyGenerator(t *testing.T) {
	testStringKeyGenerator(t, base64StringKeyGenerator{}, testBase64EncodingVectors)
}

func TestHexEncodingStringKeyGenerator(t *testing.T) {
	testStringKeyGenerator(t, hexEncodingStringKeyGenerator{}, testHexEncodingVectors)
}
