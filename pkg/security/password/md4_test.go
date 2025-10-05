package password

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMd4PasswordEncoder(t *testing.T) {
	testCases := []struct {
		name            string
		rawPassword     string
		encodedPassword string
		match           bool
	}{
		{"md4", "123456", "585028aa0f794af812ee3be8804eb14a", true},
		{"md4", "123456", "585028AA0F794AF812EE3BE8804EB14A", true},
		{"md4", "123456", "e50a3e40483e6db2870b60f0eec1c817", false},
		{"md4", "test123456", "239f07be2de82e5f3ad8a160bbe03050", true},
		{"md4", "test123456", "239F07BE2DE82E5F3AD8A160BBE03050", true},
		{"md4", "test123456", "e50a3e40483e6db2870b60f0eec1c817", false},
	}
	encoder := Md4PasswordEncoder()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			encoded := encoder.Encode(tc.rawPassword)
			if tc.match {
				assert.Equal(t, encoded, strings.ToLower(tc.encodedPassword), "PasswordEncoder => Encode(%s) = %s; expected %s", tc.rawPassword, encoded, tc.encodedPassword)
			} else {
				assert.NotEqual(t, encoded, strings.ToLower(tc.encodedPassword))
			}

			match := encoder.Matches(tc.rawPassword, tc.encodedPassword)
			assert.Equal(t, match, tc.match, "PasswordEncoder => Matches(%s,%s) = %v; expected %v", tc.rawPassword, tc.encodedPassword, match, tc.match)
		})
	}
}
