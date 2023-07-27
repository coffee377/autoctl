package semver

import (
	"encoding/json"
)

// MarshalJSON implements the encoding/json.Marshaller interface.
func (v *version) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

// UnmarshalJSON implements the encoding/json.Unmarshaler interface.
func (v *version) UnmarshalJSON(data []byte) (err error) {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	semver, err := Version(s)
	if err != nil {
		return err
	}
	p := semver.(*version)
	*v = *p
	return nil
}
