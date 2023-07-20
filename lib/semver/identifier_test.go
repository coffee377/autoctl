package semver

import (
	"testing"
)

func TestComparatorIdentifier(t *testing.T) {
	tests := []struct {
		c []string
		r string
	}{
		{[]string{"1", "2"}, "<"},
		{[]string{"1", "a"}, "<"},
		{[]string{"A", "B"}, "<"},
		{[]string{"a", "b"}, "<"},
		{[]string{"a", "A"}, ">"},
		{[]string{"a", "B"}, ">"},
		{[]string{"1", "11"}, "<"},
	}
	for _, test := range tests {
		i1 := NewIdentifier(test.c[0])
		i2 := NewIdentifier(test.c[1])
		i := i1.Compare(i2)
		r := ""
		if i > 0 {
			r = ">"
		} else if i < 0 {
			r = "<"
		} else {
			r = "="
		}
		if test.r != r {
			t.Errorf("%s compare %s expected be '%s', but '%s' got", i1.Raw, i2.Raw, test.r, r)
		}
	}
}
