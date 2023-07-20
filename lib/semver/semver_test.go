package semver

import (
	"testing"
)

type data struct {
	version  string
	opts     []Option
	expected string
}

func TestVersion(t *testing.T) {
	tests := []data{
		//{"0.1.0", []Option{WithPrePatch()}, "0.1.1-0"},
		//{"0.1.1-0", []Option{WithPreRelease()}, "0.1.1-1"},
		//{"0.1.1-0", []Option{WithPreRelease(), WithIdentifier(alpha)}, "0.1.1-alpha"},
		{"1.2.0-alpha", []Option{WithPreRelease(), WithIdentifier(alpha)}, "1.2.0-alpha.1"},
		{"1.2.0-beta.1", []Option{WithPreRelease(), WithIdentifier(beta)}, "1.2.0-beta.2"},
		//{"0.1.1-alpha.1", []Option{WithPreRelease(), WithIdentifier(alpha)}, "0.1.1-alpha.2"},
		//{"0.1.1-1.2", []Option{WithPreRelease()}, "0.1.1-1.3"},
		//{"0.1.0", []Option{WithPrePatch(), WithIdentifier("alpha")}, "0.1.1-alpha"},
		//{"0.1.1-alpha", []Option{WithPreRelease(), WithIdentifier("alpha")}, "0.1.1-alpha.1"},
	}

	for _, test := range tests {
		result := Version(test.version).Increment(test.opts...)
		if result.String() != test.expected {
			t.Errorf("version number '%s' increment, expected '%s', but '%s' got", test.version, test.expected, result.String())
		}
	}

}
