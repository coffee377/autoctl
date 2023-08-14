package semver

// VersionChanged 版本变动类型
type VersionChanged int

type PreReleaseIdentifier string

const (
	pre        VersionChanged = 1 << 0
	Major      VersionChanged = 1 << 1
	PreMajor                  = pre | Major
	Minor      VersionChanged = 1 << 2
	PreMinor                  = pre | Minor
	Patch      VersionChanged = 1 << 3
	PrePatch                  = pre | Patch
	PreRelease VersionChanged = 1 << 4
)

const (
	alpha PreReleaseIdentifier = "alpha"
	beta  PreReleaseIdentifier = "beta"
	rc    PreReleaseIdentifier = "rc"
)

const NumberIdentifierReg = "0|[1-9]\\d*"
const VersionReg = "^(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(?:-((?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\.(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\+([0-9a-zA-Z-]+(?:\\.[0-9a-zA-Z-]+)*))?$"
