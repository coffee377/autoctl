package semver

import (
	"sort"
)

type Versions []Semver

func (semver Versions) Len() int           { return len(semver) }
func (semver Versions) Less(i, j int) bool { return semver[i].Compare(semver[j]) == -1 }
func (semver Versions) Swap(i, j int)      { semver[i], semver[j] = semver[j], semver[i] }

func (semver Versions) Sort()     { semver.SortAsc() }
func (semver Versions) SortAsc()  { semver.sort(false) }
func (semver Versions) SortDesc() { semver.sort(true) }

// Sort is a convenience method: x.Sort() calls Sort(x).
func (semver Versions) sort(desc bool) {
	if desc {
		sort.Sort(sort.Reverse(semver))
	} else {
		sort.Sort(semver)
	}
}
