package module

import (
	"github.com/coffee377/autoctl/lib/semver"
)

type Graph interface {
	Raw() *string
	Path() *string
	Name() *string
	Version() semver.Semver
}
