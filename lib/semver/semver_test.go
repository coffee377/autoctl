package semver

import (
	"testing"
)

type data struct {
	version  string
	opts     []Option
	expected string
}

func TestVersion_Increment(t *testing.T) {
	tests := []data{
		{"0.1.0", []Option{WithPrePatch()}, "0.1.1-0"},
		{"0.1.0", []Option{WithPatch()}, "0.1.1"},

		{"0.1.1-0", []Option{WithPreRelease()}, "0.1.1-1"},
		{"0.1.1-1", []Option{WithPreRelease()}, "0.1.1-2"},
		{"0.1.1-2", []Option{WithPreRelease(), WithIdentifier(alpha)}, "0.1.1-alpha"},
		{"0.1.1-2", []Option{WithPatch()}, "0.1.1"},

		{"0.1.1-alpha", []Option{WithPreRelease(), WithIdentifier(alpha)}, "0.1.1-alpha.1"},
		{"0.1.1-alpha.1", []Option{WithPreRelease(), WithIdentifier(alpha)}, "0.1.1-alpha.2"},
		{"0.1.1-alpha.2", []Option{WithPreRelease(), WithIdentifier(beta)}, "0.1.1-beta"},
		{"0.1.1-beta", []Option{WithPreRelease(), WithIdentifier(beta)}, "0.1.1-beta.1"},
		{"0.1.1-beta.1", []Option{WithPreRelease(), WithIdentifier(rc)}, "0.1.1-rc"},
		{"0.1.1-rc", []Option{WithPreRelease(), WithIdentifier(rc)}, "0.1.1-rc.1"},

		{"0.1.1-rc", []Option{WithMinor()}, "0.2.0"},
		{"0.2.0", []Option{WithMajor()}, "1.0.0"},
		{"1.0.0", []Option{WithMinor()}, "1.1.0"},

		{"1.1.0", []Option{WithPreMinor()}, "1.2.0-0"},
		{"1.1.0", []Option{WithPreMinor(), WithIdentifier(alpha)}, "1.2.0-alpha"},

		{"1.2.0-0", []Option{WithPreRelease(), WithIdentifier(alpha)}, "1.2.0-alpha"},
		{"1.2.0-1.2", []Option{WithPreRelease()}, "1.2.0-1.3"},

		{"1.2.0-alpha", []Option{WithPreRelease(), WithIdentifier(alpha)}, "1.2.0-alpha.1"},
		{"1.2.0-alpha.1", []Option{WithPreRelease(), WithIdentifier(alpha)}, "1.2.0-alpha.2"},
		{"1.2.0-alpha.1", []Option{WithPreRelease(), WithIdentifier(beta)}, "1.2.0-beta"},
		{"1.2.0-beta.1", []Option{WithPreRelease(), WithIdentifier(beta)}, "1.2.0-beta.2"},
		{"1.2.0-beta.1", []Option{WithPreRelease(), WithIdentifier(alpha)}, "1.2.0-beta.1"},
		{"1.2.0-beta.1", []Option{WithPreRelease(), WithIdentifier(rc)}, "1.2.0-rc"},

		{"1.2.0-rc", []Option{WithMinor()}, "1.2.0"},

		{"1.2.0", []Option{WithIdentifier("alpha")}, "1.2.0"},
		{"1.2.0", []Option{WithPrePatchIdentifier(alpha)}, "1.2.1-alpha"},
		{"1.2.0-alpha", []Option{WithPrePatchIdentifier("alpha")}, "1.2.1-alpha"},
	}

	for _, test := range tests {
		result := Version(test.version).Increment(test.opts...)
		if result.String() != test.expected {
			t.Errorf("version number '%s' increment, expected '%s', but '%s' got", test.version, test.expected, result.String())
		}
	}

}

func TestVersion_IncrementMajor(t *testing.T) {
	v1 := Version("0.1.0")
	v2 := v1.IncrementMajor()
	if v2.String() != "1.0.0" {
		t.Errorf("IncrementMajor 错误")
	}
	v3 := v2.IncrementPatch()
	if v3.String() != "1.0.1" {
		t.Errorf("IncrementPatch 错误")
	}
	v4 := v3.IncrementPrePatch("")
	if v4.String() != "1.0.2-0" {
		t.Errorf("IncrementPrePatch 错误")
	}
	v5 := v4.IncrementPrePatch("")
	if v5.String() != "1.0.3-0" {
		t.Errorf("IncrementPrePatch 错误")
	}
	v6 := v5.IncrementPrePatch(alpha)
	if v6.String() != "1.0.4-alpha" {
		t.Errorf("IncrementPrePatch 错误")
	}
}
