package semver

import (
	"database/sql/driver"
	"fmt"
)

// Scan implements the database/sql.Scanner interface.
func (v *version) Scan(src interface{}) (err error) {
	var str string
	switch src := src.(type) {
	case string:
		str = src
	case []byte:
		str = string(src)
	default:
		return fmt.Errorf("version.Scan: cannot convert %T to string", src)
	}

	if semver, err := Version(str); err == nil {
		*v = *semver.(*version)
	}

	return err
}

// Value implements the database/sql/driver.Valuer interface.
func (v *version) Value() (driver.Value, error) {
	return v.String(), nil
}
