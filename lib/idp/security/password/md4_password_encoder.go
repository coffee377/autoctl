package password

import (
	"fmt"
	"golang.org/x/crypto/md4"
	"hash"
	"io"
	"strings"
)

type md4PasswordEncoder struct {
	md4 hash.Hash
}

func Md4PasswordEncoder() Encoder {
	return &md4PasswordEncoder{
		md4: md4.New(),
	}
}

func (e md4PasswordEncoder) Encode(rawPassword string) string {
	e.md4.Reset()
	_, err := io.WriteString(e.md4, rawPassword)
	if err != nil {
		panic(err.Error())
	}
	result := fmt.Sprintf("%x", e.md4.Sum(nil))
	return result
}

func (e md4PasswordEncoder) Matches(rawPassword string, encodedPassword string) bool {
	return strings.EqualFold(e.Encode(rawPassword), encodedPassword)
}

func (e md4PasswordEncoder) UpgradeEncoding(string) bool {
	return false
}
