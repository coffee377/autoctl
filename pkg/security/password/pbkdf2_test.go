package password

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testVector struct {
	password string
	secret   string
	iter     int
	hex      string
}

var sha1TestVectors = []testVector{
	{
		"password",
		"",
		185000,
		"a4b2dde07244ef111e6a63cbc15d90dcc405defd67e4cf05b3681de2b551532a41111e3d72675880",
	},
	{
		"password",
		"",
		185000,
		"1e202aeecc69e46941b22aa9de0e52837d3fb63d9503156fc782d6e7f57c47eef5e1ffdd2b6006d5",
	},
	{
		"password",
		"salt",
		185000,
		"97341311975378e480ab0eb95968daa2ccc3a926e81c4ccd2cf1068c843edadad89129e4113f4c07",
	},
}

var sha256TestVectors = []testVector{
	{
		"password",
		"",
		185000,
		"133d889464958be997e7c18f7ba564642c5af3640e9b583b7a098d2499825a91eb29306981f823bb",
	},
}

var sha512TestVectors = []testVector{
	{
		"password",
		"",
		185000,
		"11d1451a55c5a09310b70180939c30dc4fad21e0d457acd36fd807d00fada3c1a1a03fcfdecb737d",
	},
}

func testHash(t *testing.T, hashName string, vectors []testVector) {
	for _, v := range vectors {
		encoder := Pbkdf2PasswordEncoder(Pbkdf2WithAlgorithm(hashName))
		out := encoder.Encode(v.password)
		match1 := encoder.Matches(v.password, out)
		assert.Equal(t, match1, true)
		match2 := encoder.Matches(v.password, v.hex)
		assert.Equal(t, match2, true)
	}
}

func TestWithHMACSHA1(t *testing.T) {
	testHash(t, "SHA1", sha1TestVectors)
}

func TestWithHMACSHA256(t *testing.T) {
	testHash(t, "SHA256", sha256TestVectors)
}

func TestWithHMACSHA512(t *testing.T) {
	testHash(t, "SHA512", sha512TestVectors)
}
