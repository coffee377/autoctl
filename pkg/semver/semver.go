package semver

import (
	"errors"
	"github.com/coffee377/autoctl/pkg/log"
	"regexp"
	"strconv"
)

type options struct {
	changed    VersionChanged
	identifier PreReleaseIdentifier
	//identifierBase bool
}

type Option func(options *options) error

type Semver interface {
	Major() uint64
	Minor() uint64
	Patch() uint64

	PreRelease() []Identifier
	Build() []Identifier

	Increment(opts ...Option) Semver

	IncrementMajor() Semver
	IncrementMinor() Semver
	IncrementPatch() Semver

	IncrementPreMajor(identifier PreReleaseIdentifier) Semver
	IncrementPreMinor(identifier PreReleaseIdentifier) Semver
	IncrementPrePatch(identifier PreReleaseIdentifier) Semver

	IncrementPreRelease(identifier PreReleaseIdentifier) Semver

	String() string
	FinalizeVersion() string
	Compare(other Semver) int
	CompareWithBuildMeta(other Semver) int
}

type version struct {
	major      uint64       // 主版本号：不兼容的 API 修改
	minor      uint64       // 次版本号：向下兼容的功能性新增
	patch      uint64       // 修订号：向下兼容的问题修正
	preRelease []Identifier // 先行版本号
	build      []Identifier // 版本编译信息
	options    *options     // 配置选项
}

// parse parses version string and returns a validated Semver or error
func parse(ver string) (version, error) {
	reg := regexp.MustCompile(VersionReg)
	if !reg.MatchString(ver) {
		return version{}, errors.New("the version number does not match the semantic version number, please refer to https://semver.org/lang/zh-CN/")
	}
	match := reg.FindStringSubmatch(ver)
	v := version{options: &options{}}
	v.major, _ = strconv.ParseUint(match[1], 10, 64)
	v.minor, _ = strconv.ParseUint(match[2], 10, 64)
	v.patch, _ = strconv.ParseUint(match[3], 10, 64)

	if match[4] != "" {
		v.preRelease = parseIdentifiers(match[4])
	}
	if match[5] != "" {
		v.build = parseIdentifiers(match[5])
	}
	return v, nil
}

// Version is an alias for Parse and returns a pointer, parses version string and returns a validated Semver or error
func Version(version string) (Semver, error) {
	v, err := parse(version)
	if err != nil {
		//log.Error("the %s number does not match the semantic version number, please refer to https://semver.org/lang/zh-CN/", version)
		return nil, err
	}
	return &v, nil
}

func (v *version) Major() uint64 {
	return v.major
}

func (v *version) Minor() uint64 {
	return v.minor
}

func (v *version) Patch() uint64 {
	return v.patch
}

func (v *version) PreRelease() []Identifier {
	return v.preRelease
}

func (v *version) Build() []Identifier {
	return v.build
}

// Increment increments the version
func (v *version) Increment(opts ...Option) Semver {
	ver := *v
	increment(&ver, opts...)
	return &ver
}

func (v *version) IncrementMajor() Semver {
	return v.Increment(WithMajor())
}

func (v *version) IncrementMinor() Semver {
	return v.Increment(WithMinor())
}

func (v *version) IncrementPatch() Semver {
	return v.Increment(WithPatch())
}

func (v *version) IncrementPreMajor(identifier PreReleaseIdentifier) Semver {
	return v.Increment(WithPreMinorIdentifier(identifier))
}

func (v *version) IncrementPreMinor(identifier PreReleaseIdentifier) Semver {
	return v.Increment(WithPreMinorIdentifier(identifier))
}

func (v *version) IncrementPrePatch(identifier PreReleaseIdentifier) Semver {
	return v.Increment(WithPrePatchIdentifier(identifier))
}

func (v *version) IncrementPreRelease(identifier PreReleaseIdentifier) Semver {
	return v.Increment(WithPreReleaseIdentifier(identifier))
}

func (v *version) String() string {
	buffer := v.versionBase()

	if len(v.preRelease) > 0 {
		buffer = append(buffer, '-')
		buffer = append(buffer, v.preRelease[0].Raw...)

		for _, pre := range v.preRelease[1:] {
			buffer = append(buffer, '.')
			buffer = append(buffer, pre.Raw...)
		}
	}

	if len(v.build) > 0 {
		buffer = append(buffer, '+')
		buffer = append(buffer, v.build[0].Raw...)

		for _, build := range v.build[1:] {
			buffer = append(buffer, '.')
			buffer = append(buffer, build.Raw...)
		}
	}

	return string(buffer)
}

