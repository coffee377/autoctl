package git

import (
	"bytes"
	"fmt"
	"github.com/coffee377/automation-cli/cmd"
	"os/exec"
	"regexp"
	"strings"
)

// Run 执行 git 命令
func Run(cwd string, verbose bool, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Dir = cwd
	if verbose {
		fmt.Println(cmd.String())
	}
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	return stdout.String(), nil
}

// FetchAll fetch the latest commits and tags
func FetchAll() error {
	_, err := Run(cmd.CliOpts.Cwd, false, "fetch", "--all")
	if err != nil {
		panic(err)
	}
	return nil
}

func fetchTags(releasePattern string, sort bool) []string {
	out, err := Run("", false, "git", "log", "--tags", "--simplify-by-decoration", "--pretty=\"%ai @%d\"")
	if err != nil {
		return nil
	}

	//out, _ := exec.Command("git", "log", "--tags", "--simplify-by-decoration", "--pretty=\"%ai @%d\"").Output()
	uncuratedTags := strings.Split(out, "\n")

	var tags []string
	validTag := regexp.MustCompile(fmt.Sprintf(`(.*\s)@\s.*tag:\s(%v)`, releasePattern))

	var match [][]string
	for _, tag := range uncuratedTags {
		match = validTag.FindAllStringSubmatch(tag, -1)
		if len(match) > 0 && len(match[0]) > 2 {
			// fmt.Printf("%v => %v \n", match[0][1], match[0][2])
			tags = append(tags, strings.Replace(match[0][2], ",", "", -1))
		}
	}
	//if sort {
	//	tags = utils.OnlyStable(tags)
	//	sort.Sort(utils.ByVersion(tags))
	//}
	return tags
}
