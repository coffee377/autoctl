package password

import "testing"

func TestDelegatingPasswordEncoder_Encode(t *testing.T) {
	encoder := CreateDelegatingPasswordEncoder()
	encode := encoder.Encode("123456")
	if !encoder.Matches("123456", encode) {
		t.Errorf("incorrect password")
	}
}
