package password

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/coffee377/autoctl/pkg/security/keygen"
)

// Deprecated
type messageDigestPasswordEncoder struct {
	saltGenerator      keygen.StringKeyGenerator
	digester           Digester
	encodeHashAsBase64 bool
}

func (e messageDigestPasswordEncoder) Encode(rawPassword string) string {
	salt := fmt.Sprintf("{%s}", e.saltGenerator.GenerateKey())
	return e.digest(salt, rawPassword)
}

func (e messageDigestPasswordEncoder) Matches(rawPassword string, encodedPassword string) bool {
	salt := e.extractSalt(encodedPassword)
	rawPasswordEncoded := e.digest(salt, rawPassword)
	b := rawPasswordEncoded == encodedPassword
	return b
}

func (e messageDigestPasswordEncoder) UpgradeEncoding(encodedPassword string) bool {
	return false
}

func (e messageDigestPasswordEncoder) digest(salt string, rawPassword string) string {
	saltedPassword := strings.Join([]string{rawPassword, salt}, "")
	digest := e.digester.Digest([]byte(saltedPassword))
	encoded := e.encode(digest)
	return salt + encoded
}

func (e messageDigestPasswordEncoder) encode(digest []byte) string {
	if e.encodeHashAsBase64 {
		return base64.StdEncoding.EncodeToString(digest)
	}
	return hex.EncodeToString(digest)
}

func (e messageDigestPasswordEncoder) extractSalt(prefixEncodedPassword string) string {
	if prefixEncodedPassword != "" {
		start := strings.Index(prefixEncodedPassword, "{")
		end := strings.Index(prefixEncodedPassword, "}")
		if start == 0 && end > 0 {
			return prefixEncodedPassword[0 : end+1]
		}
	}
	return ""
}

// Deprecated
func MessageDigestPasswordEncoder(algorithm string) Encoder {
	encoder := &messageDigestPasswordEncoder{
		encodeHashAsBase64: false,
		saltGenerator:      keygen.Base64StringKeyGenerator(),
		digester:           NewDigester(algorithm, 1),
	}
	return encoder
}
