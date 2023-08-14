package semver

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

var versions = []string{
	"1.0.0-0", "1.0.0-1",
	"1.0.0-alpha", "1.0.0-alpha.1", "1.0.0-alpha.2", "1.0.0-alpha.beta",
	"1.0.0-beta", "1.0.0-beta.1",
	"1.0.0-beta.2", "1.0.0-beta.11",
	"1.0.0-rc", "1.0.0-rc.1", "1.0.0-rc.2",
	"1.0.0", "1.0.1", "1.1.0", "2.0.0-beta", "2.0.0",
}

func Shuffle(slice []string) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

func TestVersionsSortAsc(t *testing.T) {
	// 随机打乱顺序
	c := make([]string, len(versions))
	copy(c, versions)
	Shuffle(c)

	// 排序
	v := make(Versions, 0, len(c))
	for _, version := range versions {
		semver, _ := Version(version)
		v = append(v, semver)
	}
	v.Sort()

	// 实际结果
	var actual strings.Builder
	for i, version := range v {
		actual.WriteString(version.String())
		if i < len(versions)-1 {
			actual.WriteString(" < ")
		}
	}

	// 期望的结果
	expect := strings.Join(versions, " < ")

	if actual.String() != expect {
		t.Errorf("\nExpected: \n%s\nActual: \n%s\n", expect, actual.String())
	}

}

func TestVersionsSortDesc(t *testing.T) {
	// 随机打乱顺序
	c := make([]string, len(versions))
	copy(c, versions)
	Shuffle(c)

	// 排序
	v := make(Versions, 0, len(c))
	for _, version := range versions {
		semver, _ := Version(version)
		v = append(v, semver)
	}
	v.SortDesc()

	// 实际结果
	var actual strings.Builder

	for i, version := range v {
		actual.WriteString(version.String())
		if i < len(versions)-1 {
			actual.WriteString(" > ")
		}
	}

	// 期望的结果
	var expect strings.Builder
	for i := len(versions) - 1; i >= 0; i-- {
		expect.WriteString(versions[i])
		if i > 0 {
			expect.WriteString(" > ")
		}
	}

	if actual.String() != expect.String() {
		t.Errorf("\nExpected: \n%s\nActual: \n%s\n", expect.String(), actual.String())
	}
}

func BenchmarkVersionSort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//TestVersionsSortAsc()
	}
	b.StopTimer()
}
