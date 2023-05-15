package lib

import (
	"sort"
)

type SemVerSlice []SemVer

func (semver SemVerSlice) Len() int           { return len(semver) }
func (semver SemVerSlice) Less(i, j int) bool { return semver[i].Compare(semver[j]) == -1 }
func (semver SemVerSlice) Swap(i, j int)      { semver[i], semver[j] = semver[j], semver[i] }

func (semver SemVerSlice) Sort() {
	semver.SortAsc()
}

func (semver SemVerSlice) SortAsc() {
	semver.sort(false)
}

func (semver SemVerSlice) SortDesc() {
	semver.sort(true)
}

// Sort is a convenience method: x.Sort() calls Sort(x).
func (semver SemVerSlice) sort(desc bool) {
	if desc {
		sort.Sort(sort.Reverse(semver))
	} else {
		sort.Sort(semver)
	}
}
