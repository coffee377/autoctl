package git

import (
	"bytes"
	"fmt"
	git "github.com/coffee377/autoctl/git/commit"
	"github.com/spf13/cobra"
	"os/exec"
)

type Plus struct {
	Cwd     string
	Verbose bool
	args    *[]string
	cmd     *cobra.Command
}

// Exec 执行 git 命令
func (plus *Plus) Exec(args ...string) ([]byte, []byte) {
	command := exec.Command("git", args...)
	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)
	command.Stdout = &stdout
	command.Stderr = &stderr
	command.Dir = plus.Cwd
	if plus.Verbose {
		fmt.Println(command.String())
	}
	if err := command.Run(); err != nil {
		stderr.Write([]byte(err.Error()))
	}
	return stdout.Bytes(), stderr.Bytes()
}

// FetchAll 拉取所有远程仓库的最新内容到本地
func (plus *Plus) FetchAll() {
	_, _ = plus.Exec("fetch", "--all")
}

func (plus *Plus) FetchTags(releasePattern string, sort bool) []string {
	//t := plus.Exec("log", "--tags", "--simplify-by-decoration", "--pretty=\"%ai @%d\"")
	//log.Warn("%v", string(t))
	//if err != nil {
	//	return nil
	//}
	//
	////out, _ := exec.Command("git", "log", "--tags", "--simplify-by-decoration", "--pretty=\"%ai @%d\"").Output()
	//uncuratedTags := strings.Split(out, "\n")

	//var tags []string
	//validTag := regexp.MustCompile(fmt.Sprintf(`(.*\s)@\s.*tag:\s(%v)`, releasePattern))
	//
	//var match [][]string
	//for _, tag := range uncuratedTags {
	//	match = validTag.FindAllStringSubmatch(tag, -1)
	//	if len(match) > 0 && len(match[0]) > 2 {
	//		// fmt.Printf("%v => %v \n", match[0][1], match[0][2])
	//		tags = append(tags, strings.Replace(match[0][2], ",", "", -1))
	//	}
	//}
	//if sort {
	//	tags = utils.OnlyStable(tags)
	//	sort.Sort(utils.ByVersion(tags))
	//}
	var tags []string
	return tags
}

// FetchLogs will query commit records between two tags
func (plus *Plus) FetchLogs(commit1, commit2 string) ([]byte, []byte) {
	//var notation string
	//if len(tag1) > 0 && len(tag2) > 0 {
	//	notation = fmt.Sprintf("%s..%s", tag1, tag2)
	//} else if len(tag1) > 0 {
	//	notation = tag1
	//} else if len(tag2) > 0 {
	//	notation = tag2
	//}
	var args []string

	format := fmt.Sprintf("--pretty=format:%s", git.GitRecordFormat)
	args = append(args, "log", "--no-merges", "--date=format:%Y/%m/%d %H:%M:%S", format)

	//append(args, )

	//format := "--pretty=format:{\"hash\":\"%H\",\"abbrevHash\":\"%h\",\"author\":\"%an\",\"email\":\"%ae\",\"date\":\"%ad\",\"timestamp\":%at,\"message\":\"%f\"}"
	//format := "--pretty=format:{\"timestamp\":%at,\"subject\":\"%s\",\"body\":\"%b\",\"message\":\"%B\"}"
	//format := "--pretty=format:%B%n%n"
	//format := "--pretty=format:%<(30,mtrunc)%s %ai" // 下一个占位符最多占用 30 个字符,超过将进行截断处理（ltrunc|mtrunc|trunc）
	//notation := ""
	return plus.Exec(args...)
}
