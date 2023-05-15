package module

import (
	lib "github.com/coffee377/autoctl/lib/version"
)

type Graph interface {
	Raw() *string
	Path() *string
	Name() *string
	Version() lib.SemVer
}
