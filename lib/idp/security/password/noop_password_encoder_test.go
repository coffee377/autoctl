package password

import "testing"
import "github.com/stretchr/testify/assert"

func TestNoopPasswordEncoder(t *testing.T) {
	testCases := []struct {
		name        string
		rawPassword string
	}{
		{"Noop", "123456"},
		{"Noop", "test123456"},
	}
	encoder := NoopPasswordEncoder()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			encoded := encoder.Encode(tc.rawPassword)
			assert.Equal(t, encoded, tc.rawPassword, "PasswordEncoder => Encode(%s) = %s; expected %s", tc.rawPassword, encoded, tc.rawPassword)

			match := encoder.Matches(tc.rawPassword, tc.rawPassword)
			assert.True(t, match, "PasswordEncoder => Matches(%s,%s) = %v; expected %v", tc.rawPassword, tc.rawPassword, match, true)
		})
	}
}
