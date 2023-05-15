package lib

import (
	"fmt"
	"path/filepath"
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
func comparePre(a, b SemVer) int {
	return compareIdentifier(a.PreRelease(), b.PreRelease())
}

// 比较构建信息
func compareBuild(a, b SemVer) int {
	return compareIdentifier(a.Build(), b.Build())
}

func compareVersion(a SemVer, b SemVer, ignoreBuildMeta bool) int {
	// 主版本号比较
	major := compare(a.Major(), b.Major())
	if major != 0 {
		return major
	}

	// 次版本号比较
	minor := compare(a.Minor(), b.Minor())
	if minor != 0 {
		return minor
	}

	// 补丁号比较
	patch := compare(a.Patch(), b.Patch())
	if patch != 0 {
		return patch
	}

	// 先行版本号比较
	pre := comparePre(a, b)

	if !ignoreBuildMeta && pre == 0 {
		return compareBuild(a, b)
	}

	return pre
}

func CompareVersion(a, b string, ignoreBuildMeta bool) int {
	return compareVersion(NewVersion(a), NewVersion(b), ignoreBuildMeta)
}

func listModules() {
	//print()
	Path1 := "D:/Project/jqsoft/framework/frontend-base" //absolute path1

	pattern := Path1 + "/**/*/package.json"
	files, err := filepath.Glob(pattern)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, file := range files {
		relPath, _ := filepath.Rel(Path1, file)
		fmt.Println("Files:", relPath)
	}
	//fmt.Println("Files:", files)
}
