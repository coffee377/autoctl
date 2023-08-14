package password

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"strings"
)

type Digester interface {
	Digest([]byte) []byte
}

type digester struct {
	hash       hash.Hash
	hashes     map[string]hash.Hash
	iterations int
}

func (d digester) Digest(value []byte) []byte {
	for i := 0; i < d.iterations; i++ {
		d.hash.Reset()
		d.hash.Write(value)
		//str := "Hello, World!"
		//hash := sha1.New()
		//hash.Write([]byte(str))
		//encrypted := hash.Sum(nil)
		value = d.hash.Sum(nil)
	}
	return value
}

// NewDigester Create a new Digester
// @param algorithm the digest algorithm; for example, "SHA-1" or "SHA-256".
// @param iterations the number of times to apply the digest algorithm to the input
func NewDigester(algorithm string, iterations int) Digester {
	d := &digester{
		iterations: 1,
		hashes: map[string]hash.Hash{
			"MD5":     md5.New(),
			"SHA-1":   sha1.New(),
			"SHA-256": sha256.New(),
			"SHA-512": sha512.New(),
		},
	}
	if iterations > 0 {
		d.iterations = iterations
	}
	h := d.hashes[strings.ToUpper(algorithm)]
	if h == nil {
		panic("No such hashing algorithm")
	}
	d.hash = h

	return d
}
