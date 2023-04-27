package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Identifier struct {
	Raw       string
	Num       uint64
	IsNumeric bool
}

func (identifier Identifier) Compare(other Identifier) int {
	if identifier.IsNumeric && !other.IsNumeric {
		return -1
	} else if !identifier.IsNumeric && other.IsNumeric {
		return 1
	} else {
		return strings.Compare(identifier.Raw, other.Raw)
	}
}

func NewIdentifier(identifier string) Identifier {
	res := Identifier{
		Raw: identifier,
	}
	result, _ := regexp.MatchString("0|[1-9]\\d*", identifier)
	res.IsNumeric = result
	if result {
		res.Num, _ = strconv.ParseUint(identifier, 10, 64)
	}
	return res
}

func parseIdentifiers(str string) []Identifier {
	splits := strings.Split(str, ".")
	identifiers := make([]Identifier, 0, len(splits))
	for _, split := range splits {
		identifiers = append(identifiers, NewIdentifier(split))
	}
	return identifiers
}

// Version changelog
type Version struct {
	Major      uint64       // 主版本号
	Minor      uint64       // 次版本号
	Patch      uint64       // 修订号
	PreRelease []Identifier // 先行版本号
	Build      []Identifier // 版本编译信息
}

// Parse parses version string and returns a validated Version or error
func Parse(version string) (Version, error) {
	reg := regexp.MustCompile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)
	if !reg.MatchString(version) {
		return Version{}, errors.New("the version number does not match the semantic version number, please refer to https://semver.org/lang/zh-CN/")
	}
	match := reg.FindStringSubmatch(version)
	v := Version{}
	v.Major, _ = strconv.ParseUint(match[1], 10, 64)
	v.Minor, _ = strconv.ParseUint(match[2], 10, 64)
	v.Patch, _ = strconv.ParseUint(match[3], 10, 64)

	if match[4] != "" {
		v.PreRelease = parseIdentifiers(match[4])
	}
	if match[5] != "" {
		v.Build = parseIdentifiers(match[5])
	}
	return v, nil
}

// NewVersion is an alias for Parse and returns a pointer, parses version string and returns a validated Version or error
func NewVersion(version string) (Version, error) {
	v, err := Parse(version)
	return v, err
}

func (version *Version) String() string {
	buffer := make([]byte, 0, 5)
	buffer = strconv.AppendUint(buffer, version.Major, 10)
	buffer = append(buffer, '.')
	buffer = strconv.AppendUint(buffer, version.Minor, 10)
	buffer = append(buffer, '.')
	buffer = strconv.AppendUint(buffer, version.Patch, 10)

	if len(version.PreRelease) > 0 {
		buffer = append(buffer, '-')
		buffer = append(buffer, version.PreRelease[0].Raw...)

		for _, pre := range version.PreRelease[1:] {
			buffer = append(buffer, '.')
			buffer = append(buffer, pre.Raw...)
		}
	}

	if len(version.Build) > 0 {
		buffer = append(buffer, '+')
		buffer = append(buffer, version.Build[0].Raw...)

		for _, build := range version.Build[1:] {
			buffer = append(buffer, '.')
			buffer = append(buffer, build.Raw...)
		}
	}

	return string(buffer)
}

// FinalizeVersion discards prerelease and build number and only returns
// major, minor and patch number.
func (version *Version) FinalizeVersion() string {
	b := make([]byte, 0, 5)
	b = strconv.AppendUint(b, version.Major, 10)
	b = append(b, '.')
	b = strconv.AppendUint(b, version.Minor, 10)
	b = append(b, '.')
	b = strconv.AppendUint(b, version.Patch, 10)
	return string(b)
}

func compareBuildMeta(version, other *Version, compareBuildMeta bool) int {
	if version.Major != other.Major {
		return compare(version.Major, other.Major)
	}
	if version.Minor != other.Minor {
		return compare(version.Minor, other.Minor)
	}
	if version.Patch != other.Patch {
		return compare(version.Patch, other.Patch)
	}

	result := comparePre(*version, *other)

	if compareBuildMeta && result == 0 {
		return compareBuild(*version, *other)
	}
	return result
}

// Compare compares Versions v to o:
// -1 == version is less than other
// 0 == version is equal to other
// 1 == version is greater than other
func (version *Version) Compare(other Version) int {
	return compareBuildMeta(version, &other, false)
}

func (version *Version) CompareWithBuildMeta(other Version) int {
	return compareBuildMeta(version, &other, true)
}

func compare(a, b uint64) int {
	return strings.Compare(strconv.FormatUint(a, 10), strconv.FormatUint(b, 10))
}

func compareIdentifier(a, b []Identifier) int {

	// Quick comparison if a version has no prerelease versions
	if len(a) == 0 && len(b) == 0 {
		return 0
	} else if len(a) == 0 && len(b) > 0 {
		return 1
	} else if len(a) > 0 && len(b) == 0 {
		return -1
	}

	i := 0
	////compareNum := math.Min(len(version.PreRelease), len(other.PreRelease))
	for ; i < len(a) && i < len(b); i++ {
		if comp := a[i].Compare(b[i]); comp == 0 {
			continue
		} else {
			return comp
		}
	}

	// If all pre versions are the equal but one has further pre version, this one greater
	if i == len(a) && i == len(b) {
		return 0
	} else if i == len(a) && i < len(b) {
		return -1
	} else {
		return 1
	}
}

// 比较先行版本号
func comparePre(a, b Version) int {
	return compareIdentifier(a.PreRelease, b.PreRelease)
}

// 比较构建信息
func compareBuild(a, b Version) int {
	return compareIdentifier(a.Build, b.Build)
}

// IncrementPreRelease increments the patch version
func (version *Version) IncrementPreRelease() {
	//version.PreRelease.Increment()
	//version.Patch++
	//version.Pre = PreRelease{}
}

// IncrementPatch increments the patch version
func (version *Version) IncrementPatch() {
	version.Patch++
	//version.PreRelease = PreRelease{}
}

// IncrementMinor increments the minor version
func (version *Version) IncrementMinor() {
	version.Minor++
	version.Patch = 0
}

// IncrementMajor increments the major version
func (version *Version) IncrementMajor() {
	version.Major++
	version.Minor = 0
	version.Patch = 0
}
