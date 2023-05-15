package lib

// ReleaseType 发布类型
type ReleaseType int

const (
	Pre        ReleaseType = 1 << 0
	Major      ReleaseType = 1 << 1
	PreMajor               = Pre | Major
	Minor      ReleaseType = 1 << 2
	PreMinor               = Pre | Minor
	Patch      ReleaseType = 1 << 3
	PrePatch               = Pre | Patch
	PreRelease ReleaseType = 1 << 4
)

const NumberIdentifierReg = "0|[1-9]\\d*"
const VerReg = "^(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(?:-((?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\.(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\+([0-9a-zA-Z-]+(?:\\.[0-9a-zA-Z-]+)*))?$"
