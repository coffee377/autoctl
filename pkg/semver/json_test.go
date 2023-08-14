package semver

import (
	"encoding/json"
	"strconv"
	"testing"
)

func TestVersion_MarshalJSON(t *testing.T) {
	versionString := "3.1.4-alpha.1.5.9+build.2.6.5"
	semver, err := Version(versionString)
	if err != nil {
		t.Fatal(err)
	}

	versionJSON, err := json.Marshal(semver)
	if err != nil {
		t.Fatal(err)
	}

	quotedVersionString := strconv.Quote(versionString)

	if string(versionJSON) != quotedVersionString {
		t.Fatalf("JSON marshaled semantic version not equal: expected %q, got %q", quotedVersionString, string(versionJSON))
	}
}

func TestVersion_UnmarshalJSON(t *testing.T) {
	versionString := "3.1.4-alpha.1.5.9+build.2.6.5"
	quotedVersionString := strconv.Quote(versionString)

	var v version
	if err := json.Unmarshal([]byte(quotedVersionString), &v); err != nil {
		t.Fatal(err)
	}

	if v.String() != versionString {
		t.Fatalf("JSON unmarshaled semantic version not equal: expected %q, got %q", versionString, v.String())
	}

	badVersionString := strconv.Quote("3.1.4.1.5.9.2.6.5-other-digits-of-pi")
	if err := json.Unmarshal([]byte(badVersionString), &v); err == nil {
		t.Fatal("expected JSON unmarshal error, got nil")
	}

	if err := json.Unmarshal([]byte("3.1"), &v); err == nil {
		t.Fatal("expected JSON unmarshal error, got nil")
	}
}