// FinalizeVersion discards prerelease and build number and only returns major, minor and patch number.
func (v *version) FinalizeVersion() string {
	b := v.versionBase()
	return string(b)
}

func (v *version) Compare(other Semver) int {
	return compareVersion(v, other, true)
}

func (v *version) CompareWithBuildMeta(other Semver) int {
	return compareVersion(v, other, false)
}

// 判断是否是预发版本
func (v *version) isPreRelease() bool {
	return len(v.preRelease) > 0
}

// 预设版本置为空
func (v *version) resetPreRelease() {
	v.preRelease = []Identifier{}
}

func (v *version) versionBase() []byte {
	buffer := make([]byte, 0, 5)
	buffer = strconv.AppendUint(buffer, v.major, 10)
	buffer = append(buffer, '.')
	buffer = strconv.AppendUint(buffer, v.minor, 10)
	buffer = append(buffer, '.')
	buffer = strconv.AppendUint(buffer, v.patch, 10)
	return buffer
}

// 预发布版本号递增：从后往前解析到第一个是数字类型的 Identifier
func lastNumberIdentifierIncrease(v *version) {
	found := false
	l := len(v.preRelease)
	for i := 0; i < l; i++ {
		identifier := v.preRelease[l-i-1]
		if identifier.IsNumeric {
			found = true
			// 递增版本号
			v.preRelease[l-i-1] = NewIdentifier(strconv.FormatUint(identifier.Num+1, 10))
			break
		}
	}
	// 未找到含有数字的 Identifier
	// 如果PreRelease数组中未找到数字类型，则在数组后追加 base
	if !found {
		v.preRelease = append(v.preRelease, NewIdentifier("1"))
	}
}

func increment(v *version, opts ...Option) Semver {
	for _, opt := range opts {
		_ = opt(v.options)
	}

	switch v.options.changed {
	case PreMajor:
		v.resetPreRelease()
		v.patch = 0
		v.minor = 0
		v.major++

		increment(v, withPre())
		break
	case PreMinor:
		v.resetPreRelease()
		v.patch = 0
		v.minor++

		increment(v, withPre())
		break
	case PrePatch:
		// 如果这已经是一个预发行版，它将会在下一个版本中删除任何可能已经存在的预发行版，因为它们在这一点上是不相关的
		v.resetPreRelease()

		increment(v, WithPatch())
		increment(v, withPre())
		break
	case PreRelease:
		// 如果输入是一个非预发布版本，其作用与 PrePatch 相同
		if !v.isPreRelease() {
			increment(v, WithPatch())
		}
		increment(v, withPre())
		break
	case Major:
		// 如果这是一个 pre-major 版本，升级到相同的 major 版本，否则递增 major
		// 1.0.0-5 => 1.0.0
		// 1.1.0 => 2.0.0
		if v.minor != 0 || v.patch != 0 || !v.isPreRelease() {
			v.major++
		}
		v.minor = 0
		v.patch = 0
		v.resetPreRelease()
		break
	case Minor:
		// 如果这是一个 pre-minor 版本，则升级到相同的 minor 版本，否则递增 minor
		// 1.2.1 => 1.3.0
		// 1.2.0-5 => 1.2.0
		if v.patch != 0 || !v.isPreRelease() {
			v.minor++
		}
		v.patch = 0
		v.resetPreRelease()
	case Patch:
		// 如果这不是预发布版本，它将增加补丁号 1.2.0 => to 1.2.1
		if !v.isPreRelease() {
			v.patch++
		}
		// 如果它是一个预发布，它将上升到相同的补丁版本 1.2.0-5 => 1.2.0
		v.resetPreRelease()
	case pre:
		identifier := v.options.identifier
		identifiers := parseIdentifiers("0")
		if identifier != "" {
			identifiers = parseIdentifiers(string(identifier))
		}
		// 不是预发版本
		if !v.isPreRelease() {
			v.preRelease = identifiers
		} else {
			compare := v.preRelease[0].Compare(identifiers[0])
			if identifier != "" {
				if compare == 0 {
					lastNumberIdentifierIncrease(v)
				} else if compare == -1 {
					v.preRelease = identifiers
				} else if compare == 1 {
					// 不处理，使用原版本号，日志警告提示
					log.Warn("预发布版本不能进行降级，因为 %s < %s", identifiers[0].Raw, v.preRelease[0].Raw)
				}
				return v
			}
			lastNumberIdentifierIncrease(v)
		}
	}
	return v
}
