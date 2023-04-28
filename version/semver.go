package semver

import (
	"errors"
	"github.com/coffee377/autoctl/log"
	"regexp"
	"strconv"
)

type Version struct {
	Major      uint64       // 主版本号：不兼容的 API 修改
	Minor      uint64       // 次版本号：向下兼容的功能性新增
	Patch      uint64       // 修订号：向下兼容的问题修正
	PreRelease []Identifier // 先行版本号
	Build      []Identifier // 版本编译信息
}

// parse parses version string and returns a validated Version or error
func parse(version string) (Version, error) {
	reg := regexp.MustCompile(VerReg)
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
func NewVersion(version string) Version {
	v, err := parse(version)
	if err != nil {
		log.Error("the %s number does not match the semantic version number, please refer to https://semver.org/lang/zh-CN/", version)
	}
	return v
}

// 判断是否是预发版本
func (version *Version) isPreRelease() bool {
	return len(version.PreRelease) > 0
}

// 预设版本置为空
func (version *Version) resetPreRelease() {
	version.PreRelease = []Identifier{}
}

// Increment increments the version
func (version *Version) Increment(release ReleaseType, identifier string, identifierBase bool) Version {
	switch release {
	case PreMajor:
		version.resetPreRelease()
		version.Patch = 0
		version.Minor = 0
		version.Major++
		version.Increment(Pre, identifier, identifierBase)
		break
	case PreMinor:
		version.resetPreRelease()
		version.Patch = 0
		version.Minor++
		version.Increment(Pre, identifier, identifierBase)
		break
	case PrePatch:
		// 如果这已经是一个预发行版，它将会在下一个版本中删除任何可能已经存在的预发行版，因为它们在这一点上是不相关的
		version.resetPreRelease()
		version.Increment(Patch, identifier, identifierBase)
		version.Increment(Pre, identifier, identifierBase)
		break
	case PreRelease:
		// 如果输入是一个非预发布版本，其作用与 PrePatch 相同
		if !version.isPreRelease() {
			version.Increment(Patch, identifier, identifierBase)
		}
		version.Increment(Pre, identifier, identifierBase)
		break
	case Major:
		// 如果这是一个 pre-major 版本，升级到相同的 major 版本，否则递增 major
		// 1.0.0-5 => 1.0.0
		// 1.1.0 => 2.0.0
		if version.Minor != 0 || version.Patch != 0 || !version.isPreRelease() {
			version.Major++
		}
		version.Minor = 0
		version.Patch = 0
		version.resetPreRelease()
		break
	case Minor:
		// 如果这是一个 pre-minor 版本，则升级到相同的 minor 版本，否则递增 minor
		// 1.2.0-5 => 1.2.0
		// 1.2.1 => 1.3.0
		if version.Patch != 0 || !version.isPreRelease() {
			version.Minor++
		}
		version.Patch = 0
		version.resetPreRelease()
	case Patch:
		// 如果这不是预发布版本，它将增加补丁号 1.2.0 => to 1.2.1
		// 如果它是一个预发布，它将上升到相同的补丁版本 1.2.0-5 => 1.2.0
		if !version.isPreRelease() {
			version.Patch++
		}
		version.resetPreRelease()
	case Pre:
		base := "0"
		if identifierBase {
			base = "1"
		}
		preReleaseIdentifiers := []Identifier{NewIdentifier(base)}

		if !version.isPreRelease() {
			version.PreRelease = preReleaseIdentifiers
		} else {
			// 从后往前解析到第一个是数字类型的 Identifier
			i := len(version.PreRelease)
			for ; i >= 0; i-- {
				identifier := version.PreRelease[i]
				if identifier.IsNumeric {
					version.PreRelease[i] = NewIdentifier(strconv.FormatUint(identifier.Num+1, 10))
					break
				}
			}
			// 未找到含有数字的 Identifier
			if i == -1 {
				// didn't increment anything
				//if (identifier === this.prerelease.join('.') && identifierBase === false) {
				//	throw new Error('invalid increment argument: identifier already exists')
				//}
				version.PreRelease = append(version.PreRelease, NewIdentifier(base))
			}
			// 如果PreRelease数组中未找到数字类型，则在数组后追加 base
			if identifier != "" {
				// alpha
				// 1.2.0-alpha => 1.2.0-alpha.1
				// 1.2.0-beta.1 bumps to 1.2.0-beta.2,
				// 1.2.0-beta.foo.bar 1.2.0-beta.foo or 1.2.0-beta bumps to 1.2.0-beta.0
				prerelease := []Identifier{NewIdentifier(identifier)}
				if identifierBase {
					prerelease = append(prerelease, NewIdentifier(base))
				}
				if version.PreRelease[0].Compare(prerelease[0]) == 0 {
					if len(prerelease) == 1 {
						version.PreRelease = prerelease
					}
				} else {
					version.PreRelease = prerelease
				}
			}
		}
		break
	}
	return *version
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

// FinalizeVersion discards prerelease and build number and only returns major, minor and patch number.
func (version *Version) FinalizeVersion() string {
	b := make([]byte, 0, 5)
	b = strconv.AppendUint(b, version.Major, 10)
	b = append(b, '.')
	b = strconv.AppendUint(b, version.Minor, 10)
	b = append(b, '.')
	b = strconv.AppendUint(b, version.Patch, 10)
	return string(b)
}

func (version *Version) Compare(other Version) int {
	return compareVersion(*version, other, true)
}

func (version *Version) CompareWithBuildMeta(other Version) int {
	return compareVersion(*version, other, false)
}
