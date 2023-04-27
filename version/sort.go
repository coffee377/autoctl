package main

import (
	"fmt"
	"sort"
)

// Versions represents multiple versions.
type Versions []Version

// Len returns length of version collection
func (s Versions) Len() int {
	return len(s)
}

// Swap swaps two versions inside the collection by its indices
func (s Versions) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less checks if version at index i is less than version at index j
func (s Versions) Less(i, j int) bool {
	return s[i].Compare(s[j]) == -1
}

// Sort sorts a slice of versions
func Sort(versions []Version, desc bool) {
	if desc {
		sort.Sort(sort.Reverse(Versions(versions)))
	} else {
		sort.Sort(Versions(versions))
	}

}

func main() {

	//fmt.Printf("%v", strings.Compare("1.0.0-alpha.1", "1.0.0-beta"))
	version1, _ := NewVersion("1.0.0-alpha")
	version2, _ := NewVersion("1.0.0-beta.0")
	version3, _ := NewVersion("1.0.0")
	version4, _ := NewVersion("1.2.0-beta.0")
	version5, _ := NewVersion("1.0.0-alpha.1")
	//version1.Compare(version2)
	v := make(Versions, 0, 1)
	v = append(v, version1)
	v = append(v, version2)
	v = append(v, version3)
	v = append(v, version4)
	v = append(v, version5)
	//sort.Slice(v, func(i, j int) bool {
	//	//return v[i] < v[j] //升序  即前面的值比后面的小
	//	return v[i] > v[j] //降序  即前面的值比后面的大
	//})
	fmt.Println("排序之前")
	for _, version := range v {
		fmt.Println(version.String())
	}
	Sort(v, true)
	fmt.Println("排序之后")
	for _, version := range v {
		fmt.Println(version.String())
	}

	//NewVersion("1.0.0-alpha")
	//NewVersion("1.0.0-alpha.0")
	//NewVersion("1.0.0-beta.0+build.2")
}
