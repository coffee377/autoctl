package semver

import (
	"strconv"
	"strings"
)

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
	// compareNum := math.Min(len(version.PreRelease), len(other.PreRelease))
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

func compareVersion(version Version, other Version, ignoreBuildMeta bool) int {
	// 主版本号比较
	major := compare(version.Major, other.Major)
	if major != 0 {
		return major
	}

	// 次版本号比较
	minor := compare(version.Minor, other.Minor)
	if minor != 0 {
		return minor
	}

	// 补丁号比较
	patch := compare(version.Patch, other.Patch)
	if patch != 0 {
		return patch
	}

	// 先行版本号比较
	pre := comparePre(version, other)

	if !ignoreBuildMeta && pre == 0 {
		return compareBuild(version, other)
	}

	return pre
}

func CompareVersion(a, b string, ignoreBuildMeta bool) int {
	return compareVersion(NewVersion(a), NewVersion(b), ignoreBuildMeta)
}

func parseIdentifiers(str string) []Identifier {
	splits := strings.Split(str, ".")
	identifiers := make([]Identifier, 0, len(splits))
	for _, split := range splits {
		identifiers = append(identifiers, NewIdentifier(split))
	}
	return identifiers
}
