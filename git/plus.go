package git

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	git "github.com/coffee377/autoctl/git/commit"
	"github.com/coffee377/autoctl/log"
	"github.com/spf13/cobra"
	"io"
	"os/exec"
	"strings"
)

type Plus struct {
	Cwd     string
	Verbose bool
	args    *[]string
	cmd     *cobra.Command
}

// Exec 执行 git 命令
func (plus *Plus) Exec(args ...string) []byte {
	command := exec.Command("git", args...)
	command.Dir = plus.Cwd
	if plus.Verbose {
		n := strings.SplitN(command.String(), " ", 2)
		log.Debug("git %s", n[1])
	}
	output, err := command.Output()
	if err != nil {
		log.Fatal("%s", err)
	}
	return output
}

// FetchAll 拉取所有远程仓库的最新内容到本地
func (plus *Plus) FetchAll() {
	_ = plus.Exec("fetch", "--all")
}

// FetchLatestTag 获取最近一次的标签
func (plus *Plus) FetchLatestTag() []byte {
	return plus.Exec("describe", "--tags", "--abbrev=0")
}

// https://zhuanlan.zhihu.com/p/87725726 打标签
//https://blog.csdn.net/qq_21746331/article/details/120776710

func (plus *Plus) FetchTags(desc bool, patterns ...string) []byte {
	var args []string
	var min = ""
	if desc {
		min = "-"
	}

	var sort = fmt.Sprintf("--sort=%screatordate", min)

	args = append(args, "tag", "-l", sort, "--format=%(objecttype);%(objectname);%(creatordate:short);%(refname:short);%(subject)")
	// pattern 参数需要放在最后
	for _, value := range patterns {
		args = append(args, value)
	}
	// git tag -l --sort=-creatordate --format='%(objectname);%(creatordate:short);%(refname:short);%(subject)'
	bytes := plus.Exec(args...)
	s := string(bytes)
	log.Info(s)

	return nil
}

// FetchLogs will query commit records between two tags
func (plus *Plus) FetchLogs(commit1, commit2 string) []byte {
	//var notation string
	//if len(tag1) > 0 && len(tag2) > 0 {
	//	notation = fmt.Sprintf("%s..%s", tag1, tag2)
	//} else if len(tag1) > 0 {
	//	notation = tag1
	//} else if len(tag2) > 0 {
	//	notation = tag2
	//}
	var args []string

	format := fmt.Sprintf("--pretty=%s", git.RecordFormat)
	args = append(args, "log", "--no-merges", "--date=format:%Y/%m/%d %H:%M:%S", format)

	args = append(args, "-i")
	args = append(args, "-E", "--grep", "fix")
	//append(args, )

	//notation := ""
	//stdout := bytes.Buffer
	result := plus.Exec(args...)
	reader := bufio.NewReader(bytes.NewReader(result))
	//line, prefix, err := reader.ReadLine()
	var records []*git.CommitRecord
	for {
		line, prefix, err := reader.ReadLine() //以'\n'为结束符读入一行
		record := git.NewCommitRecord(string(line))
		records = append(records, record)
		if err != nil || io.EOF == err {
			break
		}
		fmt.Println(line)
		fmt.Println(prefix)
		fmt.Println(line)
	}
	marshal, _ := json.Marshal(records)
	//bytes.
	return marshal
}
